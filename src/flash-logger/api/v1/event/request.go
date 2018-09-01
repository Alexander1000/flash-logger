package event

type request struct {
	Key string `json:"key"`
	Level string `json:"level"`
	Message string `json:"message"`
	Context interface{} `json:"context"`
	Tags []string `json:"tags"`
}
