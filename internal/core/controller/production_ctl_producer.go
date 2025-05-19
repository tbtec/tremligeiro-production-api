package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type ProducerProductionController struct {
	usc *usecase.UscOrderCheckOut
}

func NewProducerProductionController(container *container.Container) *ProducerProductionController {
	return &ProducerProductionController{
		usc: usecase.NewUseCaseMessageProduction(
			gateway.NewOrderProducerGateway(container.ProducerService),
		),
	}
}

func (ctl *ProducerProductionController) Execute(ctx context.Context, input dto.Order) (dto.Order, error) {
	return ctl.usc.Production(ctx, input)
}
