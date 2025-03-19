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

var ErrHomestayBusinessNotExistsError = xerr.NewErrMsg("民宿店铺不存在")

type HomestayBusinessDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomestayBusinessDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBusinessDetailLogic {
	return &HomestayBusinessDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomestayBusinessDetailLogic) HomestayBusinessDetail(in *pb.HomestayBusinessDetailReq) (*pb.HomestayBusinessDetailResp, error) {
	homestayBusiness, err := l.svcCtx.HomestayBusinessModel.FindOne(l.ctx, in.Id)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), " HomestayBussinessDetail  FindOne db fail ,id  : %d , err : %v", in.Id, err)
	}
	if homestayBusiness == nil {
		return nil, errors.Wrapf(ErrHomestayBusinessNotExistsError, "id:%d", in.Id)
	}

	return &pb.HomestayBusinessDetailResp{
		Boss: &pb.HomestayBusinessBoss{
			//Id:       0,
			UserId:   0,
			Nickname: "",
			Avatar:   "",
			Info:     "",
			//Rank:     0,
		},
	}, nil
}
