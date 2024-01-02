package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hawkj/my_iot/common/struct"
	"strings"
)

func DealDeviceUploadMsg(ctx context.Context, redisClient *redis.Client, message string) error {
	uploadMsg := commonstruct.DeviceUploadMessage{}
	err := json.Unmarshal([]byte(message), &uploadMsg)
	if err != nil {
		return err
	}
	// 获取设备ID
	deviceCode, err := getDeviceCode(uploadMsg.Topic)
	if err != nil {
		return errors.Join(errors.New("[DealDeviceUploadMsg getDeviceCode]"), err)
	}
	fmt.Printf(uploadMsg.ClientID)
	return nil
}

func getDeviceCode(EmqTopic string) (string, error) {
	parts := strings.Split(EmqTopic, "/")
	if len(parts) >= 2 {
		return parts[len(parts)-1], nil
	} else {
		return "", fmt.Errorf("invalid input string from [getDeviceCode]")
	}
}
