package model

type Message struct {
	ID int
	Level int
	Message string
	Context interface{}
	Tags []string
}
