package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"helloworld/api/internal/config"
	"helloworld/rpc/greeter"
)

type ServiceContext struct {
	Config                  config.Config
	HelloWorldGreeterClient greeter.Greeter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		HelloWorldGreeterClient: greeter.NewGreeter(
			zrpc.MustNewClient(c.HelloWorld),
		),
	}
}
