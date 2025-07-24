package homestay_

import (
	"context"
	"github.com/jinzhu/copier"
	"home-nest/app/travel/cmd/api/internal/svc"
	"home-nest/app/travel/cmd/api/internal/types"
	"home-nest/app/travel/cmd/rpc/travel"
	"home-nest/pkg/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type GuessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGuessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GuessListLogic {
	return &GuessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GuessListLogic) GuessList(_ *types.GuessListReq) (resp *types.GuessListResp, err error) {
	list, err := l.svcCtx.TravelRpc.GuessList(l.ctx, &travel.GuessListReq{})
	if err != nil {
		return nil, err
	}

	var res []types.Homestay
	if len(list.List) > 0 {
		for _, homestay := range list.List {
			var typeHomestay types.Homestay
			_ = copier.Copy(&typeHomestay, homestay)

			typeHomestay.FoodPrice = tool.Fen2Yuan(homestay.FoodPrice)
			typeHomestay.HomestayPrice = tool.Fen2Yuan(homestay.HomestayPrice)
			typeHomestay.MarketHomestayPrice = tool.Fen2Yuan(homestay.MarketHomestayPrice)
			res = append(res, typeHomestay)
		}
	}

	return &types.GuessListResp{
		List: res,
	}, nil
}
