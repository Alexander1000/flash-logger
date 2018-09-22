package udp

import "net"

type Listener struct {
	conn *net.UDPConn
}

func NewListener(conn *net.UDPConn) *Listener {
	return &Listener{conn: conn}
}

func (l *Listener) Listen() {

}
