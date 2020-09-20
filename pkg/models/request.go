package models

import (
	"net/http"
)

type Request struct {
	Payload   string
	Method    string
	Endpoint  string
	Transport *http.Transport
	Client    *http.Client
}
