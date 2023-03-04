package order

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"test-bpjs/src/api/v1/models"
	"test-bpjs/src/api/v1/repository"
	"test-bpjs/src/api/v1/usecase"
)

type UscOrderStruct struct {
	OrderRepo         repository.RepoOrderStruct
	TransactionalRepo repository.TransactionalDBRepoInterface
}

func OrderUsecaseImpl(OrderRepo repository.RepoOrderStruct, TransactionalRepo repository.TransactionalDBRepoInterface) usecase.UscOrderInterfaces {
	return &UscOrderStruct{OrderRepo, TransactionalRepo}
}

func (Or *UscOrderStruct) CreateOrder(requestID int64, orders []*models.Order) error {

	tx := Or.TransactionalRepo.BeginTrans()
	if tx.Error != nil {
		return tx.Error
	}

	var wg sync.WaitGroup
	wg.Add(len(orders))

	errs := make(chan error, len(orders))

	for _, order := range orders {
		go func(order *models.Order) {
			defer wg.Done()

			//order.Timestamp = time.Now()
			order.RequestID = requestID

			err := Or.OrderRepo.CreateOrder(order)
			if err != nil {
				tx.Rollback()
				errs <- fmt.Errorf("failed to save order: %w", err)
			}
		}(order)
	}

	wg.Wait()

	close(errs)

	for err := range errs {
		if err != nil {
			tx.Rollback()
			return errors.New("one or more orders failed to save")
		}
	}

	tx.Commit()
	return nil
}

func (Or *UscOrderStruct) GetOrders() ([]models.Order, error) {
	return Or.OrderRepo.FindAllOrders()
}

func (Or *UscOrderStruct) UpdateOrders(id string, requestID int64, orders []*models.Order) error {
	tx := Or.TransactionalRepo.BeginTrans()
	if tx.Error != nil {
		return tx.Error
	}

	var wg sync.WaitGroup
	wg.Add(len(orders))

	errs := make(chan error, len(orders))

	for _, order := range orders {
		go func(order *models.Order) {
			defer wg.Done()

			var existingOrder models.Order
			err := Or.OrderRepo.FindOrderByID(&existingOrder, strconv.Itoa(int(order.ID)))
			if err != nil {
				tx.Rollback()
				errs <- fmt.Errorf("failed to find order: %w", err)
				return
			}

			existingOrder.RequestID = requestID
			existingOrder.Customer = order.Customer
			existingOrder.Quantity = order.Quantity
			existingOrder.Price = order.Price
			//existingOrder.Timestamp = time.Now()

			err = Or.OrderRepo.UpdateOrder(id, &existingOrder)
			if err != nil {
				tx.Rollback()
				errs <- fmt.Errorf("failed to update order: %w", err)
			}
		}(order)
	}

	wg.Wait()

	close(errs)

	for err := range errs {
		if err != nil {
			tx.Rollback()
			return errors.New("one or more orders failed to update")
		}
	}

	tx.Commit()
	return nil

}

func (Or *UscOrderStruct) DeleteOrder(id string) error {
	return Or.OrderRepo.DeleteOrder(id)
}
