package initialize

import (
	"fmt"
	"log"

	"github.com/buiminhhoat/go-ecommerce-backend-api/global"
	"github.com/segmentio/kafka-go"
)

var KafkaProducer *kafka.Writer

func InitKafka() {
	fmt.Println("InitKafka is running")
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "otp-auth-topic",
		Balancer: &kafka.LeastBytes{},
	}
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatalf("Failed to close kafka producer: %d", err)
	}
}
