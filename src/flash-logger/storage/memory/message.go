package memory

type Message struct {
	Level int
	Message string
	Context interface{}
	Tags []string
}
