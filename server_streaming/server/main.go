package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	pb "github.com/pkbhowmick/just-grpc/server_streaming/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"sync"
)

const (
	port = ":50051"
)

var locker sync.Mutex

type server struct {
	OrderMap map[string]*pb.Order
	pb.UnimplementedOrderManagementServer
}

func NewServer() *server {
	s := &server{
		OrderMap: make(map[string]*pb.Order),
	}
	return s
}

func (s *server) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	for _, order := range s.OrderMap {
		for _, item := range order.Items {
			if strings.Contains(item, searchQuery.Value) {
				err := stream.Send(order)
				if err != nil {
					return fmt.Errorf("error sending message to stream: %v", err)
				}
			}
		}
	}
	return nil
}

func (s *server) AddOrder(ctx context.Context, in *pb.Order) (*wrappers.StringValue, error) {
	orderID := uuid.New().String()
	locker.Lock()
	s.OrderMap[orderID] = in
	locker.Unlock()
	return &wrappers.StringValue{Value: orderID}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, NewServer())
	log.Println("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
