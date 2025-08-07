package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"home-nest/app/order/cmd/rpc/order"
	"home-nest/app/payment/cmd/api/internal/config"
	"home-nest/app/payment/cmd/rpc/payment"
	"home-nest/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config config.Config

	CustomCallback struct {
		NotifyUrl string
	}

	PaymentRpc    payment.Payment
	OrderRpc      order.Order
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:         c,
		CustomCallback: struct{ NotifyUrl string }{NotifyUrl: c.CustomCallback.NotifyUrl},
		PaymentRpc:     payment.NewPayment(zrpc.MustNewClient(c.PaymentRpcConf)),
		OrderRpc:       order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		UsercenterRpc:  usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
