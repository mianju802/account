package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mianju802/protocol/service/account"
	microClient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	"net/http"
)

var (
	client account.AccountService
)

func initMicro() {
	if err := cmd.Init(); err != nil {
		panic(err)
	}
	client = account.NewAccountService("micro.service.account", microClient.DefaultClient)
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
