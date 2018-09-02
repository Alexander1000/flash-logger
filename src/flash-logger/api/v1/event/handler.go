package event

import (
	"net/http"
	"encoding/json"
	"io/ioutil"

	jsonResponse "flash-logger/response/json"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		jsonResponse.Reply(resp, jsonResponse.ErrorNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	// @todo добавить validator, добавить авторизацию по заголовку Bearer

	logData := request{}
	if reqData, err := ioutil.ReadAll(req.Body); err != nil {
		jsonResponse.Reply(resp, jsonResponse.ErrorInternalServerError, http.StatusInternalServerError)
		return
	} else if err := json.Unmarshal(reqData, &logData); err != nil {
		jsonResponse.Reply(resp, jsonResponse.ErrorInternalServerError, http.StatusInternalServerError)
		return
	}

	jsonResponse.Reply(resp, response{Result: true}, http.StatusOK)
}
