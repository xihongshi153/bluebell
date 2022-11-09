package router

import (
	"bluebell/ctroller.go"
	"bluebell/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp() (e *gin.Engine, err error) {
	e = gin.New()
	e.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 1.注册
	e.POST("/register", ctroller.RegisterHandler)
	// 2.登录
	e.POST("/login", ctroller.LoginHandler)
	e.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "test success")
	})
	// 登录
	return
}
