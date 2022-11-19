package controller

import (
	"bluebell/logic"
	"bluebell/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func VoteHandler(ctx *gin.Context) {
	pidStr := ctx.Param("pid")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.S().Error("invalid param ", pidStr, " err: ", err.Error())
		response.ResponseJSON(ctx, response.CodeParamInvalid, nil)
		return
	}
	uidStr := ctx.Keys["userid"].(string)
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		zap.S().Error("invalid param ", uidStr, " err: ", err.Error())
		response.ResponseJSON(ctx, response.CodeParamInvalid, nil)
		return
	}
	err = logic.MakeVoteLogic(int(pid), int(uid))
	if err != nil {
		zap.S().Error("server error err: ", err.Error())
		response.ResponseJSON(ctx, response.CodeServerError, nil)
		return
	}
	response.ResponseJSON(ctx, response.CodeSuccess, nil)
}
func DeVoteHandler(ctx *gin.Context) {
	pidStr := ctx.Param("pid")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.S().Error("invalid param ", pidStr, " err: ", err.Error())
		response.ResponseJSON(ctx, response.CodeParamInvalid, nil)
		return
	}
	uidStr := ctx.Keys["userid"].(string)
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		zap.S().Error("invalid param ", uidStr, " err: ", err.Error())
		response.ResponseJSON(ctx, response.CodeParamInvalid, nil)
		return
	}
	err = logic.DeleteVoteLogic(int(pid), int(uid))
	if err != nil {
		zap.S().Error("server error err: ", err.Error())
		response.ResponseJSON(ctx, response.CodeServerError, nil)
		return
	}
	response.ResponseJSON(ctx, response.CodeSuccess, nil)
}
