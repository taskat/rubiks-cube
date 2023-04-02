package httpserver

type Request struct {
	Config    string `json:"config"`
	Algorithm string `json:"algorithm"`
}
