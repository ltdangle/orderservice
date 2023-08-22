package helpers

import (
	"io"
	"net/http"
	"time"
)

func MakeHttpRequest(url string, method string) (string, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}

	client := http.Client{
		Timeout: 20 * time.Second,
	}

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
