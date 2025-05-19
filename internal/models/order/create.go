package order

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func (OrderModel *OrderModel) Create(ctx context.Context, anyModel any) (any, error) {
	model, ok := anyModel.(*Order)
	if !ok {
		return &Order{}, errors.New("invalid model type")
	}

	checkOrder, err := OrderModel.GetByNumber(ctx, model.Number)
	if err != nil {
		//	var notFoundErr OrderNotFoundError
		if errors.As(err, &OrderNotFoundError{}) {
		} else {
			return &Order{}, err
		}
	} else {
		if checkOrder.UserID == model.UserID {
			return &Order{}, NewUserNumberExistsError("Create Order", err)
		}
		return &Order{}, NewNumberExists("Create Order", err)
	}

	result, err := OrderModel.DB.ExecContext(ctx,
		"INSERT INTO Orders (number, user_id) VALUES ($1, $2)", model.Number, model.UserID)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return model, NewNumberExists("Create Order", err)
		}
		return model, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model, err
	}
	if rowsAffected == 0 {
		return &Order{}, errors.New("failed to create order")
	}

	return model, nil
}

// func (OrderModel *OrderModel) Create(ctx context.Context, anyModel any) (any, error) {
// 	model, ok := (anyModel).(*Order)
// 	if !ok {
// 		return &Order{}, errors.New("invalid model type")
// 	}

// 	checkOrder, err := OrderModel.GetByNumber(ctx, model.Number)
// 	if err != nil {
// 		var notFoundErr OrderNotFoundError
// 		if errors.As(err, &notFoundErr) == false {
// 			return &Order{}, err
// 		}
// 	}

// 	if checkOrder.ID != 0 {
// 		if checkOrder.UserID == model.UserID {
// 			return &Order{}, NewUserNumberExistsError("Create Order", err)
// 		}
// 		return &Order{}, NewNumberExists("Create Order", err)
// 	}

// 	_, err = OrderModel.DB.ExecContext(ctx,
// 		"INSERT INTO Orders (number, user_id) VALUES ($1, $2)", model.Number, model.UserID)

// 	if err != nil {
// 		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
// 			return model, NewNumberExists("Create Order", err)
// 		}
// 		return model, err
// 	}

// 	return model, nil
// }
