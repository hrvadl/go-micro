package jsonh

import (
	"encoding/json"
	"io"
)

const maxBytes = 1000000

type JSONProcessor struct{}

func NewJSONProcessor() *JSONProcessor {
	return &JSONProcessor{}
}

func (p *JSONProcessor) Decode(r io.Reader, dest any) error {
	dec := json.NewDecoder(io.LimitReader(r, maxBytes))

	if err := dec.Decode(dest); err != nil {
		return err
	}

	return nil
}

func (p *JSONProcessor) Encode(w io.Writer, data any) error {
	dec := json.NewEncoder(w)

	if err := dec.Encode(data); err != nil {
		return err
	}

	return nil
}
