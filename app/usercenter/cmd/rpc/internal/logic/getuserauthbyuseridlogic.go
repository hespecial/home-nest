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

type GetUserAuthByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByUserIdLogic {
	return &GetUserAuthByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByUserIdLogic) GetUserAuthByUserId(in *pb.GetUserAuthByUserIdReq) (*pb.GetUserAuthyUserIdResp, error) {
	userAuth, err := l.svcCtx.UserAuthModel.FindOneByUserIdAuthType(l.ctx, in.UserId, in.AuthType)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user auth fail"), "err : %v , in : %+v", err, in)
	}

	return &pb.GetUserAuthyUserIdResp{
		UserAuth: &usercenter.UserAuth{
			Id:       userAuth.Id,
			UserId:   userAuth.UserId,
			AuthType: userAuth.AuthType,
			AuthKey:  userAuth.AuthKey,
		},
	}, nil
}
