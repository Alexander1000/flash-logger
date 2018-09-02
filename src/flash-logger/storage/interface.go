package storage

import (
	"flash-logger/model"
)

type Repository interface {
	SaveMessage(projectID int, level int, message string, context interface{}, tags []string) error
	GetLastMessages(projectID int, limit int, offset int) []model.Message
}
