package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"home-nest/app/order/cmd/api/internal/config"
	"home-nest/app/travel/cmd/rpc/travel"
)

type ServiceContext struct {
	Config config.Config

	TravelRpc travel.Travel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		TravelRpc: travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
	}
}
