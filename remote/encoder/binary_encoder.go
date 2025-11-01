package encoder

import (
	"bytes"
	"encoding/gob"
)

type BinaryEncoder[T any] struct {
}

func NewBinaryEncoder[T any]() *BinaryEncoder[T] {
	return &BinaryEncoder[T]{}
}

func (b *BinaryEncoder[T]) Encode(t T) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(t)
	return buf.Bytes(), err
}

func (b *BinaryEncoder[T]) Decode(bs []byte) (*T, error) {
	t := new(T)
	dec := gob.NewDecoder(bytes.NewReader(bs))
	return t, dec.Decode(t)
}
