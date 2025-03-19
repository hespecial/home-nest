package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"home-nest/app/travel/cmd/api/internal/config"
	"home-nest/app/travel/cmd/rpc/travel"
	"home-nest/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config config.Config

	UsercenterRpc usercenter.Usercenter
	TravelRpc     travel.Travel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		TravelRpc:     travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
	}
}
