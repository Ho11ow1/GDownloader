package Utils

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct{}

func (this HTTPClient) Get(url string, timeout time.Duration) (res []byte, err error){
	//
	fmt.Print("")

	client := &http.Client {
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed reading body: %w", err)
	}

	return body, nil
}
