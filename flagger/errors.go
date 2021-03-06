package flagger

import (
	"errors"
	"fmt"
	pkgErrors "github.com/pkg/errors"
)

type ErrNotSupported struct {
	err error
}

func NewErrNotSupported(msg ...interface{}) error {
	return new(ErrNotSupported).Init(msg...)
}

func (err *ErrNotSupported) Init(msg ...interface{}) error {
	err.err = pkgErrors.Wrap(errors.New("Not supported"), fmt.Sprint(msg...))
	return err
}

func (err ErrNotSupported) Error() string {
	return err.err.Error()
}
