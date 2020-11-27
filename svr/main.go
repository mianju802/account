package main

import (
	"context"
	"github.com/mianju802/lib/apollo"
	"github.com/mianju802/protocol/service/account"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/patrickmn/go-cache"
	"time"
)

type AccountService struct {
}

func (a AccountService) AccountRegister(ctx context.Context, req *account.AccountRegisterReq, rsp *account.AccountRegisterRsp) error {
	if req.UserName != "victor" || req.Passwd != "victor" {
		rsp.Code = -1
		rsp.Message = "用户名或密码不正确"
	}
	return nil
}

func main() {
	apo := &apollo.Apollo{LocalCache: cache.New(5*time.Minute, 10*time.Minute)}
	go func() {
		apo.ReadApolloConfig("consul")
	}()
	var (
		consulRegSvr registry.Registry
	)
	for {
		if consulAdd, ok := apo.LocalCache.Get("consulAdd"); ok {
			consulRegSvr = consul.NewRegistry(func(options *registry.Options) {
				options.Addrs = []string{
					consulAdd.(string),
				}
			})
			break
		}
	}
	service := micro.NewService(
		micro.Registry(consulRegSvr),
		micro.Name("micro.service.account"),
	)
	service.Init()
	if err := account.RegisterAccountServiceHandler(service.Server(), new(AccountService)); err != nil {
		panic(err)
	}
	if err := service.Run(); err != nil {
		panic(err)
	}
}
