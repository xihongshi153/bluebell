package router

import (
	"bluebell/logger"

	"github.com/gin-gonic/gin"
)

func SetUp() (e *gin.Engine, err error) {
	e = gin.New()
	e.Use(logger.GinLogger(), logger.GinRecovery(true))

	return
}
