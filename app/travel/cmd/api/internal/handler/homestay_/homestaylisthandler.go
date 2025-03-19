package homestay_

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"home-nest/app/travel/cmd/api/internal/logic/homestay_"
	"home-nest/app/travel/cmd/api/internal/svc"
	"home-nest/app/travel/cmd/api/internal/types"
)

// homestay room list
func HomestayListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := homestay_.NewHomestayListLogic(r.Context(), svcCtx)
		resp, err := l.HomestayList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
