package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"home-nest/app/order/cmd/api/internal/config"
	"home-nest/app/order/cmd/rpc/order"
	"home-nest/app/payment/cmd/rpc/payment"
	"home-nest/app/travel/cmd/rpc/travel"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc   order.Order
	PaymentRpc payment.Payment
	TravelRpc  travel.Travel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		OrderRpc:   order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		PaymentRpc: payment.NewPayment(zrpc.MustNewClient(c.PaymentRpcConf)),
		TravelRpc:  travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
	}
}
