package json

import (
	"net/http"
	"encoding/json"
)

func Reply(w http.ResponseWriter, data interface{}, code int) (int, error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	body, err := json.Marshal(data)
	if err != nil {
		body = []byte(
			`{"error": "Unknown and unpredictable error with huge, massive and catastrophic consequences!"}`)
	}

	return w.Write(body)
}
