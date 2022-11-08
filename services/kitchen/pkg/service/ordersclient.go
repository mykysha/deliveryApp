package service

import (
	"github.com/nndergunov/deliveryApp/services/order/api/v1/grpclogic/pb"
)

// OrdersClient interface shows needed signature for the Order Client.
type OrdersClient interface {
	GetIncompleteOrders(id int) (*pb.OrderResponseList, error)
}
