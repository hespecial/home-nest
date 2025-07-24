package homestay_

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"home-nest/app/travel/cmd/api/internal/svc"
	"home-nest/app/travel/cmd/api/internal/types"
	"home-nest/app/travel/cmd/rpc/travel"
	"home-nest/pkg/tool"
)

type BusinessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBusinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessListLogic {
	return &BusinessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BusinessListLogic) BusinessList(req *types.BusinessListReq) (*types.BusinessListResp, error) {
	resp, err := l.svcCtx.TravelRpc.BusinessList(l.ctx, &travel.BusinessListReq{
		LastId:             req.LastId,
		PageSize:           req.PageSize,
		HomestayBusinessId: req.HomestayBusinessId,
	})
	if err != nil {
		return nil, err
	}

	var list []types.Homestay
	if len(resp.List) > 0 {
		for _, homestay := range resp.List {
			var typeHomestay types.Homestay
			_ = copier.Copy(&typeHomestay, homestay)

			typeHomestay.FoodPrice = tool.Fen2Yuan(homestay.FoodPrice)
			typeHomestay.HomestayPrice = tool.Fen2Yuan(homestay.HomestayPrice)
			typeHomestay.MarketHomestayPrice = tool.Fen2Yuan(homestay.MarketHomestayPrice)
			list = append(list, typeHomestay)
		}
	}

	return &types.BusinessListResp{
		List: list,
	}, nil
}
