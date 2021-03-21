package server

import (
	"encoding/json"
	"net/http"
)

type ResponseJson struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"Error,omitempty"`
	Message  string      `json:"Message,omitempty"`
}

func NewSuccessResponse(writer http.ResponseWriter, data interface{}, statusCode int) {

	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(ResponseJson{
		Status: "success",
		Data:   data,
	})
}

func NewSuccessResponseWithMessage(writer http.ResponseWriter, data interface{}, statusCode int, message string) {

	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(ResponseJson{
		Status: "success",
		Data:   data,
		Message: message,
	})
}


func NewErrorResponse(writer http.ResponseWriter, error string, errorCode int) {
	writer.WriteHeader(errorCode)

	json.NewEncoder(writer).Encode(&ResponseJson{
		Status: "error",
		Error:  error,
	})
}
