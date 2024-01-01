package service

import (
	"encoding/json"
	"fmt"
	"github.com/hawkj/my_iot/common/struct"
	"strings"
)

func DealDeviceUploadMsg(message string) (string, string, error) {
	uploadMsg := commonstruct.DeviceUploadMessage{}
	err := json.Unmarshal([]byte(message), &uploadMsg)
	if err != nil {
		return "", "", err
	}
	// 获取设备ID
	deviceCode, err := getDeviceCode(uploadMsg.Topic)
	if err != nil {
		return "", "", err
	}

	return deviceCode, uploadMsg.ClientID, nil
}

func getDeviceCode(EmqTopic string) (string, error) {
	parts := strings.Split(EmqTopic, "/")

	if len(parts) >= 2 {
		return parts[len(parts)-1], nil
	} else {
		return "", fmt.Errorf("invalid input string from [getDeviceCode]")
	}
}
