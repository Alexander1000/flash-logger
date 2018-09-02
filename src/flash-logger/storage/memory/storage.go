package memory

type Storage struct {
	projects []Project
}

func New() *Storage {
	// @todo доработать или загрузку из yml-файла или динамически создавать/сохранять
	projects := make([]Project, 0, 1)
	projects = append(projects, Project{ID: 1, Name: "demo", Key: "asdfg"})

	return &Storage{
		projects: projects,
	}
}
