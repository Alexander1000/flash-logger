package event

import (
	"net/http"
	"encoding/json"
	"io/ioutil"

	jsonResponse "flash-logger/response/json"
	"flash-logger/storage"
)

type Handler struct {
	storage storage.Repository
}

func New(storage storage.Repository) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	if req.Method != "POST" {
		jsonResponse.Reply(resp, jsonResponse.ErrorNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	// @todo добавить validator

	logData := request{}
	if reqData, err := ioutil.ReadAll(req.Body); err != nil {
		jsonResponse.Reply(resp, jsonResponse.ErrorInternalServerError, http.StatusInternalServerError)
		return
	} else if err := json.Unmarshal(reqData, &logData); err != nil {
		jsonResponse.Reply(resp, jsonResponse.ErrorInternalServerError, http.StatusInternalServerError)
		return
	}

	// @todo определять projectId по key
	// @todo конвертировать строку в число fatal -> 0, error -> 1, etc
	if err := h.storage.SaveMessage(1, 1, logData.Message, logData.Context, logData.Tags); err != nil {
		jsonResponse.Reply(resp, jsonResponse.ErrorInternalServerError, http.StatusInternalServerError)
		return
	}

	jsonResponse.Reply(resp, response{Result: true}, http.StatusOK)
}
