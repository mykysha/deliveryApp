package handler

import (
	"context"
	"strconv"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/domain"

	pb "github.com/nndergunov/deliveryApp/app/services/courier/api/v1/grpc/proto"

	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/service/courierservice"
)

type Params struct {
	Logger         *logger.Logger
	CourierService courierservice.CourierService
}

// handler is the entrypoint into our application
type handler struct {
	pb.UnsafeCourierServer
	service courierservice.CourierService
	log     *logger.Logger
}

// NewHandler returns new http multiplexer with configured endpoints.
func NewHandler(p Params) *handler {
	return &handler{
		log:     p.Logger,
		service: p.CourierService,
	}
}

func (h *handler) InsertNewCourier(ctx context.Context, in *pb.NewCourierRequest) (*pb.CourierResponse, error) {
	courier := domain.Courier{
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Email:     in.Email,
		Phone:     in.Phone,
	}

	resp, err := h.service.InsertCourier(courier)
	if err != nil {
		return nil, err
	}

	return &pb.CourierResponse{
		ID:        int64(resp.ID),
		Firstname: resp.Firstname,
		Lastname:  resp.Lastname,
		Email:     resp.Email,
		Phone:     resp.Phone,
	}, nil
}

func (h *handler) GetAllCourier(ctx context.Context, in *pb.SearchParam) (*pb.CourierListResponse, error) {
	param := domain.SearchParam{}
	param["available"] = in.GetAvailable()
	respList, err := h.service.GetCourierList(param)
	if err != nil {
		return nil, err
	}

	if respList == nil {
		return nil, nil
	}

	var outList []*pb.CourierResponse

	for _, resp := range respList {
		out := &pb.CourierResponse{
			ID:        int64(resp.ID),
			Firstname: resp.Firstname,
			Lastname:  resp.Lastname,
			Email:     resp.Email,
			Phone:     resp.Phone,
		}
		outList = append(outList, out)
	}
	return &pb.CourierListResponse{CourierListResponse: outList}, nil
}

func (h *handler) DeleteCourier(ctx context.Context, in *pb.CourierID) (*pb.CourierDeleteResponse, error) {
	out, err := h.service.DeleteCourier(strconv.Itoa(int(in.CourierID)))
	if err != nil {
		return nil, err
	}

	return &pb.CourierDeleteResponse{CourierDeleteResponse: out}, nil
}

func (h *handler) UpdateCourier(ctx context.Context, in *pb.UpdateCourierRequest) (*pb.CourierResponse, error) {
	courier := domain.Courier{
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Email:     in.Email,
		Phone:     in.Phone,
	}

	resp, err := h.service.UpdateCourier(courier, strconv.FormatInt(in.ID, 10))
	if err != nil {
		return nil, err
	}

	return &pb.CourierResponse{
		ID:        int64(resp.ID),
		Firstname: resp.Firstname,
		Lastname:  resp.Lastname,
		Email:     resp.Email,
		Phone:     resp.Phone,
	}, nil
}

func (h *handler) GetCourier(ctx context.Context, in *pb.CourierID) (*pb.CourierResponse, error) {
	resp, err := h.service.GetCourier(strconv.FormatInt(in.CourierID, 10))
	if err != nil {
		return nil, err
	}

	return &pb.CourierResponse{
		ID:        int64(resp.ID),
		Firstname: resp.Firstname,
		Lastname:  resp.Lastname,
		Email:     resp.Email,
		Phone:     resp.Phone,
	}, nil
}

func (h *handler) InsertNewCourierLocation(ctx context.Context, in *pb.Location) (*pb.Location, error) {
	location := domain.Location{
		UserID:     int(in.UserID),
		Latitude:   *in.Latitude,
		Longitude:  *in.Longitude,
		Country:    *in.Country,
		City:       *in.City,
		Region:     *in.Region,
		Street:     *in.Street,
		HomeNumber: *in.HomeNumber,
		Floor:      *in.Floor,
		Door:       *in.Door,
	}

	resp, err := h.service.InsertLocation(location, strconv.Itoa(location.UserID))
	if err != nil {
		return nil, err
	}

	return &pb.Location{
		UserID:     int64(resp.UserID),
		Latitude:   &resp.Latitude,
		Longitude:  &resp.Longitude,
		Country:    &resp.Country,
		City:       &resp.City,
		Region:     &resp.Region,
		Street:     &resp.Street,
		HomeNumber: &resp.HomeNumber,
		Floor:      &resp.Floor,
		Door:       &resp.Door,
	}, nil
}

func (h *handler) UpdateCourierLocation(ctx context.Context, in *pb.Location) (*pb.Location, error) {
	location := domain.Location{
		UserID:     int(in.UserID),
		Latitude:   *in.Latitude,
		Longitude:  *in.Longitude,
		Country:    *in.Country,
		City:       *in.City,
		Region:     *in.Region,
		Street:     *in.Street,
		HomeNumber: *in.HomeNumber,
		Floor:      *in.Floor,
		Door:       *in.Door,
	}

	resp, err := h.service.UpdateLocation(location, strconv.Itoa(location.UserID))
	if err != nil {
		return nil, err
	}

	return &pb.Location{
		UserID:     int64(resp.UserID),
		Latitude:   &resp.Latitude,
		Longitude:  &resp.Longitude,
		Country:    &resp.Country,
		City:       &resp.City,
		Region:     &resp.Region,
		Street:     &resp.Street,
		HomeNumber: &resp.HomeNumber,
		Floor:      &resp.Floor,
		Door:       &resp.Door,
	}, nil
}

func (h *handler) GetCourierLocation(ctx context.Context, in *pb.UserID) (*pb.Location, error) {
	resp, err := h.service.GetLocation(strconv.FormatInt(in.UserID, 10))
	if err != nil {
		return nil, err
	}

	return &pb.Location{
		UserID:     int64(resp.UserID),
		Latitude:   &resp.Latitude,
		Longitude:  &resp.Longitude,
		Country:    &resp.Country,
		City:       &resp.City,
		Region:     &resp.Region,
		Street:     &resp.Street,
		HomeNumber: &resp.HomeNumber,
		Floor:      &resp.Floor,
		Door:       &resp.Door,
	}, nil
}