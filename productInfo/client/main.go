package main

import (
	"context"
	pb "github.com/pkbhowmick/just-grpc/productInfo/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	addr = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect to server: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.AddProduct(ctx, &pb.Product{
		Id:    "1",
		Name:  "Apple IPhone 13",
		Price: 300,
	})
	if err != nil {
		log.Fatalf("couldn't add product: %v",err)
	}
	log.Printf("Product id: %s added successfully", resp.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: resp.Value})
	if err != nil {
		log.Printf("couldn't get product: %v", err)
	}
	log.Println("Product: ", product.String())
}
