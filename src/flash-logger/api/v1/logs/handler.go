package logs

import (
	"net/http"
	"io/ioutil"
	"encoding/json"

	jsonResponse "flash-logger/response/json"
	"flash-logger/storage"
)

const (
	defaultLimit = 20
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

	// @todo добавить validator, добавить авторизацию по заголовку Bearer

	requestData := request{}
	if reqData, err := ioutil.ReadAll(req.Body); err != nil {
		jsonResponse.Reply(resp, jsonResponse.ErrorInternalServerError, http.StatusInternalServerError)
		return
	} else if err := json.Unmarshal(reqData, &requestData); err != nil {
		jsonResponse.Reply(resp, jsonResponse.ErrorInternalServerError, http.StatusInternalServerError)
		return
	}

	if requestData.Limit == 0 {
		requestData.Limit = defaultLimit
	}

	valProjectID := req.Context().Value("projectId")
	projectID := valProjectID.(int)

	messages := h.storage.GetMessages(projectID, requestData.Limit, requestData.Offset)
	jsonResponse.Reply(resp, response{Result: messages}, http.StatusOK)
}
