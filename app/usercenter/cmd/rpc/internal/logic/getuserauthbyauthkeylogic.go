package logic

import (
	"context"
	"github.com/pkg/errors"
	"home-nest/app/usercenter/cmd/rpc/usercenter"
	"home-nest/app/usercenter/model"
	"home-nest/pkg/xerr"

	"home-nest/app/usercenter/cmd/rpc/internal/svc"
	"home-nest/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByAuthKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByAuthKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByAuthKeyLogic {
	return &GetUserAuthByAuthKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByAuthKeyLogic) GetUserAuthByAuthKey(in *pb.GetUserAuthByAuthKeyReq) (*pb.GetUserAuthByAuthKeyResp, error) {
	userAuth, err := l.svcCtx.UserAuthModel.FindOneByAuthTypeAuthKey(l.ctx, in.AuthType, in.AuthKey)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user auth fail"), "err : %v , in : %+v", err, in)
	}

	return &pb.GetUserAuthByAuthKeyResp{
		UserAuth: &usercenter.UserAuth{
			Id:       userAuth.Id,
			UserId:   userAuth.UserId,
			AuthType: userAuth.AuthType,
			AuthKey:  userAuth.AuthKey,
		},
	}, nil
}
