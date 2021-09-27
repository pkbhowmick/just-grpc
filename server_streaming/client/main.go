package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/pkbhowmick/just-grpc/server_streaming/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

const (
	addr = "localhost:50051"
)

func main()  {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't run grpc server")
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	c := pb.NewOrderManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	// add a order
	order1ID, err := c.AddOrder(ctx, &pb.Order{
		Id:          "1",
		Items:       []string{"rice","oil","cheese"},
		Price:       12.50,
		Destination: "something",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Order1 ID: ", order1ID.Value)

	// add another order
	order2ID, err := c.AddOrder(ctx, &pb.Order{
		Id:          "1",
		Items:       []string{"rice-corn","oil","cheese"},
		Price:       12.50,
		Destination: "something",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Order2 ID: ", order2ID.Value)

	stream, err := c.SearchOrders(ctx, &wrappers.StringValue{Value: "rice"})
	if err != nil {
		log.Fatalf("error creating stream: %v", err)
	}

	for {
		order, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf(" error in receving message from server: %v", err)
		}
		log.Println("searched order: ", order)
	}
}
