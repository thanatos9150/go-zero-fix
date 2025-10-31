package svc

import "github.com/thanatos9150/go-zero-fix/tools/goctl/example/rpc/hello/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
