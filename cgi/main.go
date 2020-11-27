package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mianju802/lib/apollo"
	"github.com/mianju802/protocol/service/account"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

var (
	client account.AccountService
)

func initMicro() {
	apo := &apollo.Apollo{LocalCache: cache.New(5*time.Minute, 10*time.Minute)}
	go func() {
		apo.ReadApolloConfig("consul")
	}()
	var (
		consulRegCli registry.Registry
	)
	for {
		if consulAdd, ok := apo.LocalCache.Get("consulAdd"); ok {
			consulRegCli = consul.NewRegistry(func(options *registry.Options) {
				options.Addrs = []string{
					consulAdd.(string),
				}
			})
			break
		} else {
			log.Debug("apollo 未读取到 consul 配置 。。。。")
		}
	}
	service := micro.NewService(
		micro.Name("micro.client.account"),
		micro.Registry(consulRegCli),
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
