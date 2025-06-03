package draftservices

import (
	"context"
	"database/sql"

	"github.com/kirsh-nat/gophermart.git/internal/models"
	orderservices "github.com/kirsh-nat/gophermart.git/internal/services/orderServices"
	userservices "github.com/kirsh-nat/gophermart.git/internal/services/userServices"
)

func NewUserDraft(DB *sql.DB, ctx context.Context, user *models.User, number string, payment float32) error {
	activeOrder, err := orderservices.GetByNumber(DB, ctx, number)
	if err != nil {
		return err
	}

	if activeOrder.UserID != user.ID {
		return NewUserNotAuthorizedError("Create Draft", err)
	}

	if activeOrder.Accural < payment && user.Balance < payment {
		return NewPaymentRequiredError("Create Draft", err)
	}

	_, err = CreateDraft(DB, ctx, number, user.ID, payment)
	if err != nil {
		return err
	}

	newBalance := user.Balance - payment
	newSpent := user.Spent + payment

	err = userservices.UpdateSpent(DB, ctx, user.ID, newBalance, newSpent)
	if err != nil {
		return err
	}

	return nil
}
