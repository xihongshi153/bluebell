package controller

import (
	"bluebell/dao/models"
	"bluebell/logic"
	"bluebell/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(ctx *gin.Context) {
	// 1. 参数校验
	var param models.CreatePostParam
	if err := ctx.BindJSON(&param); err != nil {
		response.ResponseJSON(ctx, response.CodeParamInvalid, map[string]any{"err": err.Error()})
		zap.S().Error("postController.go CreatePostHandler err: %s", err.Error())
		return
	}
	// 2. logic
	success, err := logic.CreatePost(param)
	// 3. 结果处理
	if err != nil {
		response.ResponseJSON(ctx, response.CodeServerError, map[string]any{"err": err.Error()})
		zap.S().Error("postController.go CreatePostHandler logic.CreatePost(param)   err: %s", err.Error())
		return
	}
	if !success {
		response.ResponseJSON(ctx, response.CodeServerError, map[string]any{"err": "服务器内部错误 err为nil 但是没有成功添加"})
		zap.S().Error("postController.go CreatePostHandler logic.CreatePost(param)  服务器内部错误 err为nil 但是没有成功添加")

		return
	}
	response.ResponseJSON(ctx, response.CodeSuccess, nil)
}
func GetPostDetailByPIdHandler(ctx *gin.Context) {
	pidStr := ctx.Param("pid")
	if len(pidStr) == 0 {
		response.ResponseJSON(ctx, response.CodeParamInvalid, map[string]any{"err": "pid为空"})
		zap.S().Error("postController.go GetPostByPIdHandler err: pid为空")
		return
	}
	id, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		response.ResponseJSON(ctx, response.CodeParamInvalid, map[string]any{"err": err.Error()})
		zap.S().Error("postController.go GetPostByPIdHandler param invalid err: %s", err.Error())
		return
	}
	data, err := logic.GetPostDetailById(int(id))
	if err != nil {
		response.ResponseJSON(ctx, response.CodeServerError, map[string]any{"err": err.Error()})
		zap.S().Error("postController.go func:GetPostByPIdHandler logic.GetPostById(id) err: %s", err.Error())
		return
	}
	response.ResponseJSON(ctx, response.CodeSuccess, map[string]any{"data": data})
}
func GetPostPageHandler(ctx *gin.Context) {
	pageStr := ctx.PostForm("page")
	limitStr := ctx.PostForm("limit")
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		response.ResponseJSON(ctx, response.CodeParamInvalid, map[string]any{"err": "page参数错误"})
		zap.S().Error("param page invalid err: ", err.Error())
		return
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		response.ResponseJSON(ctx, response.CodeParamInvalid, map[string]any{"err": "limit参数错误"})
		zap.S().Error("param page invalid err: ", err.Error())
		return
	}
	data, err := logic.GetPostPage(int(page), int(limit))
	if err != nil {
		response.ResponseJSON(ctx, response.CodeServerError, map[string]any{"err": err.Error()})
		zap.S().Error("data, err := logic.GetPostPage(int(page), int(limit)) err: ", err.Error())
		return
	}
	response.ResponseJSON(ctx, response.CodeSuccess, map[string]any{"data": data})
}
