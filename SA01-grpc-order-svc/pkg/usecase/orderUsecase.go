package usecase

import (
	"context"
	"strings"

	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/domain"
	repository "github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/repository/interface"
	interfaces "github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/usecase/interface"
)

type orderUseCase struct {
	orderRepo repository.OrderRepository
}

// FetchOrder implements interfaces.OrderUseCase
func (o *orderUseCase) FetchOrder(ctx context.Context, userid int) (domain.RecOrder, error) {
	order, err := o.orderRepo.FetchOrder(ctx, userid)
	if err != nil {
		return domain.RecOrder{}, err
	}
	itemId := strings.Split(order.Item_id, ",")
	items := []domain.Item{}
	for _, itemid := range itemId {
		item, err := o.orderRepo.FetchItem(ctx, itemid)
		if err != nil {
			return domain.RecOrder{}, err
		}
		items = append(items, item)
	}
	recorder := domain.RecOrder{
		ID:           order.ID,
		Status:       order.Status,
		Item:         items,
		Total:        order.Total,
		CurrencyUnit: order.CurrencyUnit,
	}
	return recorder, err
}

// UpdateOrder implements interfaces.OrderUseCase
func (o *orderUseCase) UpdateOrder(ctx context.Context, orderid string, status string) (string, error) {
	id, err := o.orderRepo.UpdateOrder(ctx, orderid, status)
	return id, err
}

// CreateOrder implements interfaces.OrderUseCase
func (o *orderUseCase) CreateOrder(ctx context.Context, order domain.RecOrder) (string, error) {
	items := order.Item

	var it []string
	for _, item := range items {
		id, err := o.orderRepo.CreateItem(ctx, item)
		if err != nil {
			return "", err
		}
		it = append(it, id)
	}
	itemIDs := strings.Join(it, ",")
	od := domain.Order{
		ID:           order.ID,
		Status:       order.Status,
		Item_id:      itemIDs,
		Total:        order.Total,
		CurrencyUnit: order.CurrencyUnit,
	}
	id, err := o.orderRepo.CreateOrder(ctx, od)
	return id, err
}

func NewOrderUseCase(repo repository.OrderRepository) interfaces.OrderUseCase {
	return &orderUseCase{
		orderRepo: repo,
	}
}
