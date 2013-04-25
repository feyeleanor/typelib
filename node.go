package typelib

type Node struct {
	value		interface{}
	previous	*Node
	next		*Node
}

func NewNode(v interface{}) *Node {
	return &Node{ value: v }
}

func (n *Node) Append(v interface{}) (r *Node) {
	r = &Node{ value: v }
	if n != nil {
		n.next = r
		r.previous = n
		r = n
	}
	return
}

func (n Node) Next() *Node {
	return n.next
}

func (n Node) Previous() *Node {
	return n.previous
}

func (n *Node) Fix() (end *Node, length int) {
	if n != nil {
		for end = n; end.next != nil; end = end.next {
			end.next.previous = end
			length++
		}
	}
	return
}

func (n *Node) Clone() (r *Node) {
	if n != nil {
		r = &Node{ value: n.value, previous: n.previous, next: n.next }
	}
	return 
}

func (n *Node) Start() (r *Node) {
	if n != nil {
		for r = n; r.previous != nil; r = r.previous {}
	}
	return
}

func (n *Node) End() (r *Node) {
	if n != nil {
		for r = n; r.next != nil; r = r.next {}
	}
	return
}

func (n *Node) Content() (r interface{}) {
	if n != nil {
		r = n.value
	}
	return
}

func (n *Node) Store(v interface{}) {
	if n != nil {
		n.value = v
	}
}

func (n *Node) Splice(o *Node) {
	switch {
	case n.next == nil:
		o.previous = n
		n.next = o
	case o.next == nil:
		o.previous = n
		o.next = n.next
		n.next = o
	default:
		o.previous = n
		o.End().next = n.next
		n.next = o
	}
}

func (n *Node) Truncate() {
	if n != nil {
		if n.next != nil {
			n.next.previous = nil
			n.next = nil
		}
	}
}

func (n *Node) Equal(o interface{}) (r bool) {
	if n != nil {
		switch o := o.(type) {
		case *Node:
			r = n.value == o.value
		default:
			r = n.value == o
		}
	} else {
		r = o == nil
	}
	return
}

func (n *Node) EachToStart(f func(*Node)) {
	for ; n != nil; n = n.previous {
		f(n)
	}
}

func (n *Node) EachToEnd(f func(*Node)) {
	for ; n != nil; n = n.next {
		f(n)
	}
}