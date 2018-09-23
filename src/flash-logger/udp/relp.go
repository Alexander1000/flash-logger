package udp

type RELP struct {
	PRI int
	Timestamp string
	Host string
	Process string
	Message []byte
}

func (r *RELP) GetSeverity() int {
	return r.PRI % 8
}
