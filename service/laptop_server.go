package service

import (
	"context"
	"fmt"

	uuid2 "github.com/google/uuid"
	pb "github.com/pkbhowmick/just-grpc/proto"
)

type LaptopServer struct {
	Store LaptopStore
}

func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
}

func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {

	laptop := req.GetLaptop()
	uuid, err := uuid2.Parse(laptop.Id)
	if err != nil {

	}
	fmt.Println(uuid)
	res := &pb.CreateLaptopResponse{Id: uuid.String()}

	// need to save the laptop in the laptop store
	// wip

	return res, nil
}
