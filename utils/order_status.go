package utils

import "fmt"

type OrderStatus string

const (
	OrderStatusNew        OrderStatus = "new"
	OrderStatusInProgress OrderStatus = "progress"
	OrderStatusDone       OrderStatus = "done"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

func (s OrderStatus) IsValid() bool {
	switch s {
	case OrderStatusNew, OrderStatusInProgress, OrderStatusDone, OrderStatusCancelled:
		return true
	}
	return false
}

func ToOrderStatus(status string) (OrderStatus, error) {
	s := OrderStatus(status)
	if s.IsValid() {
		return s, nil
	}
	return "", fmt.Errorf("invalid order status: %s", status)
}
