package gateway

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/event"
)

type OrderProducerGateway struct {
	producerService event.IProducerService
}

func NewOrderProducerGateway(producerService event.IProducerService) *OrderProducerGateway {
	return &OrderProducerGateway{
		producerService: producerService,
	}
}

func (gtw *OrderProducerGateway) PublishMessage(ctx context.Context, order dto.Order) error {

	err := gtw.producerService.PublishMessage(ctx, order)
	if err != nil {
		return err
	}

	return nil
}
