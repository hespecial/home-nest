package homestayBussiness_

import (
	"context"
	"github.com/jinzhu/copier"
	"home-nest/app/travel/cmd/rpc/travel"

	"home-nest/app/travel/cmd/api/internal/svc"
	"home-nest/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBussinessDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// boss detail
func NewHomestayBussinessDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBussinessDetailLogic {
	return &HomestayBussinessDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBussinessDetailLogic) HomestayBussinessDetail(req *types.HomestayBussinessDetailReq) (resp *types.HomestayBussinessDetailResp, err error) {
	detail, err := l.svcCtx.TravelRpc.HomestayBusinessDetail(l.ctx, &travel.HomestayBusinessDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	var boss types.HomestayBusinessBoss
	_ = copier.Copy(&boss, detail.Boss)

	return &types.HomestayBussinessDetailResp{
		Boss: boss,
	}, nil
}
