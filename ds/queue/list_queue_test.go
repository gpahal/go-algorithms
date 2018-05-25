package queue_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/ds/queue"
)

func TestListQueue(t *testing.T) {
	testInterfaceHelper(t, queue.NewListQueue)
}
