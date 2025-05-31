package systemservices

import (
	"context"
	"database/sql"
	"log"

	"github.com/kirsh-nat/gophermart.git/internal/app"
	orderservices "github.com/kirsh-nat/gophermart.git/internal/services/orderServices"
	userservices "github.com/kirsh-nat/gophermart.git/internal/services/userServices"
)

func Worker(sqlDB *sql.DB, ctx context.Context, acrAddress string) {
	err := CreateSystemGood(app.AccrualAddress)
	if err != nil {
		log.Printf("Error creating system good: %s", err)
		return
	}
	orderNumbers, err := orderservices.GetActiveOrders(sqlDB, ctx)
	if err != nil {
		return
	}

	for _, orderNum := range orderNumbers {
		result, err := NotifyAccrualSystem(orderNum, acrAddress)

		if err != nil {
			log.Printf("Error notifying accrual system for order %s: %v", orderNum, err)
			continue
		}

		if result.Order == "" {
			continue
		}
		userID, updateErr := UpdateStatus(sqlDB, ctx, result)
		if updateErr != nil {
			log.Printf("Error updating status for order %s: %v", orderNum, updateErr)
		}

		err = userservices.UpdateBalance(sqlDB, ctx, userID, result.Accrual)
		if err != nil {
			log.Printf("Error update balance for user: %s \n", err)
			continue
		}

	}
}
