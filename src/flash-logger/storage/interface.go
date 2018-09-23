package storage

import (
	"flash-logger/model"
)

const (
	SeverityEmergency = 0
	SeverityAlert = 1
	SeverityCritical = 2
	SeverityError = 3
	SeverityWarning = 4
	SeverityNotice = 5
	SeverityInformation = 6
	SeverityDebug = 7
)

type Repository interface {
	SaveMessage(projectID int, level int, message string, context interface{}, tags []string) error
	GetMessages(projectID int, limit int, offset int) []model.Message
}
