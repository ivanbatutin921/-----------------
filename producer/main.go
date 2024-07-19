package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

func main() {





	
	// producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, nil)
	// if err != nil {
	// 	fmt.Println("продюсер не создался", err)
	// }
	// defer producer.Close()

	// resp := &sarama.ProducerMessage{
	// 	Topic: "XYU",
	// 	Value: sarama.StringEncoder("hello world"),
	// }

	// producer.Input() <- resp

	// select {
	// case err := <-producer.Errors():
	// 	fmt.Println("ошибка", err)
	// default:
	// 	fmt.Println("сообщение отправлено")
	// }
	
}
