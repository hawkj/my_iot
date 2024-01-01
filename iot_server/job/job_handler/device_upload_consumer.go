package jobhandler

import (
	"context"
	"encoding/json"
	"fmt"
	commoncons "github.com/hawkj/my_iot/common/constants"
	"github.com/hawkj/my_iot/common/pkg/queue"
	commonstruct "github.com/hawkj/my_iot/common/struct"
	"github.com/hawkj/my_iot/iot_server/pkg/common"
	"log"
	"strings"
)

func DeviceUploadConsumer(ctx context.Context, g *common.Global) {
	reader := queue.GetKafkaReader(g.Config.Kafka.BrokerAddress, commoncons.KafkaTopicUploadDevice, commoncons.DefaultConsumerGroup)
	defer reader.Close()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read message:", err)
		}
		log.Printf("Received message: %s\n", string(message.Value))

	}
}

func dealDeviceUploadMsg(message string) error {
	uploadMsg := commonstruct.DeviceUploadMessage{}
	err := json.Unmarshal([]byte(message), &uploadMsg)
	if err != nil {
		return err
	}
	// 获取设备ID
	deviceCode, err := getDeviceCode(uploadMsg.Topic)
	if err != nil {
		return err
	}
	// 解析payLoad

	siteID := uploadMsg.ClientID
	fmt.Println(fmt.Sprintf("%+v", uploadMsg))
	fmt.Println(deviceCode, siteID)
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

//func parsePayload(deviceCode, payload string) (interface{}, error) {
//	var payloadData interface{}
//	switch deviceCode {
//	case commoncons.DeviceBME280:
//		bme280data := commonstruct.BME280{}
//		err := json.Unmarshal([]byte(payload), &bme280data)
//		if err != nil {
//			return nil, err
//		}
//		payloadData = bme280data
//	default:
//
//		return nil, errors.New("[parsePayload] unknown deviceCode")
//	}
//}
