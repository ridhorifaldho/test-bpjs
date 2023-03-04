package repository

import (
	"gorm.io/gorm"
	"test-bpjs/src/api/v1/models"
)

type RepoOrderStruct interface {
	CreateOrder(order *models.Order) error
	FindAllOrders() ([]models.Order, error)
	FindOrderByID(order *models.Order, id string) error
	UpdateOrder(id string, order *models.Order) error
	DeleteOrder(id string) error
}

type TransactionalDBRepoInterface interface {
	BeginTrans() *gorm.DB
}
