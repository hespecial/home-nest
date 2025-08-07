package logic

import (
	"context"
	"github.com/pkg/errors"
	"home-nest/app/payment/model"
	"home-nest/pkg/globalkey"
	"home-nest/pkg/uniqueid"
	"home-nest/pkg/xerr"
	"time"

	"home-nest/app/payment/cmd/rpc/internal/svc"
	"home-nest/app/payment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePaymentLogic {
	return &CreatePaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建微信支付预处理订单
func (l *CreatePaymentLogic) CreatePayment(in *pb.CreatePaymentReq) (*pb.CreatePaymentResp, error) {
	data := new(model.ThirdPayment)
	data.Sn = uniqueid.GenSn(uniqueid.SnPrefixThirdPayment)
	data.UserId = in.UserId
	data.PayMode = in.PayModel
	data.PayTotal = in.PayTotal
	data.OrderSn = in.OrderSn
	data.ServiceType = model.ThirdPaymentServiceTypeHomestayOrder
	data.DeleteTime = time.Unix(0, 0)
	data.DelState = globalkey.DelStateNo
	data.PayTime = time.Unix(0, 0)

	_, err := l.svcCtx.ThirdPaymentModel.Insert(l.ctx, data)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "create wechat pay prepayorder db insert fail , err:%v ,data : %+v  ", err, data)
	}

	return &pb.CreatePaymentResp{
		Sn: data.Sn,
	}, nil
}
