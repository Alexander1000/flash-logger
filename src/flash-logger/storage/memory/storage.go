package memory

import "flash-logger/config"

type Storage struct {
	projects []config.Project
	tuples []Tuple
	sequenceMessageID int
}

func New(projects []config.Project) *Storage {
	return &Storage{
		projects: projects,
		tuples: make([]Tuple, 0, 100),
		sequenceMessageID: 0,
	}
}
