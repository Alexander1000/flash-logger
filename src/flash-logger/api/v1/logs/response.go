package logs

import (
	"flash-logger/model"
)

type response struct {
	Result []model.Message `json:"result"`
}
