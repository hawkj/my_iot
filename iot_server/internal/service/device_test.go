package service

import (
	"context"
	commoncache "github.com/hawkj/my_iot/common/pkg/cache"
	"github.com/hawkj/my_iot/iot_server/config"
	"os"
	"testing"
)

func Test_dealDeviceUploadMsg(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	configFile := currentDir + "/../../config/iot_server_conf.yaml"
	c := config.GetConfig(configFile)
	msg := `{"topic":"device/upload/bme280","payload":"{\"msg_type\":\"device-data\",\"data\":{\"Temperature\":22,\"Pressure\":10,\"Humidity\":10,\"Timestamp\":1704177170}}","clientid":"weather_station_1"}`
	err = DealDeviceUploadMsg(context.Background(), commoncache.NewRedis(c.Redis.Address), msg)
	if err != nil {
		t.Error(err)
	}
}
