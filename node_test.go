package typelib

import(
	"testing"
)

func TestNodeAppend(t *testing.T) {
	n := &Node{ value: false }
	switch n.Append(true); {
	case n.next == nil:
		t.Fatalf("%v.Append(%v) failed to create a tail node", n, true)
	case n.next.value == false:
		t.Fatalf("%v.Append(%v) tail node value is %v", n, true, false)
	case n.next.previous != n:
		t.Fatalf("%v.Append(%v) returned node %v is not the same node as the head node", n, true, n)
	}
}