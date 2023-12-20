package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hawkj/my_iot/common/constants"
	"github.com/hawkj/my_iot/common/error"
	"github.com/hawkj/my_iot/common/pkg/cache"
	"github.com/hawkj/my_iot/common/struct"
	"strings"
	"time"
)

func DealDeviceUploadMsg(ctx context.Context, redisClient *redis.Client, message string) error {
	uploadMsg := commonstruct.DeviceUploadMessage{}
	err := json.Unmarshal([]byte(message), &uploadMsg)
	if err != nil {
		return errors.Join(errors.New("[DealDeviceUploadMsg Unmarshal]"), err)
	}
	// 获取设备ID
	deviceCode, err := getDeviceCode(uploadMsg.Topic)
	if err != nil {
		return errors.Join(errors.New("[DealDeviceUploadMsg getDeviceCode]"), err)
	}
	mqttMessage := commonstruct.MqttMessage{}
	err = json.Unmarshal([]byte(uploadMsg.Payload), &mqttMessage)
	if err != nil {
		return errors.Join(errors.New("[DealDeviceUploadMsg MqttMessage]"), err)
	}

	if mqttMessage.MsgType == commoncons.MqttMsgTypeDeviceData {
		jsonData, err := json.Marshal(mqttMessage.Data)
		if err != nil {
			return errors.Join(errors.New("[DealDeviceUploadMsg json.Marshal(mqttMessage.Data)]"), err)
		}
		err = UpdateDeviceData(ctx, redisClient, deviceCode, uploadMsg.ClientID, string(jsonData))
		if err != nil {
			return errors.Join(errors.New("[DealDeviceUploadMsg UpdateDeviceData]"), err)
		}
	} else {
		return errors.New(commonerr.ErrEmqMsgType.ErrorMsg)
	}
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

func UpdateDeviceData(ctx context.Context, redisClient *redis.Client, deviceCode string, siteID string, data interface{}) error {
	//获取设备缓存key
	deviceCacheKey := commoncache.GetDeviceCacheKey(siteID, deviceCode)
	err := redisClient.HSet(ctx, deviceCacheKey, commoncache.H_F_Device_DeviceData, data).Err()
	if err != nil {
		return err
	}
	_, err = redisClient.Expire(ctx, deviceCacheKey, time.Hour*24*365).Result()
	if err != nil {
		return errors.Join(err, errors.New("[UpdateDeviceData redisClient.Expire]"))
	}
	return nil
}

func GetDeviceData(ctx context.Context, redisClient *redis.Client, deviceCode string, siteID string) (string, error) {
	value, err := redisClient.HGet(context.Background(), commoncache.GetDeviceCacheKey(siteID, deviceCode), commoncache.H_F_Device_DeviceData).Result()
	if err != nil {
		return "", err
	}
	return value, err
}
