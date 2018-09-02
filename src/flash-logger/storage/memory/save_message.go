package memory

func (s *Storage) SaveMessage(projectID int, level int, message string, context interface{}, tags []string) error {
	s.sequenceMessageID++
	sequenceID := s.sequenceMessageID
	tMessage := Message{
		ID: sequenceID,
		Level: level,
		Message: message,
		Context: context,
		Tags: tags,
	}
	tuple := Tuple{
		ProjectID: projectID,
		Message: &tMessage,
	}
	s.tuples = append(s.tuples, tuple)
	return nil
}
