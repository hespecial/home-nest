package handler

import (
	"home-nest/pkg/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"home-nest/app/payment/cmd/custom_callback/internal/logic"
	"home-nest/app/payment/cmd/custom_callback/internal/svc"
	"home-nest/app/payment/cmd/custom_callback/internal/types"
)

func handleCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CustomCallbackReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewHandleCallbackLogic(r.Context(), svcCtx)
		resp, err := l.HandleCallback(&req)
		result.HttpResult(r, w, resp, err)
	}
}
