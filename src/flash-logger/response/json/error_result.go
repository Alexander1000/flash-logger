package json

import "net/http"

type ErrorResult struct {
	Error ErrorData `json:"error"`
}

type ErrorData struct {
	Message string `json:"message"`
	Code int `json:"code"`
}

var (
	// ErrorNotAllowed
	ErrorNotAllowed = ErrorResult{
		Error: ErrorData{
			Message: "Method not allowed",
			Code: http.StatusMethodNotAllowed,
		},
	}
)
