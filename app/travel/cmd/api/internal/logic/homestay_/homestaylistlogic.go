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

type HomestayListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayListLogic {
	return &HomestayListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayListLogic) HomestayList(req *types.HomestayListReq) (resp *types.HomestayListResp, err error) {
	homestayListResp, err := l.svcCtx.TravelRpc.HomestayList(l.ctx, &travel.HomestayListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("get homestay list fail"), "get homestay list err: %v ", err)
	}

	var list []types.Homestay
	for _, homestay := range homestayListResp.List {
		list = append(list, types.Homestay{
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
		})
	}

	return &types.HomestayListResp{
		List: list,
	}, nil
}
