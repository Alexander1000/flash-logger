package udp

import (
	"net"
	"encoding/hex"
	"log"
)

type Listener struct {
	conn *net.UDPConn
}

func NewListener(conn *net.UDPConn) *Listener {
	return &Listener{conn: conn}
}

func (l *Listener) Listen() {
	for {
		var buf [2048]byte
		n, err := l.conn.Read(buf[0:])
		if err != nil {
			log.Println("Error Reading")
			return
		} else {
			log.Println(string(buf[0:n]))
			log.Println(hex.EncodeToString(buf[0:n]))
			log.Println("Package Done")
		}
	}
}
