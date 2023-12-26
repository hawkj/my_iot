package pkg

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	brokerAddress = "localhost:9092"
	topic         = "my-topic"
)

func produceMessage() {
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

func consumeMessages() {
	// Create and configure the Reader directly
	reader := kafka.Reader{
		Brokers:  []string{brokerAddress},
		GroupID:  "my-consumer-group",
		Topic:    topic,
		MaxBytes: 10e6, // 10MB
	}

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

func main() {
	// Start a goroutine to produce messages
	go produceMessage()

	// Start a goroutine to consume messages
	go consumeMessages()

	// Wait for some time to allow producers and consumers to execute
	time.Sleep(10 * time.Second)
}
