package user

import (
	"context"
	"github.com/pkg/errors"
	"home-nest/app/usercenter/cmd/rpc/usercenter"
	"home-nest/app/usercenter/model"

	"home-nest/app/usercenter/cmd/api/internal/svc"
	"home-nest/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	registerResp, err := l.svcCtx.UsercenterRpc.Register(l.ctx, &usercenter.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
		AuthKey:  req.Mobile,
		AuthType: model.UserAuthTypeSystem,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.RegisterResp{
		AccessToken:  registerResp.AccessToken,
		AccessExpire: registerResp.AccessExpire,
		RefreshAfter: registerResp.RefreshAfter,
	}, nil
}
