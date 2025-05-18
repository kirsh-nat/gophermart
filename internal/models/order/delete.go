package order

import (
	"context"
	"errors"
)

func (orderModel *OrderModel) Delete(ctx context.Context, model any) error {
	order, ok := (model).(*Order)
	if !ok {
		return errors.New("invalid model type")
	}

	_, err := orderModel.DB.ExecContext(ctx,
		"DELETE FROM orders WHERE id = ($1)",
		order.ID)

	if err != nil {
		return err
	}

	return nil
}
