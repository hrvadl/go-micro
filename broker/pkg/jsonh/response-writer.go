package jsonh

import (
	"net/http"
)

type JSONResponseWriter interface {
	Write(w http.ResponseWriter, data ResponseEnvelope, status int, headers ...http.Header) error
}

type JSONResponseReader interface {
	Read(req *http.Request, dist any) error
}

type JSONResponseWriterReader interface {
	JSONResponseReader
	JSONResponseWriter
}

// JSON Response Writer Reader - JSONRWR
type JSONRWR struct {
	processor *JSONProcessor
}

func NewJSONRWR() *JSONRWR {
	return &JSONRWR{
		processor: NewJSONProcessor(),
	}
}

func (r *JSONRWR) Write(w http.ResponseWriter, data ResponseEnvelope, status int, headers ...http.Header) error {
	if err := r.processor.Encode(w, data); err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	return nil
}

func (r *JSONRWR) Read(req *http.Request, dist any) error {
	return nil
}
