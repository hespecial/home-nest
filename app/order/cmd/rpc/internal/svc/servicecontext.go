package svc

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"home-nest/app/order/cmd/rpc/internal/config"
	"home-nest/app/order/model"
	"home-nest/app/travel/cmd/rpc/travel"
)

type ServiceContext struct {
	Config             config.Config
	AsynqClient        *asynq.Client
	TravelRpc          travel.Travel
	HomestayOrderModel model.HomestayOrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		AsynqClient:        newAsynqClient(c),
		TravelRpc:          travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
		HomestayOrderModel: model.NewHomestayOrderModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
