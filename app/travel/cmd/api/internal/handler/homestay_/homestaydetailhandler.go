package homestay_

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"home-nest/app/travel/cmd/api/internal/logic/homestay_"
	"home-nest/app/travel/cmd/api/internal/svc"
	"home-nest/app/travel/cmd/api/internal/types"
)

// homestay room detail
func HomestayDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := homestay_.NewHomestayDetailLogic(r.Context(), svcCtx)
		resp, err := l.HomestayDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
