package logic

import (
	"context"
	"github.com/pkg/errors"
	"home-nest/app/travel/model"
	"home-nest/pkg/xerr"

	"home-nest/app/travel/cmd/rpc/internal/svc"
	"home-nest/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrHomeStayNotExistsError = xerr.NewErrMsg("房源不存在")

type HomestayDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomestayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayDetailLogic {
	return &HomestayDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomestayDetailLogic) HomestayDetail(in *pb.HomestayDetailReq) (*pb.HomestayDetailResp, error) {
	homestay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, in.Id)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), " HomestayDetail db err , id : %d ", in.Id)
	}
	if homestay == nil {
		return nil, errors.Wrapf(ErrHomeStayNotExistsError, "id:%d", in.Id)
	}

	return &pb.HomestayDetailResp{
		Homestay: &pb.Homestay{
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
		},
	}, nil
}
