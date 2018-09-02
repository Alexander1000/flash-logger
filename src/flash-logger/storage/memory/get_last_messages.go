package memory

import (
	"flash-logger/model"
)

func (s *Storage) GetLastMessages(projectID int, limit int, offset int) []model.Message {
	messages := make([]model.Message, 0, limit)
	count := 0
	for _, tuple := range s.tuples {
		if projectID == tuple.ProjectID {
			count++
			messages = append(
				messages,
				model.Message{
					ID: tuple.Message.ID,
					Level: tuple.Message.Level,
					Message: tuple.Message.Message,
					Context: tuple.Message.Context,
					Tags: tuple.Message.Tags,
				},
			)
		}

		if count >= limit {
			break
		}
	}
	return messages
}
