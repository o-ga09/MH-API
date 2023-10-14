package middleware

import "runtime"

type stackError struct {
	stack []byte
	err error
}

func NewStackError(err error) *stackError {
	var buf [16 * 1024]byte
	n := runtime.Stack(buf[:],false)
	return &stackError{
		stack: buf[:n],
		err: err,
	}
}

func (e stackError) Error() string {
	return e.err.Error()
}

func (e *stackError) Unwrap() error {
	return e.err
}