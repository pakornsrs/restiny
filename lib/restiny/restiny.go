package restiny

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type RestinyRequest struct {
	Endpoint string
	Params   *url.Values
	Headers  []Header
}

type Header struct {
	Key   string
	Value string
}

func GET(endpoint, apiKey string) (*http.Response, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(apiKey) > 0 {
		req.Header.Set("Authorization", apiKey)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	return resp, nil
}

func POST[T any](endpoint string, body T, apiKey string) (*http.Response, error) {

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(apiKey) > 0 {
		req.Header.Set("Authorization", apiKey)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func PUT[T any](endpoint string, body T, apiKey string) (*http.Response, error) {

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(apiKey) > 0 {
		req.Header.Set("Authorization", apiKey)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func DELETE(endpoint string, apiKey string) (*http.Response, error) {

	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(apiKey) > 0 {
		req.Header.Set("Authorization", apiKey)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	return resp, nil
}

func GETWithQueryParam(endpoint string, params *url.Values, apiKey string) (*http.Response, error) {

	path := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(apiKey) > 0 {
		req.Header.Set("Authorization", apiKey)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	return resp, nil
}

func DELETEWithQueryParam(endpoint string, params *url.Values, apiKey string) (*http.Response, error) {
	path := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(apiKey) > 0 {
		req.Header.Set("Authorization", apiKey)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	return resp, nil
}
