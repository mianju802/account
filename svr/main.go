package main

import (
	"context"
	"github.com/mianju802/protocol/service/account"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
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
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"172.16.68.131:30769",
		}
	})
	service := micro.NewService(
		micro.Registry(consulReg),
		micro.Name("micro.service.account"),
	)
	service.Init()
	if err := account.RegisterAccountServiceHandler(service.Server(), new(AccountService));err != nil {
		panic(err)
	}
	if err := service.Run();err != nil {
		panic(err)
	}
}
