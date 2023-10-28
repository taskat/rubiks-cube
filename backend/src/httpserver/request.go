package httpserver

type Request struct {
	Config string `json:"config"`
	Algo   string `json:"algo"`
}
