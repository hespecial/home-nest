package logic

import (
	"context"
	"home-nest/app/payment/cmd/custom_callback/internal/svc"
	"home-nest/app/payment/cmd/custom_callback/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandleCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleCallbackLogic {
	return &HandleCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleCallbackLogic) HandleCallback(req *types.CustomCallbackReq) (resp *types.CustomCallbackResp, err error) {
	l.svcCtx.UserPayment.Store(req.UserId, req)
	return &types.CustomCallbackResp{
		Status: "success",
	}, nil
}
