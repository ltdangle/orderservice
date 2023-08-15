package util

import (
	"bytes"
	"encoding/gob"
)

type GobSerializer struct{}

func NewGobSerializer() *GobSerializer {
	return &GobSerializer{}
}
func (g *GobSerializer) EncodeToString(v interface{}) (string, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (g *GobSerializer) DecodeFromString(s string, v interface{}) error {
	buf := bytes.NewBufferString(s)
	dec := gob.NewDecoder(buf)
	return dec.Decode(v)
}
