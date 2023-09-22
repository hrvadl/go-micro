package jsonh

type ResponseEnvelope struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
	Data    any    `json:"data,omitempty"`
}
