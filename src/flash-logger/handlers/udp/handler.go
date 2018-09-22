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
		priRegular: regexp.MustCompile(`^<\d+>`),

	}
}

// @example message
// https://tools.ietf.org/html/rfc3164
// <142>Sep 22 11:28:28 cs330699 nginx: 31.130.148.145 - - [22/Sep/2018:11:28:28 +0300] "GET / HTTP/1.1" 200 444 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.34
//  97.100 Safari/537.36"

func (h *Handler) Handle(buffer []byte) {
	pri := h.priRegular.Find(buffer[0:5])
	if pri == nil {
		log.Println("Not found PRI")
		return
	}

	timestamp := buffer[len(pri):len(pri) + 15]
	log.Printf("Timestamp: %s", string(timestamp))

	isHost := true
	host := make([]byte, 0, 10)
	process := make([]byte, 0, 10)
	newBuffer := buffer[len(pri) + len(timestamp) + 1:]
	length := 0
	for _, char := range newBuffer {
		length++
		if char == byte(':') {
			break
		}

		if char == byte(' ') {
			isHost = false
			continue
		}

		if isHost {
			host = append(host, char)
		} else {
			process = append(process, char)
		}
	}

	log.Printf("Host: %s", string(host))
	if len(process) > 0 {
		log.Printf("Process: %s", string(process))
	}
	newBuffer = newBuffer[length:]
	log.Println(string(newBuffer))
}
