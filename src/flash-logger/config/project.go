package config

type Project struct {
	ID int `yaml:"id"`
	Title string `yaml:"title"`
	Name string `yaml:"name"`
	Token string `yaml:"token"`
	UDP bool `yaml:"udp,omitempty"`
}

type ProjectList struct {
	Projects []Project `yaml:"projects,flow"`
}
