package repository

import (
	"context"

	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/domain"
	interfaces "github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/repository/interface"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderDatabase struct {
	DB *gorm.DB
}

// UpdateOrder implements interfaces.OrderRepository
func (o *orderDatabase) UpdateOrder(ctx context.Context, orderid string, status string) (string, error) {
	order := domain.Order{}
	err := o.DB.Model(&order).Where("id = ?", orderid).Update("status", status).Error
	return orderid, err
}

// CreateItem implements interfaces.OrderRepository
func (o *orderDatabase) CreateItem(ctx context.Context, item domain.Item) (string, error) {
	err := o.DB.Create(&item).Error
	return item.ID, err
}

// CreateOrder implements interfaces.OrderRepository
func (o *orderDatabase) CreateOrder(ctx context.Context, order domain.Order) (string, error) {
	order.ID = uuid.New().String()
	err := o.DB.Create(&order).Error
	return order.ID, err
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &orderDatabase{DB}
}
