package homestay_

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"home-nest/app/travel/cmd/api/internal/logic/homestay_"
	"home-nest/app/travel/cmd/api/internal/svc"
	"home-nest/app/travel/cmd/api/internal/types"
)

// guess your favorite homestay room
func GuessListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GuessListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := homestay_.NewGuessListLogic(r.Context(), svcCtx)
		resp, err := l.GuessList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
