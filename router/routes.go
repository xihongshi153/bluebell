package router

import (
	"bluebell/controller"
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
	e.POST("/register", controller.RegisterHandler)
	// 2.登录
	e.POST("/login", controller.LoginHandler)
	e.POST("/islogin", jwt.JwtAuthMiddleware(), func(ctx *gin.Context) {
		response.ResponseJSON(ctx, response.CodeSuccess, nil)
	})
	v1 := e.Group("/api/v1", jwt.JwtAuthMiddleware())
	{
		// 社区相关
		v1.GET("/community", controller.CommunityInfoHandler)
		// 根据id获取社区的详细信息c
		v1.GET("/communityIntroduction/:cid", controller.CommunityIntroductionHandler)
		// 创建post
		v1.POST("/createpost", controller.CreatePostHandler)
		// 查看某一个帖子
		v1.GET("/post/:pid", controller.GetPostDetailByPIdHandler)
		// 投票
		v1.POST("/doVote/:pid", controller.VoteHandler)
		v1.POST("/deVote/:pid", controller.DeVoteHandler)
	}
	// 分页展示帖子
	e.GET("/postpage", controller.GetPostPageHandler)
	e.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "test success")
	})
	// 登录
	return
}
