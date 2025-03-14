package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"home-nest/app/usercenter/cmd/rpc/internal/config"
	"home-nest/app/usercenter/model"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	UserModel     model.UserModel
	UserAuthModel model.UserAuthModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		RedisClient:   redis.MustNewRedis(c.Redis.RedisConf),
		UserModel:     model.NewUserModel(conn, c.Cache),
		UserAuthModel: model.NewUserAuthModel(conn, c.Cache),
	}
}
