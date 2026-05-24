package Utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

type HTTPClient struct {
    Client *http.Client
}

func (this HTTPClient) Get(pageURL string) (res []byte, err error) {
	//
	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Creating HTTP request failed: %w", err)
	}

	this.SetDefaultHeaders(req, pageURL)

	resp, err := this.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Sending HTTP GET failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) 
	if err != nil {
		return nil, fmt.Errorf("Failed reading body: %w", err)
	}

	return body, nil
}

func (this HTTPClient) Post(pageURL string, postURL string, body []byte) (res []byte, err error) {
	//
    req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(body))
    if err != nil {
        return nil, fmt.Errorf("Creating HTTP request failed: %w", err)
    }

    this.SetDefaultHeaders(req, pageURL)
    req.Header.Set("Content-Type", "application/json")

    resp, err := this.Client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("Sending HTTP POST failed: %w", err)
    }
    defer resp.Body.Close()

    resBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("Failed reading body: %w", err)
    }
	
    return resBody, nil
}

func (this HTTPClient) Download(pageURL string, downloadURL string, destPath string) error {
	//
    req, err := http.NewRequest("GET", downloadURL, nil)
    if err != nil {
        return fmt.Errorf("Creating request failed: %w", err)
    }

    this.SetDefaultHeaders(req, pageURL)

    resp, err := this.Client.Do(req)
    if err != nil {
        return fmt.Errorf("Request failed: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("Unexpected status: %d", resp.StatusCode)
    }

    this.WriteToDisc(destPath, resp)

    return nil
}

func (this HTTPClient) WriteToDisc(destPath string, resp *http.Response) error {
    //
    file, err := os.Create(destPath)
    if err != nil {
        return fmt.Errorf("Failed to create file: %w", err)
    }
    defer file.Close()

    _, err = io.Copy(file, resp.Body)
    if err != nil {
        return fmt.Errorf("Failed writing to disk: %w", err)
    }

    return nil
}

func (this HTTPClient) SetDefaultHeaders(req *http.Request, pageURL string) {
	//
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:126.0) Gecko/20100101 Firefox/126.0",
	)

	req.Header.Set("Referer", pageURL)
}
