package udp

import "log"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(buffer []byte) {
	log.Println(string(buffer))
}
