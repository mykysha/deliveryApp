// Package ordersclient implements client that is used to send requests to the
// order service.
package ordersclient

import (
	"context"
	"fmt"
	"log"

	"github.com/nndergunov/deliveryApp/app/services/order/api/v1/grpclogic/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// OrdersClient is used to comunicate with the order service.
type OrdersClient struct {
	orderServiceBaseURL string
}

// NewOrdersClient creates new OrdersClient instance.
func NewOrdersClient(orderServiceBaseURL string) *OrdersClient {
	return &OrdersClient{orderServiceBaseURL: orderServiceBaseURL}
}

// GetIncompleteOrders returns all orders that are not marked as complete.
func (o OrdersClient) GetIncompleteOrders(id int) (*pb.OrderResponseList, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(o.orderServiceBaseURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	c := pb.NewRestaurantServiceClient(conn)

	ctx := context.TODO()

	param := &pb.Parameters{
		FromRestaurantID: int32(id),
		Statuses:         nil,
		ExcludeStatuses:  nil,
	}

	r, err := c.GetOrderList(ctx, param)
	if err != nil {
		return nil, fmt.Errorf("could not get order: %v", err)
	}
	return r, nil
}
