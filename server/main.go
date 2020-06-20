package main

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/yenchieh/grpc-with-vue/server/gen/pb"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Server struct{
	Topics []*pb.Topic
}

func (s *Server) GetTopics(c context.Context, e *empty.Empty) (*pb.Topics, error) {
	c, cancel := context.WithTimeout(c, 1 * time.Second)
	defer cancel()

	return &pb.Topics{Topics: s.Topics}, nil
}

func (s *Server) VoteTopic(c context.Context, request *pb.RequestByID) (*pb.Topic, error) {
	c, cancel := context.WithTimeout(c, 1 * time.Second)
	defer cancel()

	var target *pb.Topic
	for _, topic := range s.Topics {
		if topic.Id == request.Id {
			target = topic
			break
		}
	}
	if target == nil {
		return nil, errors.New("no topic")
	}
	target.VoteCount++
	return target, nil
}

func main() {
	// Generate Topics
	server := Server{}
	server.Topics = append(server.Topics, &pb.Topic{
		Id:                   1,
		Message:              "You Like Go",
		VoteCount:            0,
	})
	server.Topics = append(server.Topics, &pb.Topic{
		Id:                   2,
		Message:              "You Like Vue",
		VoteCount:            0,
	})

	s := grpc.NewServer()
	pb.RegisterVoteServer(s, &server)
	lis, err := net.Listen("tcp", ":5110")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server on 5110")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
