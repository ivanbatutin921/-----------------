package server

import (
	"context"
	"encoding/json"
	"fmt"

	db "root/mk/internal/database"
	pb "root/mk/internal/proto"

	"github.com/IBM/sarama"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	err := db.DB.DB.Exec(`INSERT INTO users (login, password) VALUES ($1, $2)`, req.Login, req.Password)
	if err != nil {
		fmt.Println("ошибка", err.Error)
	}

	sendMessageToKafka(req)

	return &pb.User{
			Login:    req.Login,
			Password: req.Password},
		nil
}

func sendMessageToKafka(req *pb.User) {
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		fmt.Println("продюсер не создался", err.Error())
	}
	defer producer.Close()

	jsonBytes,err:= json.Marshal(req)
	if err != nil {
		fmt.Println("ошибка", err)
	}

	resp := &sarama.ProducerMessage{
		Topic: "register_users",
		Value: sarama.ByteEncoder(jsonBytes),
	}

	producer.Input() <- resp

	select {
	case err := <-producer.Errors():
		fmt.Println("ошибка", err)
	default:
		fmt.Println("сообщение отправлено")
	}
}
