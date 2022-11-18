package controller

import (
	"bluebell/dao/models"
	"bluebell/logic"
	"bluebell/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterHandler(c *gin.Context) {
	// 1.参数校验
	var registerParam models.RegisterParam
	if err := c.ShouldBind(&registerParam); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "unexpected param err: " + err.Error(),
		})
		zap.L().Error("unexpected register param ", zap.String("err", err.Error()))
		return
	}
	// 2.loggic
	err := logic.Register(registerParam.UserName, registerParam.Passwrod, registerParam.Email)
	// 3.返回值
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	c.JSON(http.StatusOK, "success123")
}
func LoginHandler(ctx *gin.Context) {
	var loginParam models.LoginParam
	if err := ctx.ShouldBind(&loginParam); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "unexpected param err: " + err.Error(),
		})
		zap.L().Error("unexpected param ", zap.String("err", err.Error()))
	}
	success, token, err := logic.Login(loginParam.UserName, loginParam.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	if success {
		response.ResponseJSON(ctx, response.CodeSuccess, map[string]any{"token": token})
	} else {
		ctx.JSON(500, gin.H{
			"err": "something wrong,please contact with Administrator,err is nil but fail login",
		})
	}
}
