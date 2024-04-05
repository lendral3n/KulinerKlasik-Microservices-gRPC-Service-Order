package main

import (
	"fmt"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Order/pkg/client"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Order/pkg/config"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Order/pkg/db"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Order/pkg/pb"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Order/pkg/services"

	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(
		c.DB_USERNAME,
		c.DB_PASSWORD,
		c.DB_HOSTNAME,
		c.DB_PORT,
		c.DB_NAME,
	)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	menuSvc := client.InitMenuServiceClient(c.MenuSvcUrl)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Order Svc on", c.Port)

	s := services.Server{
		H:       h,
		MenuSVc: menuSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
