package memory

type Storage struct {
	projects []Project
	tuples []Tuple
	sequenceMessageID int
}

func New() *Storage {
	// @todo доработать или загрузку из yml-файла или динамически создавать/сохранять
	projects := make([]Project, 0, 1)
	projects = append(projects, Project{ID: 1, Name: "demo", Key: "asdfg"})

	return &Storage{
		projects: projects,
		tuples: make([]Tuple, 0, 100),
		sequenceMessageID: 0,
	}
}
