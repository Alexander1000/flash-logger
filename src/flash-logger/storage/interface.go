package storage

type Repository interface {
	SaveMessage(projectID int, level int, message string, context interface{}, tags []string) error
}
