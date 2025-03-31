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

func GET(request RestinyRequest) (*http.Response, error) {
	req, err := http.NewRequest("GET", request.Endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(request.Headers) > 0 {
		req = setHTTPRequestHeader(req, request.Headers)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	return resp, nil
}

func POST[T any](request RestinyRequest, body T) (*http.Response, error) {

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", request.Endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(request.Headers) > 0 {
		req = setHTTPRequestHeader(req, request.Headers)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func PUT[T any](request RestinyRequest, body T) (*http.Response, error) {

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", request.Endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(request.Headers) > 0 {
		req = setHTTPRequestHeader(req, request.Headers)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func DELETE(request RestinyRequest) (*http.Response, error) {

	req, err := http.NewRequest("DELETE", request.Endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(request.Headers) > 0 {
		req = setHTTPRequestHeader(req, request.Headers)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	return resp, nil
}

func GETWithQueryParam(request RestinyRequest) (*http.Response, error) {

	path := fmt.Sprintf("%s?%s", request.Endpoint, request.Params.Encode())

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(request.Headers) > 0 {
		req = setHTTPRequestHeader(req, request.Headers)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	return resp, nil
}

func DELETEWithQueryParam(request RestinyRequest) (*http.Response, error) {
	path := fmt.Sprintf("%s?%s", request.Endpoint, request.Params.Encode())

	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(request.Headers) > 0 {
		req = setHTTPRequestHeader(req, request.Headers)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	return resp, nil
}

func setHTTPRequestHeader(req *http.Request, headers []Header) *http.Request {
	for _, header := range headers {
		if header.Key == "Content-Type" && header.Value == "application/json" {
			continue
		}

		req.Header.Set(header.Key, header.Value)
	}

	return req
}
