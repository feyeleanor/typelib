package boolean

import(
	"fmt"
	"strings"
	"typelib"
)

type Chain struct {
	start 		*typelib.Node
	end			*typelib.Node
	length		int
}

func MakeChain(n int) (r *Chain) {
	if n > 0 {
		node := typelib.NewNode(false)
		r = &Chain{ start: node, end: node, length: n }
		for n--; n > 0; n-- {
			r.end = r.end.Append(false)
		}
	}
	return
}

func NewChain(n... bool) (r *Chain) {
	if length := len(n); length > 0 {
		node := typelib.NewNode(n[0])
		r = &Chain{ start: node, end: node, length: length }
		for _, v := range n[1:] {
			r.end = r.end.Append(v)
		}
	} else {
		r = new(Chain)
	}
	return
}

func (c *Chain) IsNil() bool {
	return c == nil || c.start == nil || c.end == nil || c.length == 0
}

func (c *Chain) Fix() {
	if c != nil {
		switch {
		case c.length == 0:
			c.start = nil
			c.end = nil
		case c.start == nil:
			c.end = nil
			c.length = 0
		case c.end == nil:
			c.start = nil
			c.length = 0
		default:
			c.end, c.length = c.start.Fix()
		}
	}
}

func (c *Chain) fetchNode(i int) (n *typelib.Node) {
	if c != nil {
		for n = c.start; i > 0; i-- {
			if n.Next() == nil {
				n = nil
				break
			}
			n = n.Next()
		}
	}
	return
}

func (c *Chain) Fetch(i int) (r bool) {
	if n := c.fetchNode(i); n != nil {
		r = n.Content().(bool)
	}
	return
}

func (c *Chain) At(i... interface{}) (r interface{}) {
	if n := c.fetchNode(i[0].(int)); n != nil {
		r = n.Content().(bool)
	}
	return
}

func (c *Chain) Truncate(i int) {
	switch node := c.fetchNode(i); {
	case i == 0:
		c.start = nil
		c.end = nil
		c.length = 0
	case node != nil:
		if previous := node.Previous(); previous != nil {
			c.end = previous
			node.Truncate()
		} else {
			c.start = nil
			c.end = nil
		}
		c.length = i
	}
}

func (c *Chain) Store(value interface{}, i... interface{}) {
	x := i[0].(int)
	switch value := value.(type) {
	case bool:
		if node := c.fetchNode(x); node == nil {
		} else {
			node.Store(value)
		}
	case []bool:
		if len(value) == 0 {
			c.Truncate(x)
		} else {
			if node := c.fetchNode(x); node == nil {
				if x == 0 {
					chain := NewChain(value...)
					*c = *chain
				} else {
					panic(i)
				}
			} else {
				chain := NewChain(value...)
				node.Splice(chain.start)
				if chain.end.Next() == nil {
					c.end = chain.end
				}
				c.length += chain.length
			}
		}
	case *Vector:
		c.Store(value.Chain(), i...)
	case *Chain:
		if value.Len() == 0 {
			c.Truncate(x)
		} else {
			if node := c.fetchNode(x); node == nil {
				if x == 0 {
					chain := value.Clone().(*Chain)
					*c = *chain
				} else {
					panic(i)
				}
			} else {
				node.Splice(value.Clone().(*Chain).start)
				if value.end.Next() == nil {
					c.end = value.end
				}
				c.length += value.length
			}
		}
	default:
		panic(value)
	}
}

func (c Chain) String() string {
	elements := []string{}
	c.Each(func(v bool) {
		elements = append(elements, fmt.Sprintf("%v", v))
	})
	return fmt.Sprintf("(chain boolean (%v))", strings.Join(elements, " "))
}

func (c *Chain) equal(o *Chain) (r bool) {
	if c != nil && o != nil {
		o_end := o.start
		for end := c.start; end != nil; {
			if o_end == nil || !end.Equal(o_end) {
				return false
			}
			end = end.Next()
			o_end = o_end.Next()
		}
		return o_end == nil
	} else {
		return c == o
	}
}

func (c *Chain) Equal(o interface{}) (r bool) {
	if o, ok := o.(*Chain); ok {
		r = c.equal(o)
	}
	return
}

func (c Chain) Len() int {
	return c.length
}

func (c *Chain) Append(v interface{}) (r typelib.List) {
	if c != nil {
		var chain	*Chain

		end := c.end
		switch v := v.(type) {
		case bool:
			c.end = end.Append(v)
			c.length++
		case []bool:
			chain = NewChain(v...)
		case *Vector:
			chain = v.Chain()
		case *Chain:
			chain = v.Clone().(*Chain)
		default:
			panic(v)
		}

		if chain != nil {
			end.Splice(chain.start)
			c.length += chain.length
			c.end = chain.end
		}
		r = c
	} else {
		switch v := v.(type) {
		case bool:
			r = NewChain(v)
		case []bool:
			r = NewChain(v...)
		case *Vector:
			r = v.Chain()
		case *Chain:
			r = v.Clone().(*Chain)
		default:
			panic(v)
		}
	}
	return
}

