package usecase

import "test-bpjs/src/api/v1/models"

type UscOrderInterfaces interface {
	CreateOrder(requestID int64, orders []*models.Order) error
	GetOrders() ([]models.Order, error)
	UpdateOrders(id string, requestID int64, orders []*models.Order) error
	DeleteOrder(id string) error
}
