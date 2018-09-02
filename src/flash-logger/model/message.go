package model

type Message struct {
	ID int `json:"id"`
	Level int `json:"level"`
	Message string `json:"message"`
	Context interface{} `json:"context"`
	Tags []string `json:"tags"`
}