func (c *Chain) Prepend(value interface{}) (r typelib.List) {
	switch value := value.(type) {
	case bool:
		r = NewChain(value)
	case []bool:
		r = NewChain(value...)
	case *Vector:
		r = value
	default:
		panic(value)
	}

	if r == nil {
		r = c
	} else {
		r = r.Append(value).(*Chain)
	}
	return
}

func (c *Chain) Clone() (r typelib.List) {
	if c != nil {
		if node := c.start.Clone(); node != nil {
			chain := &Chain{ start: node, length: 1 }
			for end := node.Next(); end != nil; end.Next() {
				node = node.Append(end.Content())
				chain.length++
			}
			chain.end = node
			r = chain
		}
	} else {
		r = new(Chain)
	}
	return
}

func (c Chain) Each(f interface{}) {
	if c.length > 0 {
		n := c.start
		switch f := f.(type) {
		case func(bool):
			for ; n != nil; n = n.Next() {
				f(n.Content().(bool))
			}
		case func(int, bool):
			for i := 0; n != nil; n = n.Next() {
				f(i, n.Content().(bool))
				i++
			}
		case func(interface{}, bool):
			for i := 0; n != nil; n = n.Next() {
				f(i, n.Content().(bool))
				i++
			}
		case func(interface{}):
			for ; n != nil; n = n.Next() {
				f(n.Content().(bool))
			}
		case func(int, interface{}):
			for i := 0; n != nil; n = n.Next() {
				f(i, n.Content().(bool))
				i++
			}
		case func(interface{}, interface{}):
			for i := 0; n != nil; n = n.Next() {
				f(i, n.Content().(bool))
				i++
			}
		}
	}
}

func (c *Chain) Collect(f interface{}) (r typelib.List) {
	chain := new(Chain)
	if c != nil && c.length > 0 {
		chain.length = c.length
		switch f := f.(type) {
		case func(bool) bool:
			end := c.start
			node := typelib.NewNode(f(end.Content().(bool)))
			chain.start = node
			for ; end != nil; end = end.Next() {
				node = node.Append(f(end.Content().(bool)))
			}
			chain.end = node
			r = chain
		case func(int, bool) bool:
			end := c.start
			node := typelib.NewNode(f(0, end.Content().(bool)))
			chain.start = node
			for i := 1; end != nil; end = end.Next() {
				node = node.Append(f(i, end.Content().(bool)))
				i++
			}
			chain.end = node
			r = chain
		case func(interface{}, bool) bool:
			end := c.start
			node := typelib.NewNode(f(0, end.Content().(bool)))
			chain.start = node
			for i := 1; end != nil; end = end.Next() {
				node = node.Append(f(i, end.Content().(bool)))
				i++
			}
			chain.end = node
			r = chain
		}
	}
	return
}

func (c *Chain) Delete(f interface{}) (r typelib.List) {
	if l := c; l != nil {
/*		switch f := f.(type) {
		case bool:

		case []bool:

		case *Chain:

		case func(bool) bool:

		case func(int, bool) bool:

		case func(interface{}, bool) bool:

		}
*/	} else {
		r = c
	}
	return
}

func (c *Chain) Reduce(f interface{}) (r typelib.List) {
	if c != nil {
		if f, ok := f.(func(bool, bool) bool); ok {
			n := c.start
			x := n.Content().(bool)
			for n = n.Next(); n != nil; n = n.Next() {
				x = f(x, n.Content().(bool))
			}
			n = typelib.NewNode(x)
			r = &Chain{ start: n, end: n, length: 1 }
		}
	}
	return
}

func (c *Chain) Reverse() (r typelib.List) {
	return
}
/*func (l DoublyLinkedList) ReverseEach(f interface{}) bool {
	if l.length > 0 {
		n := l.end
		switch f := f.(type) {
		case func(bool):
			for ; n.head != nil; {
				f(n.value)
				n = n.head
			}
		case func(int, bool):
			for i := 0; n.head != nil; i++ {
				f(i, n.value)
				n = n.head
			}
		case func(interface{}, bool):
			for i := 0; n.head != nil; i++ {
				f(i, n.value)
				n = n.head
			}
		case func(interface{}):
			for ; n.head != nil; {
				f(n.value)
				n = n.head
			}
		case func(int, interface{}):
			for i := 0; n.head != nil; i++ {
				f(i, n.value)
				n = n.head
			}
		case func(interface{}, interface{}):
			for i := 0; n.head != nil; i++ {
				f(i, n.value)
				n = n.head
			}
		default:
			return false
		}
		return true
	}
	return false
}*/