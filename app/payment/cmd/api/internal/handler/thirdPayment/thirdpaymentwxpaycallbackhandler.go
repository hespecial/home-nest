package thirdPayment

import (
	"home-nest/pkg/result"
	"net/http"

	"home-nest/app/payment/cmd/api/internal/logic/thirdPayment"
	"home-nest/app/payment/cmd/api/internal/svc"
)

// third payment：wechat pay callback
func ThirdPaymentWxPayCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := thirdPayment.NewThirdPaymentWxPayCallbackLogic(r.Context(), svcCtx)
		resp, err := l.ThirdPaymentWxPayCallback(w, r)
		result.HttpResult(r, w, resp, err)
	}
}
