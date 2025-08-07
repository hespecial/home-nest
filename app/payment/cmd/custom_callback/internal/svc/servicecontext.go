package svc

import (
	"home-nest/app/payment/cmd/custom_callback/internal/config"
	"sync"
)

type ServiceContext struct {
	Config      config.Config
	UserPayment sync.Map
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserPayment: sync.Map{},
	}
}
