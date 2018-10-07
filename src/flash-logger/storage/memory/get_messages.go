package memory

import (
	"flash-logger/model"
)

func (s *Storage) GetMessages(projectID int, limit int, offset int) []model.Message {
	messages := make([]model.Message, 0, limit)
	count := 0

	var tuples []Tuple
	if len(s.tuples) >= offset + limit {
		tuples = s.tuples[offset : offset + limit]
	} else if len(s.tuples) > offset {
		tuples = s.tuples[offset:]
	} else {
		tuples = make([]Tuple, 0, 0)
	}

	for _, tuple := range tuples {
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
					Timestamp: tuple.Message.Timestamp,
				},
			)
		}

		if count >= limit {
			break
		}
	}
	return messages
}
