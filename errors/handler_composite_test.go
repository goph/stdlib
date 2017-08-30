package errors_test

import (
	"testing"

	"github.com/goph/stdlib/errors"
	"github.com/goph/stdlib/internal/mocks"
)

func TestCompositeHandler(t *testing.T) {
	handler1 := new(mocks.Handler)
	handler2 := new(mocks.Handler)

	handler := errors.NewCompositeHandler(handler1, handler2)

	err := errors.New("error")

	handler1.On("Handle", err).Once().Return()
	handler2.On("Handle", err).Once().Return()

	handler.Handle(err)

	handler1.AssertExpectations(t)
	handler2.AssertExpectations(t)
}
