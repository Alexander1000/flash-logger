package storage

type Repository interface {
	SaveMessage(level string, message string, context interface{}, tags []string) error
}
