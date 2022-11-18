package controller

import (
	"bluebell/logic"
	"bluebell/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CommunityInfoHandler
// 查询 社区信息
func CommunityInfoHandler(ctx *gin.Context) {
	data, err := logic.SelectCommunityInfo()
	if err != nil {
		zap.S().Debugf("controller CommunityInfoHandler err: %s", err.Error())
		response.ResponseJSON(ctx, response.CodeServerError, map[string]any{"err": err.Error()})
		return
	}
	response.ResponseJSON(ctx, response.CodeSuccess, map[string]any{"communityInfo": data})
}

// 查询 社区的介绍
func CommunityIntroductionHandler(ctx *gin.Context) {
	cidstr := ctx.Param("cid")
	if len(cidstr) == 0 {
		zap.S().Debug("controller CommunityIntroductionHandler id is \"\" ")
		response.ResponseJSON(ctx, response.CodeParamInvalid, map[string]any{"err": "参数为\""})
		return
	}
	cid, err := strconv.ParseInt(cidstr, 10, 64)
	if err != nil {
		zap.S().Debug("controller CommunityIntroductionHandler param invalid")
		response.ResponseJSON(ctx, response.CodeParamInvalid, map[string]any{"err": err.Error()})
		return
	}
	intro, err := logic.SelectCommunityIntroBycId(int(cid))
	if err != nil {
		zap.S().Debug("controller CommunityIntroductionHandler err:%s", err.Error())
		response.ResponseJSON(ctx, response.CodeServerError, map[string]any{"err": err.Error()})
		return
	}
	response.ResponseJSON(ctx, response.CodeSuccess, map[string]any{"introduction": intro})
}
