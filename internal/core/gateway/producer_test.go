package gateway

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tbtec/tremligeiro/internal/dto"
)

// Mock para event.IProducerService
type MockProducerService struct {
	mock.Mock
}

func (m *MockProducerService) PublishMessage(ctx context.Context, message interface{}) error {
	args := m.Called(ctx, message)
	return args.Error(0)
}

func TestOrderProducerGatewayPublishMessageSuccess(t *testing.T) {
	mockProducer := new(MockProducerService)
	gateway := NewOrderProducerGateway(mockProducer)

	order := dto.Order{
		ID:        "order-1",
		Status:    "PENDING",
		CreatedAt: time.Now(),
	}

	mockProducer.On("PublishMessage", mock.Anything, mock.Anything).Return(nil)

	err := gateway.PublishMessage(context.Background(), order)
	assert.NoError(t, err)
	mockProducer.AssertExpectations(t)
}

func TestOrderProducerGatewayPublishMessageError(t *testing.T) {
	mockProducer := new(MockProducerService)
	gateway := NewOrderProducerGateway(mockProducer)

	order := dto.Order{
		ID:        "order-2",
		Status:    "FAILED",
		CreatedAt: time.Now(),
	}

	mockProducer.On("PublishMessage", mock.Anything, mock.Anything).Return(errors.New("publish error"))

	err := gateway.PublishMessage(context.Background(), order)
	assert.Error(t, err)
	assert.EqualError(t, err, "publish error")
	mockProducer.AssertExpectations(t)
}
