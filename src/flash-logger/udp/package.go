package udp

import (
	"log"
	"regexp"
	"fmt"
	"strconv"
)

type RELP struct {
	PRI int
	Timestamp string
	Host string
	Process string
	Message []byte
}

type RelpParser struct {
	priRegular *regexp.Regexp
}

func NewRelpParser() *RelpParser {
	return &RelpParser{
		priRegular: regexp.MustCompile(`^<\d+>`),
	}
}

func (rp *RelpParser) Scan(buffer []byte) (*RELP, error) {
	pri := rp.priRegular.Find(buffer[0:5])
	if pri == nil {
		log.Println("Not found PRI")
		return nil, fmt.Errorf("not found PRI: %s", string(buffer[0:5]))
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
	} else {
		return nil, fmt.Errorf("not found header with proceess")
	}
	newBuffer = newBuffer[length:]

	nPri, _ := strconv.ParseInt(string(pri[1:len(pri) - 1]), 10, 32)

	return &RELP{
		PRI: int(nPri),
		Timestamp: string(timestamp),
		Host: string(host),
		Process: string(process),
		Message: newBuffer,
	}, nil
}
