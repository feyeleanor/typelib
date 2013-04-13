package boolean

type node struct {
	value	bool
	head	*node
	tail	*node
}

func (n *node) Append(v bool) *node {
	n.tail = &node{ value: v, head: n }
	return n.tail
}

func (n node) Next() *node {
	return n.tail
}

func (n node) Previous() *node {
	return n.head
}