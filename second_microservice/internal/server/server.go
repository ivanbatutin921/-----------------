package server

import (
	"context"
	"encoding/json"
	"fmt"

	pb "root/second_microservice/internal/proto"

	"github.com/IBM/sarama"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) OutputUser(ctx context.Context, req *pb.User) (*pb.Empty, error) {

	receiveMessageFromKafka()

	return &pb.Empty{}, nil
}

func receiveMessageFromKafka() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		fmt.Println("консьюмер не создался", err)
	}

	//подписка на тему
	topic := "register_users"
	partitionConsumer, err := consumer.ConsumePartition(topic, int32(0), sarama.OffsetNewest)
	if err != nil {
		fmt.Println("консьюмер не подписался", err)
	}

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var user pb.User
			err = json.Unmarshal(msg.Value, &user)
			if err != nil {
				fmt.Println("ошибка декодирования", err)
			} else {
				fmt.Println("полученное сообщение:", &user)
			}
		case err := <-partitionConsumer.Errors():
			fmt.Println("ошибка:", err)

		}
	}
}
