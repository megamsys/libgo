package geard

import (
	"net/http"
)

const (
	rawResponse = iota
	normalResponse
)

type responseType int

type RawResponse struct {
	StatusCode int
	Body       []byte
	Header     http.Header
}

var (
	validHttpStatusCode = map[int]bool{
		http.StatusCreated:            true,
		http.StatusOK:                 true,
		http.StatusBadRequest:         true,
		http.StatusNotFound:           true,
		http.StatusPreconditionFailed: true,
		http.StatusForbidden:          true,
	}
)

