package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"home-nest/app/travel/model"
	"home-nest/app/usercenter/cmd/rpc/usercenter"
	"home-nest/pkg/xerr"

	"home-nest/app/travel/cmd/rpc/internal/svc"
	"home-nest/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrHomestayBusinessNotExistsError = xerr.NewErrCodeMsg(xerr.ServerCommonError, "民宿店铺不存在")

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

	var boss pb.HomestayBusinessBoss
	userResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		Id: homestayBusiness.UserId,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("get boss info fail"), "get boss info fail ,  userId : %d ,err:%v", homestayBusiness.UserId, err)
	}
	if userResp.User != nil && userResp.User.Id > 0 {
		_ = copier.Copy(&boss, userResp.User)
	}

	return &pb.HomestayBusinessDetailResp{
		Boss: &boss,
	}, nil
}
