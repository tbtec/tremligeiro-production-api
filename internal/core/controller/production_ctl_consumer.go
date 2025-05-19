package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type ConsumerProductionController struct {
	usc *usecase.UscCreateOrder
}

func NewConsumerProductionController(container *container.Container) *ConsumerProductionController {
	return &ConsumerProductionController{
		usc: usecase.NewUscCreateOrder(
			gateway.NewOrderGateway(container.OrderRepository),
		),
	}
}

func (ctl *ConsumerProductionController) Execute(ctx context.Context, order dto.Order) error {
	return ctl.usc.Create(ctx, order)
}
