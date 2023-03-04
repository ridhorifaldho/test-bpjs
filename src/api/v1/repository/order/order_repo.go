package order

import (
	"gorm.io/gorm"
	"test-bpjs/src/api/v1/models"
	"test-bpjs/src/api/v1/repository"
)

type RepoOrderStruct struct {
	db *gorm.DB
}

func NewRepoOrderImpl(db *gorm.DB) repository.RepoOrderStruct {
	return &RepoOrderStruct{db}
}

func (Or *RepoOrderStruct) CreateOrder(order *models.Order) error {
	return Or.db.Debug().Create(order).Error
}

func (Or *RepoOrderStruct) FindAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := Or.db.Find(&orders).Error
	return orders, err
}

func (Or *RepoOrderStruct) FindOrderByID(order *models.Order, id string) error {
	return Or.db.Where("id = ?", id).First(order).Error
}

func (Or *RepoOrderStruct) UpdateOrder(id string, order *models.Order) error {
	return Or.db.Where("id = ?", id).Updates(&order).Error
}

func (Or *RepoOrderStruct) DeleteOrder(id string) error {
	return Or.db.Where("id = ?", id).Delete(&models.Order{}).Error
}
