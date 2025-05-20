package gateway

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/internal/infra/database/repository"
)

type OrderGateway struct {
	orderRepository repository.IOrderRepository
}

func NewOrderGateway(orderRepository repository.IOrderRepository) *OrderGateway {
	return &OrderGateway{
		orderRepository: orderRepository,
	}
}

func (gtw *OrderGateway) Create(ctx context.Context, order *entity.Order) error {

	orderModel := model.Order{
		ID:        order.ID,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
	}

	err := gtw.orderRepository.Create(ctx, &orderModel)

	if err != nil {
		return err
	}

	return nil
}
