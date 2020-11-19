package main

import (
	"context"
	"github.com/mianju802/protocol/service/account"
	"github.com/micro/go-micro"
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
	service := micro.NewService(
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
