package repository

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/infra/database"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
)

type IOrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
}

type OrderRepository struct {
	database database.RDBMS
}

func NewOrderRepository(database database.RDBMS) IOrderRepository {
	return &OrderRepository{
		database: database,
	}
}

func (repository *OrderRepository) Create(ctx context.Context, order *model.Order) error {

	result := repository.database.DB.WithContext(ctx).Create(&order)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
