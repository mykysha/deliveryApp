package deliveryhandler

import (
	"github.com/nndergunov/deliveryApp/services/delivery/api/v1/rest/deliveryapi"
	"github.com/nndergunov/deliveryApp/services/delivery/pkg/domain"
)

func requestToOrder(req *deliveryapi.AssignOrderRequest) *domain.Order {
	return &domain.Order{FromUserID: req.FromUserID, FromRestaurantID: req.RestaurantID}
}
