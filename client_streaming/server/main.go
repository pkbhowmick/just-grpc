package main

import (
	"context"
	"io"
	"log"
	"net"
	"sync"

	"github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/pkbhowmick/just-grpc/client_streaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

var (
	locker sync.Mutex
)

type server struct {
	orderMap map[string]*pb.Order
	pb.UnimplementedOrderManagementServer
}

func NewServer() *server {
	return &server{
		orderMap: make(map[string]*pb.Order),
	}
}

func (s *server) AddOrder(ctx context.Context, in *pb.Order) (*wrappers.StringValue, error) {
	_, found := s.orderMap[in.Id]
	if found {
		return nil, status.New(codes.AlreadyExists, "order with given id is already existed").Err()
	}
	locker.Lock()
	s.orderMap[in.Id] = in
	locker.Unlock()
	return &wrappers.StringValue{Value: in.Id}, nil
}

func (s *server) UpdateOrders(ordersServer pb.OrderManagement_UpdateOrdersServer) error {
	orderStr := "Updated orders id:"
	for {
		order, err := ordersServer.Recv()
		if err == io.EOF {
			return ordersServer.SendAndClose(&wrappers.StringValue{Value: orderStr})
		} else if err != nil {
			return status.New(codes.Internal, "failed to receive msg from client").Err()
		}
		locker.Lock()
		s.orderMap[order.Id] = order
		locker.Unlock()
		log.Printf("Updated order with id: %s", order.Id)
		orderStr += " " + order.Id + ","
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, NewServer())
	log.Println("Starting grpc server on port ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
