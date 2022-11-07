// Package grpclogic implements a server using gRPC connection.
package grpclogic

import (
	"context"
	"fmt"

	"github.com/nndergunov/deliveryApp/services/order/api/v1/grpclogic/pb"
	"github.com/nndergunov/deliveryApp/services/order/pkg/domain"
	"github.com/nndergunov/deliveryApp/services/order/pkg/service"
	"google.golang.org/grpc"
)

type orderGRPC struct {
	pb.UnimplementedRestaurantServiceServer
	service service.AppService
}

// NewOrderRawGRPCServer returns new instance of raw order service gRPC server.
func NewOrderRawGRPCServer(service service.AppService) *grpc.Server {
	grpcLogic := new(orderGRPC)

	grpcLogic.service = service

	srv := grpc.NewServer()

	pb.RegisterRestaurantServiceServer(srv, grpcLogic)

	return srv
}

// GetOrder handles request of getting the order from the service.
func (r *orderGRPC) GetOrder(ctx context.Context, req *pb.Request) (*pb.OrderResponse, error) {
	order, err := r.service.ReturnOrder(int(req.GetID()))
	if err != nil {
		return nil, fmt.Errorf("getting order: %w", err)
	}

	items := make([]int32, len(order.OrderItems))

	for ind, item := range order.OrderItems {
		items[ind] = int32(item)
	}

	return &pb.OrderResponse{
		OrderID:      int32(order.OrderID),
		FromUserID:   int32(order.FromUserID),
		RestaurantID: int32(order.RestaurantID),
		OrderItems:   items,
		Status: &pb.OrderStatus{
			OrderID: int32(order.Status.OrderID),
			Status:  order.Status.Status,
		},
	}, nil
}

// GetOrderList handles request of getting the order from the service.
func (r *orderGRPC) GetOrderList(ctx context.Context, req *pb.Parameters) (*pb.OrderResponseList, error) {
	i := int(req.FromRestaurantID)
	param := domain.SearchParameters{
		FromRestaurantID: &i,
		Statuses:         req.Statuses,
		ExcludeStatuses:  req.ExcludeStatuses,
	}
	orderList, err := r.service.ReturnOrderList(param)
	if err != nil {
		return nil, fmt.Errorf("getting order list: %w", err)
	}

	var orderResponseList []*pb.OrderResponse
	for _, order := range orderList {
		orderResponse := &pb.OrderResponse{
			OrderID:      int32(order.OrderID),
			FromUserID:   int32(order.FromUserID),
			RestaurantID: int32(order.RestaurantID),
			OrderItems:   convertSliceToInt32(order.OrderItems),
			Status: &pb.OrderStatus{
				OrderID: int32(order.Status.OrderID),
				Status:  order.Status.Status,
			},
		}
		orderResponseList = append(orderResponseList, orderResponse)
	}

	return &pb.OrderResponseList{
		OrderResponseList: orderResponseList,
	}, nil
}

func convertSliceToInt32(list []int) []int32 {
	list32 := make([]int32, len(list))

	for i := range list32 {
		i32 := int32(i)
		list32 = append(list32, i32)
	}
	return list32
}
