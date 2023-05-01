package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/domain"
	mock "github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/mock/repoMock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"gopkg.in/go-playground/assert.v1"
)

func TestUpdateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockrepo := mock.NewMockOrderRepository(ctrl)

	usecase := NewOrderUseCase(mockrepo)
	ctx := context.Background()
	stringid := uuid.New().String()
	testData := []struct {
		name       string
		orderid    string
		status     string
		beforeTest func(userRepo *mock.MockOrderRepository)
		expectErr  error
	}{
		{
			name:    "Test success response",
			orderid: stringid,
			status:  "pending",
			beforeTest: func(userRepo *mock.MockOrderRepository) {
				userRepo.EXPECT().UpdateOrder(ctx, stringid, "pending").Return(stringid, nil)
			},
			expectErr: nil,
		},
		{
			name:    "Test when user alredy exist response",
			orderid: stringid,
			status:  "pending",
			beforeTest: func(userRepo *mock.MockOrderRepository) {
				userRepo.EXPECT().UpdateOrder(ctx, stringid, "pending").Return(stringid, errors.New("repo error"))
			},
			expectErr: errors.New("repo error"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(mockrepo)
			actualId, err := usecase.UpdateOrder(ctx, tt.orderid, tt.status)
			assert.Equal(t, tt.expectErr, err)
			if err == nil {
				assert.Equal(t, tt.orderid, actualId)
			}
		})
	}
}

func TestCreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockrepo := mock.NewMockOrderRepository(ctrl)

	usecase := NewOrderUseCase(mockrepo)
	ctx := context.Background()
	stringid := uuid.New().String()
	testData := []struct {
		name       string
		order      domain.RecOrder
		beforeTest func(userRepo *mock.MockOrderRepository)
		expectErr  error
	}{
		{
			name: "Test order with no item",
			order: domain.RecOrder{
				ID:           stringid,
				Status:       "PENDING_INVOICE",
				Item:         []domain.Item{},
				Total:        0,
				CurrencyUnit: "USD",
			},
			beforeTest: func(userRepo *mock.MockOrderRepository) {},
			expectErr:  errors.New("there no item in order"),
		},
		{
			name: "Test with create item repo error",
			order: domain.RecOrder{
				ID:     stringid,
				Status: "PENDING_INVOICE",
				Item: []domain.Item{
					{
						ID:          "12345",
						Description: "Sample product description",
						Price:       12.40,
						Quantity:    2,
					},
				},
				Total:        24.80,
				CurrencyUnit: "USD",
			},
			beforeTest: func(userRepo *mock.MockOrderRepository) {

				userRepo.EXPECT().CreateItem(ctx, domain.Item{
					ID:          "12345",
					Description: "Sample product description",
					Price:       12.40,
					Quantity:    2,
				}).Return("12345", errors.New("item repo error"))
			},
			expectErr: errors.New("item repo error"),
		},
		{
			name: "Test with order repo error",
			order: domain.RecOrder{
				ID:     stringid,
				Status: "PENDING_INVOICE",
				Item: []domain.Item{
					{
						ID:          "12345",
						Description: "Sample product description",
						Price:       12.40,
						Quantity:    2,
					},
				},
				Total:        24.80,
				CurrencyUnit: "USD",
			},
			beforeTest: func(userRepo *mock.MockOrderRepository) {

				userRepo.EXPECT().CreateItem(ctx, domain.Item{
					ID:          "12345",
					Description: "Sample product description",
					Price:       12.40,
					Quantity:    2,
				}).Return("12345", nil)

				userRepo.EXPECT().CreateOrder(ctx, domain.Order{
					ID:           stringid,
					Status:       "PENDING_INVOICE",
					Item_id:      "12345",
					Total:        24.80,
					CurrencyUnit: "USD",
				}).Return(stringid, errors.New("create rder repo error"))
			},
			expectErr: errors.New("create rder repo error"),
		},
		{
			name: "Test success response",
			order: domain.RecOrder{
				ID:     stringid,
				Status: "PENDING_INVOICE",
				Item: []domain.Item{
					{
						ID:          "12345",
						Description: "Sample product description",
						Price:       12.40,
						Quantity:    2,
					},
				},
				Total:        24.80,
				CurrencyUnit: "USD",
			},
			beforeTest: func(userRepo *mock.MockOrderRepository) {

				userRepo.EXPECT().CreateItem(ctx, domain.Item{
					ID:          "12345",
					Description: "Sample product description",
					Price:       12.40,
					Quantity:    2,
				}).Return("12345", nil)

				userRepo.EXPECT().CreateOrder(ctx, domain.Order{
					ID:           stringid,
					Status:       "PENDING_INVOICE",
					Item_id:      "12345",
					Total:        24.80,
					CurrencyUnit: "USD",
				}).Return(stringid, nil)
			},
			expectErr: nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(mockrepo)
			actualId, err := usecase.CreateOrder(ctx, tt.order)
			assert.Equal(t, tt.expectErr, err)
			if err == nil {
				assert.Equal(t, tt.order.ID, actualId)
			}
		})
	}
}
