package courierclient

import (
	"context"
	"fmt"
	pb "github.com/nndergunov/deliveryApp/app/services/courier/api/v1/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CourierClient struct {
	courierURL string
}

func NewCourierClient(url string) *CourierClient {
	return &CourierClient{courierURL: url}
}

func (a CourierClient) GetLocation(city string) (*pb.LocationList, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(a.courierURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCourierClient(conn)

	// Contact the server and print out its response.
	ctx := context.TODO()

	r, err := c.GetLocationList(ctx, &pb.SearchParamLocation{City: &city})
	if err != nil {
		return nil, fmt.Errorf("could not get locations: %v", err)
	}
	return r, nil
}

func (a CourierClient) UpdateCourierAvailable(courierID int, available string) (*pb.CourierResponse, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(a.courierURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCourierClient(conn)

	// Contact the server and print out its response.
	ctx := context.TODO()

	r, err := c.UpdateCourier(ctx, &pb.SearchParamLocation{City: &city})
	if err != nil {
		return nil, fmt.Errorf("could not get locations: %v", err)
	}
	return r, nil
}
