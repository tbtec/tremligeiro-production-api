package controller

import (
	"context"
	"log/slog"
	"time"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
	"github.com/tbtec/tremligeiro/internal/validator"
)

type ProducerProductionRestController struct {
	controller *ctl.ProducerProductionController
}

func NewProducerProductionRestController(container *container.Container) httpserver.IController {
	return &ProducerProductionRestController{
		controller: ctl.NewProducerProductionController(container),
	}
}

func (ctl *ProducerProductionRestController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {
	orderRequest := dto.Order{
		ID:        request.ParseParamString("id"),
		Status:    request.ParseParamString("status"),
		CreatedAt: time.Now(),
	}

	errBody := request.ParseBody(ctx, &orderRequest)
	if errBody != nil {
		return httpserver.HandleError(ctx, errBody)
	}

	err := validator.Validate(orderRequest)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	output, err := ctl.controller.Execute(ctx, orderRequest)
	if err != nil {
		slog.ErrorContext(ctx, "Error on checkout order", "error", err)
		return httpserver.HandleError(ctx, err)
	}

	return httpserver.Ok(output)
}
