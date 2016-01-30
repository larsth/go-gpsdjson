package gpsdjson

import "errors"

var (
	ErrNilByteSlice   = errors.New("<nil> byte slice input")
	ErrEmptyByteSlice = errors.New("non-nil, but empty byte slice")
)
