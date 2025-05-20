package usecase

import (
	"context"
	"log/slog"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type UscCreateOrder struct {
	orderGateway *gateway.OrderGateway
}

func NewUscCreateOrder(orderGateway *gateway.OrderGateway) *UscCreateOrder {
	return &UscCreateOrder{
		orderGateway: orderGateway,
	}
}

func (usc *UscCreateOrder) Create(ctx context.Context, order dto.Order) error {
	slog.InfoContext(ctx, "Create order menssage", "orderId", order.ID, "orderStatus", order.Status)

	orderEntity := &entity.Order{
		ID:        order.ID,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
	}

	err := usc.orderGateway.Create(ctx, orderEntity)
	if err != nil {
		return err
	}

	return nil
}
