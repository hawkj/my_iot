package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hawkj/my_iot/common/error"
	"github.com/hawkj/my_iot/common/function"
	"github.com/hawkj/my_iot/raspi/config"
	"github.com/hawkj/my_iot/raspi/job/job_handler"
	"github.com/hawkj/my_iot/raspi/pkg/common"

	"log"
	"os"
	"time"
)

var handlerMap = map[string]jobhandler.JobHandler{
	"test":        jobhandler.Test,
	"bme280":      jobhandler.Bme280,
	"bme280_mock": jobhandler.Bme280Mock,
}

func main() {
	configFile := os.Getenv("RASPI_SERVER_CONFIG")
	c := config.GetConfig(configFile)
	if c.AppEnv.Name == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 设置时区为8区
	time.Local = time.FixedZone("UTC+8", 8*60*60)

	ctx := context.Background()

	g := &common.Global{
		Config: c,
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
