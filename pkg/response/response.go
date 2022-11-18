package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	{
		code: 200
		msg : ""
		data: map[string]interface{}
	}
*/

type response struct {
	Code ResCode        `json:"code"`
	Msg  string         `json:"msg"`
	Data map[string]any `json:"data"`
}

/*
	代替例如
	c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "auth格式错误",
	})
*/
// 自定义msg
func ResponseJSONM(ctx *gin.Context, code ResCode, msg string, data map[string]any) {
	ctx.JSON(http.StatusOK, response{Code: code, Msg: msg, Data: data})
}

// code有对应的msg
func ResponseJSON(ctx *gin.Context, code ResCode, data map[string]any) {
	ctx.JSON(http.StatusOK, response{Code: code, Msg: codeToString(code), Data: data})
}

type ResCode int64

const (
	CodeSuccess = 1000 + iota
	CodeParamInvalid
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeTokenEmpty   // Auth为空
	CodeTokenInvalid // Auth不合法
)

var ResCodeToErrString = map[ResCode]string{
	CodeSuccess:         "success",
	CodeParamInvalid:    "参数不合法",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "密码错误",
	CodeServerBusy:      "服务器忙碌",
	CodeTokenEmpty:      "Auth为空",
	CodeTokenInvalid:    "Auth不合法",
}

func codeToString(code ResCode) string {
	if s, ok := ResCodeToErrString[code]; ok {
		return s
	}
	return ""
}
