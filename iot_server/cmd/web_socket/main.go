package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hawkj/my_iot/common/pkg/cache"
	"github.com/hawkj/my_iot/iot_server/config"
	"github.com/hawkj/my_iot/iot_server/internal/handler"
	"github.com/hawkj/my_iot/iot_server/internal/middleware"
	"github.com/hawkj/my_iot/iot_server/pkg/common"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const ServiceName = "iot_server"

func main() {
	servierInfo := config.GetServerInfo(ServiceName)
	// 创建Gin引擎
	r := gin.Default()
	srv := &http.Server{
		Addr:    servierInfo.Addr,
		Handler: r,
	}
	configFile := os.Getenv("IOT_SERVER_CONFIG")
	c := config.GetConfig(configFile)
	if c.AppEnv.Name == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	// 设置时区为8区
	time.Local = time.FixedZone("UTC+8", 8*60*60)
	ctx := context.Background()
	redisClient, err := commoncache.NewRedis(ctx, c.Redis.Address)
	if err != nil {
		panic(err)
	}
	g := &common.Global{
		Redis:  redisClient,
		Config: c,
	}

	router(r, g, c)

	// 启动服务器
	log.Printf("start listen %s\n", servierInfo.Addr)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				// 记录日志或进行其他处理
				log.Println("Recovered:", r)
			}
		}()
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号（如按下Ctrl+C）来优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 设置超时时间
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭 HTTP 服务器
	log.Println("shutting down server...")
	if err := srv.Shutdown(timeoutCtx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Println("server exiting")
}

func router(r *gin.Engine, g *common.Global, c *config.Config) {
	//基础路由组
	baseGroup := r.Group("/api", middleware.CommonContext(g, c))

	baseGroup.GET("/device/info", middleware.CommonHandlerWrapper(handler.DeviceInfo))
}
