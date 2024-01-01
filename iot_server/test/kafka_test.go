package pkg

import (
	"context"
	"fmt"
	commoncons "github.com/hawkj/my_iot/common/constants"
	commonpkg "github.com/hawkj/my_iot/common/pkg/queue"
	"github.com/hawkj/my_iot/iot_server/config"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"testing"
)

// cd /path/to/your/package
// go test -run Test_kafka
func Test_kafka(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	configFile := currentDir + "/../config/iot_server_conf.yaml"
	c := config.GetConfig(configFile)
	testTopic := commoncons.KafkaTopicUploadDevice

	msg := kafka.Message{
		Key:   []byte("test"),
		Value: []byte("test_msg"),
	}
	err = commonpkg.KafkaProducer(c.Kafka.BrokerAddress, testTopic, msg)
	if err != nil {
		t.Error(err)
	}

	reader := commonpkg.GetKafkaReader(c.Kafka.BrokerAddress, testTopic, "default_group")
	defer reader.Close()
	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read message:", err)
		}
		fmt.Printf("Received message: %s\n", string(message.Value))
	}
}
