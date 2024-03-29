package deliveryservice

import "github.com/nndergunov/deliveryApp/services/delivery/pkg/domain"

// DeliveryStorage is the interface for the delivery storage.
type DeliveryStorage interface {
	AssignOrder(order domain.AssignOrder) (*domain.AssignOrder, error)
	DeleteAssignedOrder(orderID int) error
}
