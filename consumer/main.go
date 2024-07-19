package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		fmt.Println("консьюмер не создался", err)
	}

	//подписка на тему
	topic := "XYU"
	partitionConsumer, err := consumer.ConsumePartition(topic, int32(0), sarama.OffsetNewest)
	if err != nil {
		fmt.Println("консьюмер не подписался", err)
	}

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Println("полученное сообщение:", msg.Topic)
		case err := <-partitionConsumer.Errors():
			fmt.Println("ошибка:", err)

		}
	}

}
