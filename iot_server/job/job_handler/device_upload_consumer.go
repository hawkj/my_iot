package jobhandler

import (
	"context"
	commoncons "github.com/hawkj/my_iot/common/constants"
	"github.com/hawkj/my_iot/common/pkg/queue"
	"github.com/hawkj/my_iot/iot_server/internal/service"
	"github.com/hawkj/my_iot/iot_server/pkg/common"
	"log"
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
		err = service.DealDeviceUploadMsg(ctx, g.Redis, string(message.Value))
		if err != nil {
			log.Fatal("failed to read message:", err)
		}

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
