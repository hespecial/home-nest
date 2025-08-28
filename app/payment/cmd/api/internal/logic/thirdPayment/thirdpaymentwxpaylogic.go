package thirdPayment

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"home-nest/app/order/cmd/rpc/order"
	"home-nest/app/payment/cmd/api/internal/svc"
	"home-nest/app/payment/cmd/api/internal/types"
	"home-nest/app/payment/cmd/rpc/payment"
	"home-nest/app/payment/model"
	"home-nest/app/usercenter/cmd/rpc/usercenter"
	usercenterModel "home-nest/app/usercenter/model"
	"home-nest/pkg/ctxdata"
	"home-nest/pkg/xerr"
	"io"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrWxPayError = xerr.NewErrMsg("wechat pay fail")

type ThirdPaymentwxPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThirdPaymentwxPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThirdPaymentwxPayLogic {
	return &ThirdPaymentwxPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThirdPaymentwxPayLogic) ThirdPaymentwxPay(req *types.ThirdPaymentWxPayReq) (*types.ThirdPaymentWxPayResp, error) {
	var totalPrice int64   // Total amount paid for current order(cent)
	var description string // Current Payment Description.

	switch req.ServiceType {
	case model.ThirdPaymentServiceTypeHomestayOrder:

		homestayTotalPrice, homestayDescription, err := l.getPayHomestayPriceDescription(req.OrderSn)
		if err != nil {
			return nil, errors.Wrapf(ErrWxPayError, "getPayHomestayPriceDescription err : %v req: %+v", err, req)
		}
		totalPrice = homestayTotalPrice
		description = homestayDescription

	default:
		return nil, errors.Wrapf(xerr.NewErrMsg("Payment for this business type is not supported"), "Payment for this business type is not supported req: %+v", req)
	}

	// Create WechatPay pre-processing orders
	prepayRsp, err := l.createWxPrePayOrder(req.ServiceType, req.OrderSn, totalPrice, description)
	if err != nil {
		return nil, err
	}

	b, _ := io.ReadAll(prepayRsp.Body)
	var Res struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
		Data    struct {
			Status string `json:"status"`
		} `json:"data"`
	}
	err = json.Unmarshal(b, &Res)
	if err != nil {
		return nil, err
	}

	return &types.ThirdPaymentWxPayResp{
		Status: Res.Data.Status,
	}, nil
}

// Get the price and description information of the current order of the paid B&B
func (l *ThirdPaymentwxPayLogic) createWxPrePayOrder(serviceType, orderSn string, totalPrice int64, description string) (*http.Response, error) {

	// 1、get user openId
	userId := ctxdata.GetUidFromCtx(l.ctx)
	userResp, err := l.svcCtx.UsercenterRpc.GetUserAuthByUserId(l.ctx, &usercenter.GetUserAuthByUserIdReq{
		UserId:   userId,
		AuthType: usercenterModel.UserAuthTypeSystem,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrWxPayError, "Get user wechat openid err : %v , userId: %d , orderSn:%s", err, userId, orderSn)
	}
	if userResp.UserAuth == nil || userResp.UserAuth.Id == 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("Get user wechat openid fail，Please pay before authorization by weChat"), "Get user WeChat openid does not exist  userId: %d , orderSn:%s", userId, orderSn)
	}
	//openId := userResp.UserAuth.AuthKey

	// 2、create local third payment record
	createPaymentResp, err := l.svcCtx.PaymentRpc.CreatePayment(l.ctx, &payment.CreatePaymentReq{
		UserId:      userId,
		PayModel:    model.ThirdPaymentPayModelWechatPay,
		PayTotal:    totalPrice,
		OrderSn:     orderSn,
		ServiceType: serviceType,
	})
	if err != nil || createPaymentResp.Sn == "" {
		return nil, errors.Wrapf(ErrWxPayError,
			"create local third payment record fail : err: %v , userId: %d,totalPrice: %d , orderSn: %s",
			err, userId, totalPrice, orderSn)
	}

	body := struct {
		Description string `json:"description"`
		PaymentSn   string `json:"paymentSn"`
		NotifyUrl   string `json:"notify_url"`
		TotalPrice  int64  `json:"total_price"`
		UserId      int64  `json:"user_id"`
	}{
		Description: description,
		PaymentSn:   createPaymentResp.Sn,
		NotifyUrl:   l.svcCtx.CustomCallback.NotifyUrl,
		TotalPrice:  totalPrice,
		UserId:      userId,
	}
	b, _ := json.Marshal(body)
	r := bytes.NewReader(b)
	resp, err := http.Post("http://127.0.0.1:8383/custom-callback", "application/json", r)
	if err != nil {
		return nil, errors.Wrapf(ErrWxPayError, "Failed to initiate payment pre-order err : %v , userId: %d , orderSn:%s", err, userId, orderSn)
	}
	return resp, nil
}

// Get the price and description information of the current order of the paid B&B
func (l *ThirdPaymentwxPayLogic) getPayHomestayPriceDescription(orderSn string) (int64, string, error) {

	description := "homestay pay"

	// get user openid
	resp, err := l.svcCtx.OrderRpc.HomestayOrderDetail(l.ctx, &order.HomestayOrderDetailReq{
		Sn: orderSn,
	})
	if err != nil {
		return 0, description, errors.Wrapf(ErrWxPayError, "OrderRpc.HomestayOrderDetail err: %v, orderSn: %s", err, orderSn)
	}
	if resp.HomestayOrder == nil || resp.HomestayOrder.Id == 0 {
		return 0, description, errors.Wrapf(xerr.NewErrMsg("order no exists"), "WeChat payment order does not exist orderSn : %s", orderSn)
	}

	return resp.HomestayOrder.OrderTotalPrice, description, nil
}
