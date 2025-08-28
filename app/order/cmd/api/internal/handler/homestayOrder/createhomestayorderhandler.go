package homestayOrder

import (
	"home-nest/pkg/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"home-nest/app/order/cmd/api/internal/logic/homestayOrder"
	"home-nest/app/order/cmd/api/internal/svc"
	"home-nest/app/order/cmd/api/internal/types"
)

// 创建民宿订单
func CreateHomestayOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateHomestayOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayOrder.NewCreateHomestayOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateHomestayOrder(&req)
		result.HttpResult(r, w, resp, err)
	}
}
