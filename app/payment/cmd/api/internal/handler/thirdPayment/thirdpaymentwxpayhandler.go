package thirdPayment

import (
	"home-nest/pkg/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"home-nest/app/payment/cmd/api/internal/logic/thirdPayment"
	"home-nest/app/payment/cmd/api/internal/svc"
	"home-nest/app/payment/cmd/api/internal/types"
)

// third payment：wechat pay
func ThirdPaymentwxPayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ThirdPaymentWxPayReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := thirdPayment.NewThirdPaymentwxPayLogic(r.Context(), svcCtx)
		resp, err := l.ThirdPaymentwxPay(&req)
		result.HttpResult(r, w, resp, err)
	}
}
