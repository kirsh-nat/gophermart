package systemservices

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Good struct {
	Match      string `json:"match"`
	Reward     int    `json:"reward"`
	RewardType string `json:"reward_type"`
}

func CreateSystemGood(acrAddress string) error {
	url := fmt.Sprintf("http://%s/api/goods", acrAddress)
	goodReq := Good{
		Match:      "test",
		Reward:     10,
		RewardType: "%",
	}

	body, err := json.Marshal(goodReq)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("ошибка при чтении тела ответа: %w", err)
		}

		currError := string(bodyBytes)
		if currError == "mechanic already registered" {
			return nil
		}

		return err
	}

	return nil
}
