package client

import (
	"context"
	"fmt"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Order/pkg/pb"

	"google.golang.org/grpc"
)

type MenuServiceClient struct {
	Client pb.MenuServiceClient
}

func InitMenuServiceClient(url string) MenuServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := MenuServiceClient{
		Client: pb.NewMenuServiceClient(cc),
	}

	return c
}

func (c *MenuServiceClient) FindOne(menuId int64) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: menuId,
	}

	return c.Client.FindOne(context.Background(), req)
}

func (c *MenuServiceClient) DecreasedStock(menuId, orderId int64) (*pb.DecreaseStockResponse, error) {
	req := &pb.DecreaseStockRequest{
		Id:      menuId,
		OrderId: orderId,
	}

	return c.Client.DecreaseStock(context.Background(), req)
}
