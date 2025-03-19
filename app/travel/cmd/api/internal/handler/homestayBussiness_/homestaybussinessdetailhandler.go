package homestayBussiness_

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"home-nest/app/travel/cmd/api/internal/logic/homestayBussiness_"
	"home-nest/app/travel/cmd/api/internal/svc"
	"home-nest/app/travel/cmd/api/internal/types"
)

// boss detail
func HomestayBussinessDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayBussinessDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := homestayBussiness_.NewHomestayBussinessDetailLogic(r.Context(), svcCtx)
		resp, err := l.HomestayBussinessDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
