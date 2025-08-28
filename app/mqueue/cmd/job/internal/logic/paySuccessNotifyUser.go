package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"home-nest/app/mqueue/cmd/job/internal/svc"
	"home-nest/app/mqueue/cmd/job/jobtype"
	"home-nest/app/usercenter/cmd/rpc/usercenter"
	usercenterModel "home-nest/app/usercenter/model"
	"home-nest/pkg/xerr"
)

var ErrPaySuccessNotifyFail = xerr.NewErrMsg("pay success notify user fail")

// PaySuccessNotifyUserHandler pay success notify user
type PaySuccessNotifyUserHandler struct {
	svcCtx *svc.ServiceContext
}

func NewPaySuccessNotifyUserHandler(svcCtx *svc.ServiceContext) *PaySuccessNotifyUserHandler {
	return &PaySuccessNotifyUserHandler{
		svcCtx: svcCtx,
	}
}

func (l *PaySuccessNotifyUserHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {

	var p jobtype.PaySuccessNotifyUserPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(ErrPaySuccessNotifyFail, "PaySuccessNotifyUserHandler payload err:%v, payLoad:%+v", err, t.Payload())
	}

	// 1、get user openid
	usercenterResp, err := l.svcCtx.UsercenterRpc.GetUserAuthByUserId(ctx, &usercenter.GetUserAuthByUserIdReq{
		UserId:   p.Order.UserId,
		AuthType: usercenterModel.UserAuthTypeSystem,
	})
	if err != nil {
		return errors.Wrapf(ErrPaySuccessNotifyFail, "pay success notify user fail, rpc get user err:%v , orderSn:%s , userId:%d", err, p.Order.Sn, p.Order.UserId)
	}
	if usercenterResp.UserAuth == nil || len(usercenterResp.UserAuth.AuthKey) == 0 {
		return errors.Wrapf(ErrPaySuccessNotifyFail, "pay success notify user , user no exists err:%v , orderSn:%s , userId:%d", err, p.Order.Sn, p.Order.UserId)
	}

	// 2、send notify todo
	fmt.Println("success notify user, user_id: ", usercenterResp.UserAuth.UserId)

	return nil
}
