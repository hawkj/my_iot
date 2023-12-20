package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hawkj/my_iot/common/error"
	"github.com/hawkj/my_iot/common/function"
	commoncache "github.com/hawkj/my_iot/common/pkg/cache"
	"github.com/hawkj/my_iot/iot_server/config"
	"github.com/hawkj/my_iot/iot_server/job/job_handler"
	"github.com/hawkj/my_iot/iot_server/pkg/common"

	"log"
	"os"
	"time"
)

var handlerMap = map[string]jobhandler.JobHandler{
	"test":                   jobhandler.Test,
	"device_upload_consumer": jobhandler.DeviceUploadConsumer,
}

func main() {
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
		Config: c,
		Redis:  redisClient,
	}

	var params string
	var job string

	flag.StringVar(&job, "job", "", "Name of the job to run")
	flag.StringVar(&params, "params", "{}", "json of params")
	flag.Parse()

	if !commonfunc.IsValidJSON(params) {
		panic(commonerr.ErrParamsJson.ErrorMsg + " input is: " + params)
	}
	ctx = context.WithValue(ctx, "params", params)
	handler, ok := handlerMap[job]
	if !ok {
		panic(fmt.Sprintf("job name not exists: %s\n", job))
	}

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic("无法加载时区:" + err.Error())
	}
	log.Println(fmt.Sprintf("Job: %s start at %s", job, time.Now().In(location).Format("2006-01-02 15:04:05")))
	handler(ctx, g)
}
