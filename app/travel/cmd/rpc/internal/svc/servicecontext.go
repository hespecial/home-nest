package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"home-nest/app/travel/cmd/rpc/internal/config"
	"home-nest/app/travel/model"
	"home-nest/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config config.Config

	HomestayModel         model.HomestayModel
	HomestayActivityModel model.HomestayActivityModel
	HomestayBusinessModel model.HomestayBusinessModel

	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:                c,
		HomestayModel:         model.NewHomestayModel(sqlConn, c.Cache),
		HomestayActivityModel: model.NewHomestayActivityModel(sqlConn, c.Cache),
		HomestayBusinessModel: model.NewHomestayBusinessModel(sqlConn, c.Cache),
		UsercenterRpc:         usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
