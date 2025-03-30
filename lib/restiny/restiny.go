package restiny

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	helper "github.com/pakornsrs/restiny/internal"
	"github.com/pakornsrs/restiny/models"
)

func GET(request models.RestinyRequest) (*http.Response, error) {
	req, err := http.NewRequest("GET", request.Endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(request.Headers) > 0 {
		req = helper.SetHTTPRequestHeader(req, request.Headers)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	return resp, nil
}

func POST[T any](request models.RestinyRequest, body T) (*http.Response, error) {

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
		req = helper.SetHTTPRequestHeader(req, request.Headers)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func PUT[T any](request models.RestinyRequest, body T) (*http.Response, error) {

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
		req = helper.SetHTTPRequestHeader(req, request.Headers)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func DELETE(request models.RestinyRequest) (*http.Response, error) {

	req, err := http.NewRequest("DELETE", request.Endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(request.Headers) > 0 {
		req = helper.SetHTTPRequestHeader(req, request.Headers)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	return resp, nil
}

func GETWithQueryParam(request models.RestinyRequest) (*http.Response, error) {

	path := fmt.Sprintf("%s?%s", request.Endpoint, request.Params.Encode())

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(request.Headers) > 0 {
		req = helper.SetHTTPRequestHeader(req, request.Headers)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	return resp, nil
}

func DELETEWithQueryParam(request models.RestinyRequest) (*http.Response, error) {
	path := fmt.Sprintf("%s?%s", request.Endpoint, request.Params.Encode())

	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(request.Headers) > 0 {
		req = helper.SetHTTPRequestHeader(req, request.Headers)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	return resp, nil
}
