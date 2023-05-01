package repository

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/domain"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// func TestCreateOrder(t *testing.T) {
// 	mockDB, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("error creating sqlmock: %s", err)
// 	}
// 	defer mockDB.Close()

// 	// Create a new gorm.DB instance using the mock database connection
// 	gormDB, err := gorm.Open(mysql.New(mysql.Config{
// 		Conn:                      mockDB,
// 		SkipInitializeWithVersion: true,
// 	}), &gorm.Config{
// 		SkipDefaultTransaction: true,
// 	})
// 	fmt.Println(err)
// 	if err != nil {
// 		t.Fatalf("failed to open GORM connection: %v", err)
// 	}

// 	// Create a new order repository using the mock gorm.DB instance
// 	orderRepo := NewOrderRepository(gormDB)

// 	// Set up expectations for the SQL query
// 	// expectedID := uuid.New().String()

// 	// Call the CreateOrder method
// 	mockOrder := domain.Order{
// 		ID:           "",
// 		Status:       "",
// 		Item_id:      "",
// 		Total:        0,
// 		CurrencyUnit: "",
// 	}
// 	mock.ExpectExec("INSERT INTO `orders`").WithArgs( mockOrder.Status, mockOrder.Item_id, mockOrder.Total, mockOrder.CurrencyUnit,mockOrder.ID)
// 	_, actualErr := orderRepo.CreateOrder(context.Background(), mockOrder)

// 	// Assert that the expected and actual results are equal
// 	assert.NoError(t, actualErr)
// 	// err = mock.ExpectationsWereMet()
// 	// assert.Equal(t, expectedID, actualID)
// }
