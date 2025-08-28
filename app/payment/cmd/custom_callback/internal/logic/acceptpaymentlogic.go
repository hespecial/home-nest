package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"home-nest/app/payment/cmd/custom_callback/internal/svc"
	"home-nest/app/payment/cmd/custom_callback/internal/types"
	"home-nest/pkg/tool"
	"home-nest/pkg/xerr"
	"io"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type AcceptPaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAcceptPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AcceptPaymentLogic {
	return &AcceptPaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AcceptPaymentLogic) AcceptPayment(req *types.AcceptPaymentReq) (*types.AcceptPaymentResp, error) {
	value, _ := l.svcCtx.UserPayment.Load(req.UserId)
	data := value.(*types.CustomCallbackReq)
	body := struct {
		Description   string `json:"description"`
		PaymentSn     string `json:"paymentSn"`
		TotalPrice    int64  `json:"total_price"`
		TradeState    string `json:"trade_state"`
		TransactionId string `json:"transaction_id"`
		TradeType     string `json:"trade_type"`
		PayStatus     int    `json:"pay_status"`
	}{
		Description:   data.Description,
		PaymentSn:     data.PaymentSn,
		TotalPrice:    data.TotalPrice,
		TradeState:    "SUCCESS",
		TransactionId: tool.Krand(32, 3),
		TradeType:     "Normal",
		PayStatus:     0,
	}
	jsonStr, _ := json.Marshal(body)
	b := bytes.NewReader(jsonStr)
	resp, err := http.Post(data.NotifyUrl, "application/json", b)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ServerInternalError), "request notify url failed err: %v", err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ServerInternalError), "read response body failed err: %v", err)
	}
	var Res struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
		Data    struct {
			ReturnCode string `json:"return_code"`
		} `json:"data"`
	}
	err = json.Unmarshal(respBody, &Res)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ServerInternalError), "unmarshal response body failed err: %v", err)
	}
	l.svcCtx.UserPayment.Delete(req.UserId)
	return &types.AcceptPaymentResp{
		Status: Res.Data.ReturnCode,
	}, nil
}
