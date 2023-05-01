// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interface/orderRepoInterface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	domain "github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockOrderRepository is a mock of OrderRepository interface.
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository.
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance.
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// CreateItem mocks base method.
func (m *MockOrderRepository) CreateItem(ctx context.Context, item domain.Item) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", ctx, item)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockOrderRepositoryMockRecorder) CreateItem(ctx, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockOrderRepository)(nil).CreateItem), ctx, item)
}

// CreateOrder mocks base method.
func (m *MockOrderRepository) CreateOrder(ctx context.Context, order domain.Order) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, order)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderRepositoryMockRecorder) CreateOrder(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderRepository)(nil).CreateOrder), ctx, order)
}

// FetchItem mocks base method.
func (m *MockOrderRepository) FetchItem(ctx context.Context, itemid string) (domain.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchItem", ctx, itemid)
	ret0, _ := ret[0].(domain.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchItem indicates an expected call of FetchItem.
func (mr *MockOrderRepositoryMockRecorder) FetchItem(ctx, itemid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchItem", reflect.TypeOf((*MockOrderRepository)(nil).FetchItem), ctx, itemid)
}

// FetchOrder mocks base method.
func (m *MockOrderRepository) FetchOrder(ctx context.Context, userid int, filter domain.Filter) ([]domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchOrder", ctx, userid, filter)
	ret0, _ := ret[0].([]domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchOrder indicates an expected call of FetchOrder.
func (mr *MockOrderRepositoryMockRecorder) FetchOrder(ctx, userid, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchOrder", reflect.TypeOf((*MockOrderRepository)(nil).FetchOrder), ctx, userid, filter)
}

// UpdateOrder mocks base method.
func (m *MockOrderRepository) UpdateOrder(ctx context.Context, orderid, status string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", ctx, orderid, status)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockOrderRepositoryMockRecorder) UpdateOrder(ctx, orderid, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockOrderRepository)(nil).UpdateOrder), ctx, orderid, status)
}
