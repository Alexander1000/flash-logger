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
	// ErrorUnauthorized
	ErrorUnauthorized = ErrorResult{
		Error: ErrorData{
			Message: "Not authorized",
			Code: http.StatusUnauthorized,
		},
	}
	// ErrorNotAllowed
	ErrorNotAllowed = ErrorResult{
		Error: ErrorData{
			Message: "Method not allowed",
			Code: http.StatusMethodNotAllowed,
		},
	}
	// ErrorInternalServerError
	ErrorInternalServerError = ErrorResult{
		Error: ErrorData{
			Message: "Internal server error",
			Code: http.StatusInternalServerError,
		},
	}
)
