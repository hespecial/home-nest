package homestay_

import (
	"home-nest/pkg/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"home-nest/app/travel/cmd/api/internal/logic/homestay_"
	"home-nest/app/travel/cmd/api/internal/svc"
	"home-nest/app/travel/cmd/api/internal/types"
)

// boss's all homestay room
func BusinessListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BusinessListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestay_.NewBusinessListLogic(r.Context(), svcCtx)
		resp, err := l.BusinessList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
