package pkg

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func KafkaProducer(brokerAddress, topic string) {
	// Create and configure the Writer directly
	writer := kafka.Writer{
		Addr:     kafka.TCP(brokerAddress),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{}, // Use the LeastBytes balancer
	}

	defer writer.Close()

	// Send a message
	message := kafka.Message{
		Key:   []byte("Key-1"),
		Value: []byte("Hello, Kafka!"),
	}

	err := writer.WriteMessages(context.Background(), message)
	if err != nil {
		log.Fatal("failed to write message:", err)
	}

	fmt.Println("Message sent successfully!")
}

func KafkaConsumer(brokerAddress string, groupID string, topic string) {
	// Create and configure the Reader directly
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerAddress},
		GroupID:  groupID,
		Topic:    topic,
		MaxBytes: 10e6, // 10MB
	})

	defer reader.Close()

	// Receive messages
	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read message:", err)
		}

		fmt.Printf("Received message: %s\n", string(message.Value))
	}
}
