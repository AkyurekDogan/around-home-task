/*
The Handler package to manage the request-response pipeline handlers
*/
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Base handles the base operations for handlers
type Base interface {
	WriteSuccessRespone(w http.ResponseWriter, status int, data interface{})
	WriteErrorRespone(w http.ResponseWriter, status int, message string, err error)
}

type httpResponse struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
	Error      any    `json:"error,omitempty"`
}

type base struct {
}

// WriteErrorRespone writes a JSON error response to the client
func (s *base) WriteErrorRespone(w http.ResponseWriter, status int, message string, err error) {
	w.WriteHeader(status)
	var errList []string
	if err != nil {
		errList = formatError(err)
	}
	r := httpResponse{
		StatusCode: status,
		Message:    fmt.Sprintf("[%s] %s", http.StatusText(status), message),
		Error:      errList,
	}

	json.NewEncoder(w).Encode(r)
}

// WriteSuccessRespone writes a JSON error response to the client
func (s *base) WriteSuccessRespone(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	r := httpResponse{
		StatusCode: status,
		Message:    http.StatusText(status),
		Data:       data,
	}

	json.NewEncoder(w).Encode(r)
}

func formatError(err error) []string {
	var result []string
	splitErr := strings.Split(err.Error(), "\n")
	for i, v := range splitErr {
		result = append(result, fmt.Sprintf("%d-%s", i, v))
	}
	return result
}