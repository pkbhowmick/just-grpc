package main

import (
	"context"
	"log"
	"time"

	pb "github.com/pkbhowmick/just-grpc/client_streaming/proto"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to build tcp connection: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()

	c := pb.NewOrderManagementClient(conn)

	order1 := &pb.Order{
		Id:    "1",
		Name:  "Order1",
		Price: 10,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	id, err := c.AddOrder(ctx, order1)
	if err != nil {
		log.Fatalf("failed to add order %v", err)
	}
	log.Printf("Order with id %s is added successfully", id.Value)

	order2 := &pb.Order{
		Id:    "2",
		Name:  "Order2",
		Price: 20,
	}
	id, err = c.AddOrder(ctx, order2)
	if err != nil {
		log.Fatalf("failed to add order %v", err)
	}
	log.Printf("Order with id %s is added successfully", id.Value)

	// update order
	updatedOrders := []*pb.Order{
		{
			Id:    "1",
			Name:  "Updated Order 1",
			Price: 20,
		},
		{
			Id:    "2",
			Name:  "Updated Order 2",
			Price: 30,
		},
	}
	stream, err := c.UpdateOrders(ctx)
	if err != nil {
		log.Fatalf("failed to update orders: %v", err)
	}

	for _, order := range updatedOrders {
		err := stream.Send(order)
		if err != nil {
			log.Fatalf("failed to update orders: %v", err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to update orders: %v", err)
	}
	log.Printf("Server response: %s", resp)
}
