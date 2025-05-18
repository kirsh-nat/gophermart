package order

import (
	"context"
	"errors"
)

func (orderModel *OrderModel) Update(ctx context.Context, model any) (any, error) {
	order, ok := (model).(*Order)
	if !ok {
		return &Order{}, errors.New("invalid model type")
	}
	//TODO: add updated at time
	_, err := orderModel.DB.ExecContext(ctx,
		"UPDATE orders (status, accural) VALUES ($1, $2) WHERE id = ($3)",
		order.Status, order.Accural, order.ID)

	if err != nil {
		return order, err
	}

	return order, nil
}
