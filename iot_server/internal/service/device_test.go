package service

import (
	"testing"
)

func Test_dealDeviceUploadMsg(t *testing.T) {
	msg := `{"topic":"device/upload/bme280","payload":"{\"Temperature\":25,\"Pressure\":10,\"Humidity\":10,\"Timestamp\":1703832944}","clientid":"weather_station_1"}`
	err := DealDeviceUploadMsg(msg)
	if err != nil {
		t.Error(err)
	}
}
