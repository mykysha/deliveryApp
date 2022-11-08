// Package restaurantclient implements communication with restaurant service.
package restaurantclient

import (
	"context"
	"fmt"
	"log"

	"github.com/nndergunov/deliveryApp/services/order/pkg/domain"
	pb "github.com/nndergunov/deliveryApp/services/restaurant/api/v1/grpclogic/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RestaurantClient is responsible for communicating with restaurant service.
type RestaurantClient struct {
	restaurantURL string
}

// NewRestaurantClient returns new RestaurantClient instance.
func NewRestaurantClient(url string) *RestaurantClient {
	return &RestaurantClient{restaurantURL: url}
}

// CheckIfAvailable returns whether the restaurant can accept orders.
func (r RestaurantClient) CheckIfAvailable(restaurantID int) (bool, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(r.restaurantURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return false, fmt.Errorf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	c := pb.NewRestaurantServiceClient(conn)

	ctx := context.TODO()

	resp, err := c.GetRestaurant(ctx, &pb.Request{ID: int32(restaurantID)})
	if err != nil {
		return false, fmt.Errorf("getting restaurant: %w", err)
	}

	return resp.AcceptingOrders, nil
}

// CalculateOrderPrice returns the price of the order.
func (r RestaurantClient) CalculateOrderPrice(order domain.Order) (float64, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(r.restaurantURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 0, fmt.Errorf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	c := pb.NewRestaurantServiceClient(conn)

	ctx := context.TODO()

	menu, err := c.GetMenu(ctx, &pb.Request{ID: int32(order.RestaurantID)})
	if err != nil {
		return 0, fmt.Errorf("getting restaurant: %w", err)
	}

	var price float64

	for _, itemID := range order.OrderItems {
		for _, menuItem := range menu.MenuItems {
			if int(menuItem.ID) != itemID {
				continue
			}

			price += float64(menuItem.Price)
		}
	}

	return price, nil
}
