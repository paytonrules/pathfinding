package graph

import (
	"testing"
)

func TestHasLocation(t *testing.T) {
	node := NewNavNode()

	if nil == node {
		t.Errorf("node is nil")
	}
}
