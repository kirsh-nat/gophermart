package systemservices

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func RegistrationSystemOrder(orderID, acrAddress string) error {
	url := fmt.Sprintf("http://%s/api/orders", acrAddress)
	good := GoodDesc{
		Description: "test",
		Price:       7299.8,
	}
	orderReq := OrderRequest{
		Order: orderID,
		Good:  []GoodDesc{good},
	}

	body, err := json.Marshal(orderReq)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return errors.New("invalid order")
	}

	return nil
}
