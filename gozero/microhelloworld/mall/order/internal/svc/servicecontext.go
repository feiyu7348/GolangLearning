package svc

import (
	"GolangLearning/gozero/microhelloworld/mall/order/internal/config"
	"GolangLearning/gozero/microhelloworld/mall/user/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UesrRpc)),
	}
}
