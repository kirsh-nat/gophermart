package order

import (
	"context"
	"database/sql"
	"log"
)

func Worker(sqlDb *sql.DB, ctx context.Context, acrAddress string) {
	orderNumbers, err := GetActiveOrders(sqlDb, ctx)
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

		if updateErr := UpdateStatus(sqlDb, ctx, result); updateErr != nil {
			log.Printf("Error updating status for order %s: %v", orderNum, updateErr)
		}
	}
}
