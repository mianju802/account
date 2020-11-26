package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mianju802/protocol/service/account"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"net/http"
)

var (
	client account.AccountService
)

func initMicro() {
	consulReq := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"172.16.68.131:30769",
		}
	})
	service := micro.NewService(
		micro.Name("micro.client.account"),
		micro.Registry(consulReq),
	)
	client = account.NewAccountService("micro.service.account", service.Client())
}

func newRouter() *gin.Engine {
	route := gin.Default()
	route.POST("/account/register", func(c *gin.Context) {
		userName := c.PostForm("username")
		passWd := c.PostForm("password")
		fmt.Println("============", userName, passWd)
		rsp, err := client.AccountRegister(context.TODO(), &account.AccountRegisterReq{
			UserName: userName,
			Passwd:   passWd,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    rsp.Code,
				"message": rsp.Message,
			})
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    rsp.Code,
			"message": rsp.Message,
		})
		return
	})
	return route
}

func main() {
	initMicro()
	r := newRouter()
	if err := r.Run("0.0.0.0:1234"); err != nil {
		panic(err)
	}
}
