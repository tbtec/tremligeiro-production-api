package usecase

import (
	"context"
	"log/slog"

	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type UscOrderProducer struct {
	orderProducerGateway *gateway.OrderProducerGateway
}

func NewUseCaseMessageProduction(orderProducerGateway *gateway.OrderProducerGateway,
) *UscOrderProducer {
	return &UscOrderProducer{
		orderProducerGateway: orderProducerGateway,
	}
}

func (usc *UscOrderProducer) Production(ctx context.Context, order dto.Order) (dto.Order, error) {
	slog.InfoContext(ctx, "Intiating message production", "orderId", order.ID)

	_ = usc.orderProducerGateway.PublishMessage(ctx, order)

	return order, nil
}
