package boolean

import(
	"testing"
)

func TestNodeAppend(t *testing.T) {
	n := &node{ value: false }
	switch r := n.Append(true); {
	case n.tail == nil:
		t.Fatalf("%v.Append(%v) failed to create a tail node", n, true)
	case !n.tail.value:
		t.Fatalf("%v.Append(%v) tail node value is %v", n, true, false)
	case r != n.tail:
		t.Fatalf("%v.Append(%v) returned node %v is not the same node as the tail", n, true, r)
	case r.head != n:
		t.Fatalf("%v.Append(%v) the head of returned node %v is not the same node as the head", n, true, r)
	case !r.value:
		t.Fatalf("%v.Append(%v) returned node value is %v", n, true, false)
	case r.head.value != false:
		t.Fatalf("%v.Append(%v) head node value is %v", n, true, true)
	}
}