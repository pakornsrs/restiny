package helper

import (
	"net/http"

	"github.com/pakornsrs/restiny/models"
)

func SetHTTPRequestHeader(req *http.Request, headers []models.Header) *http.Request {
	for _, header := range headers {
		if header.Key == "Content-Type" && header.Value == "application/json" {
			continue
		}

		req.Header.Set(header.Key, header.Value)
	}

	return req
}
