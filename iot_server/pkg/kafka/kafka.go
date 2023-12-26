package pkg

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func KafkaProducer(brokerAddress, topic string, message kafka.Message) error {
	// Create and configure the Writer directly
	writer := kafka.Writer{
		Addr:     kafka.TCP(brokerAddress),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{}, // Use the LeastBytes balancer
	}

	defer writer.Close()

	// Send a message
	//message := kafka.Message{
	//	Key:   []byte("Key-1"),
	//	Value: []byte("Hello, Kafka!"),
	//}

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
	})
}
