package helpers

import (
	"encoding/json"
	"fmt"
)

func DeserializeJson[T any](body []byte) (*T, error) {
	var response T
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error deserializing json: %w", err)
	}
	return &response, nil
}
