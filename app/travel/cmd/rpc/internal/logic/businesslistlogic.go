package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"home-nest/pkg/xerr"

	"home-nest/app/travel/cmd/rpc/internal/svc"
	"home-nest/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBusinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessListLogic {
	return &BusinessListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BusinessListLogic) BusinessList(in *pb.BusinessListReq) (*pb.BusinessListResp, error) {
	whereBuilder := l.svcCtx.HomestayModel.SelectBuilder().Where(squirrel.Eq{"homestay_business_id": in.HomestayBusinessId})
	list, err := l.svcCtx.HomestayModel.FindPageListByIdDESC(l.ctx, whereBuilder, in.LastId, in.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "HomestayBusinessId: %d ,err : %v", in.HomestayBusinessId, err)
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

	return &pb.BusinessListResp{
		List: resp,
	}, nil
}
