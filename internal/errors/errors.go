package errors

import (
	"io"

	"github.com/demonoid81/golibav/avutil"
	"github.com/pkg/errors"
)

func Error(code int32) error {
	if code == 0 {
		return nil
	}

	switch avutil.Error(code) {
	case avutil.ErrEOF:
		return io.EOF

	case avutil.ErrNoMem:
		panic(avutil.ErrNoMem)
	}

	return errors.WithStack(avutil.Error(code))
}

func Return(code int32) (int, error) {
	if code > 0 {
		return int(code), nil
	}

	return 0, Error(code)
}
