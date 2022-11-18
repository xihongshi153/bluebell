package router

import (
	"bluebell/ctroller.go"
	"bluebell/pkg/jwt"
	"bluebell/pkg/logger"
	"bluebell/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp() (e *gin.Engine, err error) {
	e = gin.New()
	e.Use(gin.Logger(), logger.GinLogger(), logger.GinRecovery(true))
	// 1.注册
	e.POST("/register", ctroller.RegisterHandler)
	// 2.登录
	e.POST("/login", ctroller.LoginHandler)
	e.POST("/islogin", jwt.JwtAuthMiddleware(), func(ctx *gin.Context) {
		response.ResponseJSON(ctx, response.CodeSuccess, nil)
	})
	e.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "test success")
	})
	// 登录
	return
}
