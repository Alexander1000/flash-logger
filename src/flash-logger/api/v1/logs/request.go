package logs

type request struct {
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}