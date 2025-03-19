package homestay_

import (
	"context"
	"github.com/pkg/errors"
	"home-nest/app/travel/cmd/rpc/travel"
	"home-nest/pkg/tool"
	"home-nest/pkg/xerr"

	"home-nest/app/travel/cmd/api/internal/svc"
	"home-nest/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayDetailLogic {
	return &HomestayDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayDetailLogic) HomestayDetail(req *types.HomestayDetailReq) (resp *types.HomestayDetailResp, err error) {
	homestayResp, err := l.svcCtx.TravelRpc.HomestayDetail(l.ctx, &travel.HomestayDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("get homestay detail fail"), " get homestay detail err, id: %d, err: %v ", req.Id, err)
	}

	homestay := homestayResp.Homestay
	return &types.HomestayDetailResp{
		Homestay: types.Homestay{
			Id:                  homestay.Id,
			Title:               homestay.Title,
			SubTitle:            homestay.SubTitle,
			Banner:              homestay.Banner,
			Info:                homestay.Info,
			PeopleNum:           homestay.PeopleNum,
			HomestayBusinessId:  homestay.HomestayBusinessId,
			UserId:              homestay.UserId,
			RowState:            homestay.RowState,
			RowType:             homestay.RowType,
			FoodInfo:            homestay.FoodInfo,
			FoodPrice:           tool.Fen2Yuan(homestay.FoodPrice),
			HomestayPrice:       tool.Fen2Yuan(homestay.HomestayPrice),
			MarketHomestayPrice: tool.Fen2Yuan(homestay.MarketHomestayPrice),
		},
	}, nil
}
