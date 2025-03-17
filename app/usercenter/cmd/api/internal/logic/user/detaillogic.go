package user

import (
	"context"
	"home-nest/app/usercenter/cmd/rpc/usercenter"
	"home-nest/pkg/ctxdata"

	"home-nest/app/usercenter/cmd/api/internal/svc"
	"home-nest/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(_ *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	userInfoResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResp{
		UserInfo: types.User{
			Id:       userInfoResp.User.Id,
			Mobile:   userInfoResp.User.Mobile,
			Nickname: userInfoResp.User.Nickname,
			Sex:      userInfoResp.User.Sex,
			Avatar:   userInfoResp.User.Avatar,
			Info:     userInfoResp.User.Info,
		},
	}, nil
}
