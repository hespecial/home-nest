package homestayBussiness_

import (
	"context"
	"github.com/jinzhu/copier"
	"home-nest/app/travel/cmd/rpc/travel"

	"home-nest/app/travel/cmd/api/internal/svc"
	"home-nest/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodBossLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodBossLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodBossLogic {
	return &GoodBossLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodBossLogic) GoodBoss(_ *types.GoodBossReq) (resp *types.GoodBossResp, err error) {
	list, err := l.svcCtx.TravelRpc.GoodBoss(l.ctx, &travel.GoodBossReq{})
	if err != nil {
		return nil, err
	}

	var res []types.HomestayBusinessBoss
	if len(list.List) > 0 {
		for _, v := range list.List {
			var boss types.HomestayBusinessBoss
			_ = copier.Copy(&boss, v)
			res = append(res, boss)
		}
	}

	return &types.GoodBossResp{
		List: res,
	}, nil
}
