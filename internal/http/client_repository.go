package client

import (
	"fmt"
	"io"
	"net/http"
)

func (c *HttpClient) Get(url string) ([]byte, error) {
	url = Base_url + url

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error while getting the data: %w", err)
	}

	response, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while getting the data: %w", err)
	}
	defer response.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error while getting the data: %w", err)
	}
	if response.StatusCode > 299 {
		return nil, fmt.Errorf(
			"response failed with status code: %d and\nbody: %s",
			response.StatusCode,
			body,
		)
	}

	return body, nil
}
