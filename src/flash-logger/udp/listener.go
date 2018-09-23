package udp

import (
	"net"
	"log"
)

type Listener struct {
	conn *net.UDPConn
	handler Handler
}

func NewListener(conn *net.UDPConn, handler Handler) *Listener {
	return &Listener{conn: conn, handler: handler}
}

func (l *Listener) Listen() {
	for {
		var buf [2048]byte
		n, err := l.conn.Read(buf[0:])
		if err != nil {
			log.Println("Error Reading")
			return
		} else {
			l.handler.Handle(buf[0:n])
		}
	}
}
