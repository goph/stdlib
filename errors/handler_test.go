package errors

import (
	"testing"

	"fmt"

	logrus "github.com/sirupsen/logrus/hooks/test"
)

func TestLogHandler_Handle_Logrus(t *testing.T) {
	logger, hook := logrus.NewNullLogger()
	handler := NewLogHandler(logger)

	err := fmt.Errorf("internal error")

	handler.Handle(err)

	if got, want := hook.LastEntry().Message, "internal error"; got != want {
		t.Fatalf("expected to log a specific error, received: %v", got)
	}
}
