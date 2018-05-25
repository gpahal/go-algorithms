package queue_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/ds/queue"
)

func TestSliceQueue(t *testing.T) {
	testInterfaceHelper(t, queue.NewSliceQueue)
}
