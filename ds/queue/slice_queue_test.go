package queue_test

import (
	"testing"

	"github.com/gpahal/golib/ds/queue"
)

func TestSliceQueue(t *testing.T) {
	testInterfaceHelper(t, queue.NewSliceQueue)
}
