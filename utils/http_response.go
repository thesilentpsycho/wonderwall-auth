package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Object  string `json:"object"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

type Pagination struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

func WriteJSON(w http.ResponseWriter, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

func WriteJSONWithStatus(w http.ResponseWriter, statusCode int, v interface{}) error {
	w.WriteHeader(statusCode)
	return WriteJSON(w, v)
}

func WriteError(r *http.Request, w http.ResponseWriter, statusCode int, message string) error {
	w.WriteHeader(statusCode)
	return WriteJSON(w, &ErrorResponse{
		Object:  "error",
		Message: message,
	})
}
