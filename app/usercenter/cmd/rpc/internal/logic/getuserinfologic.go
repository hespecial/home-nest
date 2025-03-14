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

var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetUserInfo find user db err , id:%d , err:%v", in.Id, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "id:%d", in.Id)
	}

	return &usercenter.GetUserInfoResp{
		User: &usercenter.User{
			Id:       user.Id,
			Mobile:   user.Mobile,
			Nickname: user.Nickname,
			Sex:      user.Sex,
			Avatar:   user.Avatar,
			Info:     user.Info,
		},
	}, nil
}
