package udp

type Handler interface {
	Handle([]byte)
}

