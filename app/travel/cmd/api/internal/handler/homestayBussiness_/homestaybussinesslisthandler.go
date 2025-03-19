package homestayBussiness_

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"home-nest/app/travel/cmd/api/internal/logic/homestayBussiness_"
	"home-nest/app/travel/cmd/api/internal/svc"
	"home-nest/app/travel/cmd/api/internal/types"
)

// business list
func HomestayBussinessListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayBussinessListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := homestayBussiness_.NewHomestayBussinessListLogic(r.Context(), svcCtx)
		resp, err := l.HomestayBussinessList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
