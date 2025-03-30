package models

import "net/url"

type RestinyRequest struct {
	Endpoint string
	Params   *url.Values
	Headers  []Header
}

type Header struct {
	Key   string
	Value string
}
