package user

import (
	"home-nest/pkg/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"home-nest/app/usercenter/cmd/api/internal/logic/user"
	"home-nest/app/usercenter/cmd/api/internal/svc"
	"home-nest/app/usercenter/cmd/api/internal/types"
)

// get user info
func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req)
		result.HttpResult(r, w, resp, err)
	}
}
