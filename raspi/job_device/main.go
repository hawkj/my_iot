package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hawkj/my_iot/raspi/pkg/common"

	"github.com/hawkj/my_iot/raspi/config"
	jobhandler "github.com/hawkj/my_iot/raspi/job_device/job_handler"

	"log"
	"os"
	"time"
)

var handlerMap = map[string]jobhandler.JobHandler{
	"test": jobhandler.Test,
}

func main() {
	configFile := os.Getenv("APP_REMIND_ME_CONFIG")
	c := config.GetConfig(configFile)
	if c.AppEnv.Name == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 设置时区为8区
	time.Local = time.FixedZone("UTC+8", 8*60*60)

	ctx := context.Background()

	g := &common.Global{}
	var params string
	var job string
	flag.StringVar(&job, "job", "", "Name of the job to run")
	flag.StringVar(&params, "params", "{}", "json of params")

	flag.Parse()
	if !common.IsValidJSON(params) {
		panic(common.ErrParamsJson.ErrorMsg + " input is: " + params)
	}
	ctx = context.WithValue(ctx, "params", params)
	handler, ok := handlerMap[job]
	if !ok {
		panic(fmt.Sprintf("job name not exists: %s\n", job))
	}

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Println("无法加载时区:" + err.Error())
		return
	}
	log.Println(fmt.Sprintf("Job: %s start at %s", job, time.Now().In(location).Format("2006-01-02 15:04:05")))
	handler(ctx, g)
}
