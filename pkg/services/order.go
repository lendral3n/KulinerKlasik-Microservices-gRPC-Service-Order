package services

import (
	"context"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Order/pkg/client"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Order/pkg/db"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Order/pkg/models"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Order/pkg/pb"
	"net/http"
)

type Server struct {
	H       db.Handler
	MenuSVc client.MenuServiceClient
	pb.UnimplementedOrderServiceServer
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	menu, err := s.MenuSVc.FindOne(req.MenuId)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	} else if menu.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{
			Status: menu.Status,
			Error:  menu.Error,
		}, nil
	} else if menu.Data.Stock < req.Quantity {
		return &pb.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  "Stock too les",
		}, nil
	}

	order := models.Order{
		Price:  menu.Data.Price,
		MenuId: menu.Data.Id,
		UserId: req.UserId,
	}
	s.H.DB.Create(&order)

	res, err := s.MenuSVc.DecreasedStock(req.MenuId, order.Id)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error()}, nil
	} else if res.Status == http.StatusConflict {
		s.H.DB.Delete(&models.Order{}, order.Id)

		return &pb.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  res.Error}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id: order.Id,
	}, nil
}
