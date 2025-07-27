package svc

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/zrpc"
	"home-nest/app/mqueue/cmd/job/internal/config"
	"home-nest/app/order/cmd/rpc/order"
	"home-nest/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server

	OrderRpc      order.Order
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		AsynqServer:   newAsynqServer(c),
		OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
