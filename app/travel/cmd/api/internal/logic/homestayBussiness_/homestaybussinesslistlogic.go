package homestayBussiness_

import (
	"context"
	"github.com/jinzhu/copier"
	"home-nest/app/travel/cmd/rpc/travel"

	"home-nest/app/travel/cmd/api/internal/svc"
	"home-nest/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBussinessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// business list
func NewHomestayBussinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBussinessListLogic {
	return &HomestayBussinessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBussinessListLogic) HomestayBussinessList(req *types.HomestayBussinessListReq) (*types.HomestayBussinessListResp, error) {
	list, err := l.svcCtx.TravelRpc.HomestayBusinessList(l.ctx, &travel.HomestayBusinessListReq{
		LastId:   req.LastId,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var resp []types.HomestayBusinessListInfo

	if len(list.List) > 0 {
		for _, item := range list.List {
			var typeHomestayBusinessListInfo types.HomestayBusinessListInfo
			_ = copier.Copy(&typeHomestayBusinessListInfo, item)

			resp = append(resp, typeHomestayBusinessListInfo)
		}
	}

	return &types.HomestayBussinessListResp{
		List: resp,
	}, nil
}
