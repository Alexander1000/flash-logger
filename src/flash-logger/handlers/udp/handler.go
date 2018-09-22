package udp

import (
	"log"
	"regexp"
)

type Handler struct {
	priRegular *regexp.Regexp
}

func New() *Handler {
	return &Handler{
		priRegular: regexp.MustCompile(`/^<\d+>/`),
	}
}

// @example message
// https://tools.ietf.org/html/rfc3164
// <142>Sep 22 11:28:28 cs330699 nginx: 31.130.148.145 - - [22/Sep/2018:11:28:28 +0300] "GET / HTTP/1.1" 200 444 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.34
//  97.100 Safari/537.36"

func (h *Handler) Handle(buffer []byte) {
	// parsing PRI part
	log.Println(string(buffer[0:5]))
	pri := h.priRegular.Find(buffer[0:5])
	if pri != nil {
		log.Println(string(pri))
	}
	log.Println(string(buffer))
}
