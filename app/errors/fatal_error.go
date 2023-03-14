package errors

import "fmt"

type FatalAppError struct {
	Err error
}

func NewFatalError(err error) error {
	return &FatalAppError{
		Err: err,
	}
}

func (f *FatalAppError) Error() string {
	return fmt.Sprintf("%v %v", "Fatal error: ", f.Err)
}
