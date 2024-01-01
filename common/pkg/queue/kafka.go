package queue

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

func KafkaProducer(brokerAddress, topic string, message kafka.Message) error {
	// Create and configure the Writer directly
	writer := kafka.Writer{
		Addr:     kafka.TCP(brokerAddress),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{}, // Use the LeastBytes balancer
	}

	defer writer.Close()

	err := writer.WriteMessages(context.Background(), message)
	if err != nil {
		return err
	}

	return nil
}

func GetKafkaReader(brokerAddress string, topic string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerAddress},
		GroupID:  groupID,
		Topic:    topic,
		MaxBytes: 10e6, // 10MB
		Dialer: &kafka.Dialer{
			Timeout: 5 * time.Second, // 设置超时报错时间为5秒，根据你的需求调整
		},
	})
}
