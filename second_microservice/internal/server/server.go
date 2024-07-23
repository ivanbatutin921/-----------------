package server

import (
	"encoding/json"
	"fmt"

	pb "root/second_microservice/internal/proto"

	"github.com/IBM/sarama"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) OutputUser(stream pb.UserService_OutputUserServer) error {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		fmt.Println("консьюмер не создался", err)
		return err
	}
	defer consumer.Close()

	//подписка на тему
	topic := "register_users"
	partitionConsumer, err := consumer.ConsumePartition(topic, int32(0), sarama.OffsetNewest)
	if err != nil {
		fmt.Println("консьюмер не подписался", err)
		return err
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var user pb.User
			err = json.Unmarshal(msg.Value, &user)
			if err!= nil {
				fmt.Println("ошибка декодирования", err)
			} else {
				jsonBytes, err := json.MarshalIndent(&user, "", "  ")
				if err!= nil {
					fmt.Println("ошибка форматирования JSON", err)
				} else {
					fmt.Println(string(jsonBytes))
				}
			}
		case err := <-partitionConsumer.Errors():
			fmt.Println("ошибка:", err)
		}
	}
}
