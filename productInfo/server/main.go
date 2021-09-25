package main

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/pkbhowmick/just-grpc/productInfo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	port = ":50051"
)


type server struct {
	pb.UnimplementedProductInfoServer
	productMap map[string] *pb.Product
}

func NewServer() *server{
	s := &server{}
	s.productMap = make(map[string]*pb.Product)
	return s
}

func (s *server)AddProduct(ctx context.Context,in *pb.Product) (*pb.ProductID, error) {
	id := uuid.New().String()
	s.productMap[id] = in
	return &pb.ProductID{Value: id}, status.New(codes.OK, "").Err()
}

func (s *server)GetProduct(ctx context.Context,in *pb.ProductID) (*pb.Product, error) {
	val, found := s.productMap[in.Value]
	if !found {
		return nil, status.Errorf(codes.NotFound, "Product doesn't exist,", in.Value)
	}
	return val, status.New(codes.OK, "").Err()
}


func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, NewServer())
	log.Println("Starting gRPC listener on port "+ port)
	if err:= s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}