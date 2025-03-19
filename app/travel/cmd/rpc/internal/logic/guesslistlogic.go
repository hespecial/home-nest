package logic

import (
	"context"
	"github.com/pkg/errors"
	"home-nest/pkg/xerr"

	"home-nest/app/travel/cmd/rpc/internal/svc"
	"home-nest/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GuessListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGuessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GuessListLogic {
	return &GuessListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GuessListLogic) GuessList(_ *pb.GuessListReq) (*pb.GuessListResp, error) {
	list, err := l.svcCtx.HomestayModel.FindPageListByIdDESC(l.ctx, l.svcCtx.HomestayModel.SelectBuilder(), 0, 5)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GuessList db err : %v", err)
	}

	var resp []*pb.Homestay
	for _, homestay := range list {
		resp = append(resp, &pb.Homestay{
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
			FoodPrice:           homestay.FoodPrice,
			HomestayPrice:       homestay.HomestayPrice,
			MarketHomestayPrice: homestay.MarketHomestayPrice,
		})
	}

	return &pb.GuessListResp{
		List: resp,
	}, nil
}
