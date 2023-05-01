package repository

import (
	"context"
	"fmt"

	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/domain"
	interfaces "github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/repository/interface"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderDatabase struct {
	DB *gorm.DB
}

// FetchItem implements interfaces.OrderRepository
func (o *orderDatabase) FetchItem(ctx context.Context, itemid string) (domain.Item, error) {

	item := domain.Item{}
	err := o.DB.Where("id = ?", itemid).First(&item).Error

	return item, err
}

// FetchOrder implements interfaces.OrderRepository
func (o *orderDatabase) FetchOrder(ctx context.Context, userid int, filter domain.Filter) ([]domain.Order, error) {

	sql := "SELECT * FROM orders WHERE 1=1"
	if filter.Status != "" {
		sql += fmt.Sprintf(" AND status='%s'", filter.Status)
	}
	if filter.MinTotal != 0 {
		sql += fmt.Sprintf(" AND total>=%v", filter.MinTotal)
	}
	if filter.MaxTotal != 0 {
		sql += fmt.Sprintf(" AND total<=%v", filter.MaxTotal)
	}
	if filter.SortBy != "" {
		sql += fmt.Sprintf(" ORDER BY %s", filter.SortBy)
		if filter.SortOrder == "desc" {
			sql += " DESC"
		}
	}
	order := []domain.Order{}
	err := o.DB.Raw(sql).Scan(&order).Error
	return order, err
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
