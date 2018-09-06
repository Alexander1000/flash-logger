package config

import (
	"os"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func (c *Config) LoadProjects() error {
	file, err := os.Open(c.ProjectList)
	if err != nil {
		return err
	}
	defer file.Close()
	// file.Read()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	projectList := ProjectList{}
	err = yaml.Unmarshal(bytes, &projectList)
	if err != nil {
		return err
	}
	c.Projects = projectList.Projects
	return nil
}
