package errors

import (
	"errors"
)

type HandlerErrors struct {
	errorsChan chan error
	stopChan   chan bool
}

func NewHandlerErrors(
	errorsChan chan error,
	stopAppChan chan bool,
) *HandlerErrors {
	return &HandlerErrors{
		errorsChan: errorsChan,
		stopChan:   stopAppChan,
	}
}

func (h *HandlerErrors) Handle() {
	for {
		err := <-h.errorsChan
		var f *FatalAppError
		if errors.As(err, &f) {
			h.stopChan <- true
		}
	}
}
