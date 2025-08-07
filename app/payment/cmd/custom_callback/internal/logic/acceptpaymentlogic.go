package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"home-nest/app/payment/cmd/custom_callback/internal/svc"
	"home-nest/app/payment/cmd/custom_callback/internal/types"
	"home-nest/pkg/tool"
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
		return nil, err
	}

	respBody, _ := io.ReadAll(resp.Body)
	var Res struct {
		ReturnCode string `json:"return_code"`
	}
	err = json.Unmarshal(respBody, &Res)
	if err != nil {
		return nil, err
	}
	l.svcCtx.UserPayment.Delete(req.UserId)
	return &types.AcceptPaymentResp{
		Status: Res.ReturnCode,
	}, nil
}
