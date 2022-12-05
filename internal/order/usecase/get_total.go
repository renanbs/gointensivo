package usecase

import (
	"github.com/renanbs/gointensivo/internal/order/entity"
	"github.com/renanbs/gointensivo/internal/order/infra/database"
)

type GetTotalOutputDTO struct {
	Total int
}

type GetTotalUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository database.OrderRepository) *GetTotalUseCase {
	return &GetTotalUseCase{
		OrderRepository: &orderRepository,
	}
}

func (c *GetTotalUseCase) Execute() (*GetTotalOutputDTO, error) {
	total, err := c.OrderRepository.GetTotal()
	if err != nil {
		return nil, err
	}

	return &GetTotalOutputDTO{
		Total: total,
	}, nil
}
