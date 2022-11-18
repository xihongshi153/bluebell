package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/pkg/logger"
	"bluebell/pkg/snowflake"
	"bluebell/router"
	"bluebell/setting"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 1.读取配置
	if err := setting.Init(); err != nil {
		log.Println("Init  Setting fail err: ", err)
		return
	}
	// return
	// 2.配置logger
	if err := logger.Init(setting.Conf.LogConfig); err != nil {
		log.Println("Init logger Setting fail err: ", err)
		return
	}
	defer zap.L().Sync() // flushing any buffered log  entries.
	zap.L().Info("logger init success")
	// 3.配置mysql
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		log.Println("Init mysql Setting fail err: ", err)
		return
	}
	defer mysql.Close()
	//4.配置redis
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		log.Println("Init reids Setting fail err: ", err)
		return
	}
	defer redis.Close()
	if err := snowflake.Init(setting.Conf.StartTime, int64(setting.Conf.MachineId)); err != nil {
		zap.L().Debug(fmt.Sprint("Init snowflake fail err: ", err.Error()))
	}
	//5.设置路由
	gin.SetMode(setting.Conf.Mode)

	e, err := router.SetUp()
	if err != nil {
		log.Println("Init  router fail err: ", err)
		return
	}
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Conf.Port),
		Handler: e,
	}

	// 6. 启动运行
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Debug("main.go ListenAndServe ", zap.String("err:", err.Error()))
		}
	}()
	// 优雅关机
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit //等待信号
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("shundown err: ", err)
	}
	log.Println("Server exiting")

}
