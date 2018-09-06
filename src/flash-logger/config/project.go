package config

type Project struct {
	ID int `yaml:"id"`
	Title string `yaml:"title"`
	Name string `yaml:"name"`
	Token string `yaml:"token"`
}

type ProjectList struct {
	Projects []Project `yaml:"projects,flow"`
}
