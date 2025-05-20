package order

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Result struct {
	Order   string `json:"order"`
	Status  string `json:"status"`
	Accrual int    `json:"accrual"`
}

type OrderRequest struct {
	Order string `json:"order"`
	Good  []GoodDesc
}

type GoodDesc struct {
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func NotifyAccrualSystem(orderID, acrAddress string) (Result, error) {
	url := fmt.Sprintf("http://%s/api/orders/%s", acrAddress, orderID)
	resp, err := http.Get(url)
	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()
	errEmpty := errors.New("empty result")
	switch resp.StatusCode {
	case http.StatusOK:
		var result Result
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return Result{}, err
		}
		return result, nil
	case http.StatusTooManyRequests:
		return Result{}, errEmpty
	case http.StatusInternalServerError:
		return Result{}, errEmpty
	case http.StatusNoContent:
		err := RegistrationSystemOrder(orderID, acrAddress)
		if err != nil {
			return Result{Order: orderID, Status: INVALID}, nil
		}
		return Result{}, nil
	}

	return Result{}, nil
}
