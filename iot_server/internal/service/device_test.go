package service

import (
	"context"
	"fmt"
	"github.com/hawkj/my_iot/common/pkg/cache"
	"github.com/hawkj/my_iot/iot_server/config"
	"os"
	"testing"
)

// go test -run Test_dealDeviceUploadMsg
func Test_dealDeviceUploadMsg(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	configFile := currentDir + "/../../config/iot_server_conf.yaml"
	c := config.GetConfig(configFile)
	msg := `{"topic":"device/upload/bme280","payload":"{\"msg_type\":\"device-data\",\"data\":{\"temperature\":25,\"pressure\":10,\"humidity\":10,\"timestamp\":1704191215}}","clientid":"weather_station_1"}
`

	redisClient, err := commoncache.NewRedis(context.Background(), c.Redis.Address)
	if err != nil {
		t.Error(err)
	}

	err = DealDeviceUploadMsg(context.Background(), redisClient, msg)
	if err != nil {
		t.Error(err)
	}

	value, err := redisClient.HGet(context.Background(), commoncache.GetDeviceCacheKey("weather_station_1", "bme280"), commoncache.H_F_Device_DeviceData).Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("----------------------------")
	fmt.Println(value)
	fmt.Println("----------------------------")
}
