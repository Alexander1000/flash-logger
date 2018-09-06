package config

import "os"

func (c *Config) LoadProjects() error {
	file, err := os.Open(c.ProjectList)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
