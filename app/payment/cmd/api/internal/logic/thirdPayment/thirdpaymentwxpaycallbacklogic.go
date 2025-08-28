package thirdPayment

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"home-nest/app/payment/cmd/rpc/payment"
	"home-nest/app/payment/model"
	"home-nest/pkg/xerr"
	"io"
	"net/http"
	"time"

	"home-nest/app/payment/cmd/api/internal/svc"
	"home-nest/app/payment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrPaymentCallbackError = xerr.NewErrMsg("wechat pay callback fail")

type ThirdPaymentWxPayCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThirdPaymentWxPayCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThirdPaymentWxPayCallbackLogic {
	return &ThirdPaymentWxPayCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThirdPaymentWxPayCallbackLogic) ThirdPaymentWxPayCallback(_ http.ResponseWriter, req *http.Request) (*types.ThirdPaymentWxPayCallbackResp, error) {
	body, err := io.ReadAll(req.Body)
	var data Transaction
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	returnCode := "SUCCESS"
	err = l.verifyAndUpdateState(&data)
	if err != nil {
		returnCode = "FAIL"
	}

	return &types.ThirdPaymentWxPayCallbackResp{
		ReturnCode: returnCode,
	}, err
}

type Transaction struct {
	Description   string `json:"description"`
	PaymentSn     string `json:"paymentSn"`
	TotalPrice    int64  `json:"total_price"`
	TradeState    string `json:"trade_state"`
	TransactionId string `json:"transaction_id"`
	TradeType     string `json:"trade_type"`
	PayStatus     int    `json:"pay_status"`
}

// Verify and update relevant flow data
func (l *ThirdPaymentWxPayCallbackLogic) verifyAndUpdateState(notifyTrasaction *Transaction) error {

	paymentResp, err := l.svcCtx.PaymentRpc.GetPaymentBySn(l.ctx, &payment.GetPaymentBySnReq{
		Sn: notifyTrasaction.PaymentSn,
	})
	if err != nil || paymentResp.PaymentDetail.Id == 0 {
		return errors.Wrapf(ErrPaymentCallbackError, "Failed to get payment flow record err:%v ,notifyTrasaction:%+v ", err, notifyTrasaction)
	}

	//比对金额
	notifyPayTotal := notifyTrasaction.TotalPrice
	if paymentResp.PaymentDetail.PayTotal != notifyPayTotal {
		return errors.Wrapf(ErrPaymentCallbackError, "Order amount exception  notifyPayTotal:%v , notifyTrasaction:%v ", notifyPayTotal, notifyTrasaction)
	}

	// Judgment status
	payStatus := l.getPayStatusByWXPayTradeState(notifyTrasaction.TradeState)
	if payStatus == model.ThirdPaymentPayTradeStateSuccess {
		//Payment Notification.

		if paymentResp.PaymentDetail.PayStatus != model.ThirdPaymentPayTradeStateWait {
			return nil
		}

		// Update the flow status.
		if _, err = l.svcCtx.PaymentRpc.UpdateTradeState(l.ctx, &payment.UpdateTradeStateReq{
			Sn:             notifyTrasaction.PaymentSn,
			TradeState:     notifyTrasaction.TradeState,
			TransactionId:  notifyTrasaction.TransactionId,
			TradeType:      notifyTrasaction.TradeType,
			TradeStateDesc: notifyTrasaction.Description,
			PayStatus:      l.getPayStatusByWXPayTradeState(notifyTrasaction.TradeState),
			PayTime:        time.Now().Unix(),
		}); err != nil {
			return errors.Wrapf(ErrPaymentCallbackError, "更新流水状态失败  err:%v , notifyTrasaction:%v ", err, notifyTrasaction)
		}

	} else if payStatus == model.ThirdPaymentPayTradeStateWait {
		//Refund notification @todo to be done later, not needed at this time
	}

	return nil

}

const (
	SUCCESS    = "SUCCESS"    //支付成功
	REFUND     = "REFUND"     //转入退款
	NOTPAY     = "NOTPAY"     //未支付
	CLOSED     = "CLOSED"     //已关闭
	REVOKED    = "REVOKED"    //已撤销（付款码支付）
	USERPAYING = "USERPAYING" //用户支付中（付款码支付）
	PAYERROR   = "PAYERROR"   //支付失败(其他原因，如银行返回失败)
)

func (l *ThirdPaymentWxPayCallbackLogic) getPayStatusByWXPayTradeState(wxPayTradeState string) int64 {

	switch wxPayTradeState {
	case SUCCESS: //支付成功
		return model.ThirdPaymentPayTradeStateSuccess
	case USERPAYING: //支付中
		return model.ThirdPaymentPayTradeStateWait
	case REFUND: //已退款
		return model.ThirdPaymentPayTradeStateWait
	default:
		return model.ThirdPaymentPayTradeStateFAIL
	}
}
