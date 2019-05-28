package queue_test

import (
	"testing"

	"github.com/gpahal/golib/ds/queue"
)

func TestListQueue(t *testing.T) {
	testInterfaceHelper(t, queue.NewListQueue)
}
