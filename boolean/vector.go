package boolean

import(
	"typelib"
	"fmt"
	"strings"
)

type Vector struct {
	value	bool
	next	*Vector
}

func MakeVector(n int) (r *Vector) {
	if n > 0 {
		var node *Vector
		for ; n > 0; n-- {
			node = &Vector{ next: node }
		}
		r = node
	}
	return
}

func NewVector(n ...bool) (r *Vector) {
	if length := len(n); length > 0 {
		x := &Vector{ value: n[0] }
		end := x
		for _, value := range n[1:] {
			end.next = &Vector{ value: value }
			end = end.next
		}
		r = x
	}
	return
}

func (v *Vector) IsNil() bool {
	return v == nil
}

func (v *Vector) Content() (r interface{}) {
	if v != nil {
		r = v.value
	}
	return
}

func (v *Vector) Tail() (r typelib.List) {
	if v != nil {
		r = v.next
	} else {
		r = v
	}
	return
}

func (v *Vector) Fetch(i int) bool {
	var node *Vector
	for node = v; i > 0; i-- {
		if node.next == nil {
			node = nil
			break
		}
		node = node.next
	}
	return node.value
}

func (v *Vector) Node(i int) typelib.List {
	var node *Vector
	for node = v; i > 0; i-- {
		if node.next == nil {
			node = nil
			break
		}
		node = node.next
	}
	return node
}

func (v *Vector) At(i... interface{}) interface{} {
	x := i[0].(int)
	var node *Vector
	for node = v; x > 0; x-- {
		if node.next == nil {
			node = nil
			break
		}
		node = node.next
	}
	return node.value
}

func (v *Vector) Store(value interface{}, i... interface{}) {
	x := i[0].(int)
	var node *Vector
	for node = v; x > 0 && node.next != nil; x-- {
		node = node.next
	}
	if node != nil {
		switch value := value.(type) {
		case bool:
			node.value = value
		case []bool:
			switch len(value) {
			case 0:
			case 1:
				node.value = value[0]
				node.next = nil
			default:
				node.value = value[0]
				node.next = NewVector(value[1:]...)
			}
		case *Vector:
			node.value = value.value
			node.next = value.next
		default:
			panic(value)
		}
	}
}

func (v *Vector) End() (r typelib.List) {
	var node *Vector
	if node = v; node != nil {
		for ; node.next != nil; node = node.next {}
	}
	if node != nil {
		r = node
	} else {
		r = v
	}
	return
}

func (v *Vector) String() string {
	elements := []string{}
	v.Each(func(value bool) {
		elements = append(elements, fmt.Sprintf("%v", value))
	})
	return fmt.Sprintf("(list boolean (%v))", strings.Join(elements, " "))
}

func (v *Vector) equal(o *Vector) (r bool) {
	o_end := o
	for end := v; end != nil; {
		if o_end == nil || end.value != o_end.value {
			return false
		}
		end = end.next
		o_end = o_end.next
	}
	return o_end == nil
}

func (v *Vector) Equal(o interface{}) (r bool) {
	if o, ok := o.(*Vector); ok {
		r = v.equal(o)
	}
	return
}

func (v *Vector) Len() (i int) {
	for end := v; end != nil; i++ {
		end = end.next
	}
	return
}

func (v *Vector) Append(value interface{}) (r typelib.List) {
	var end *Vector

	if v != nil {
		end = v.End().(*Vector)
		switch value := value.(type) {
		case bool:
			end.next = &Vector{ value: value }
		case []bool:
			end.next = NewVector(value...)
		case *Vector:
			end.next = value
		default:
			panic(value)
		}
		r = v
	} else {
		switch value := value.(type) {
		case bool:
			r = &Vector{ value: value }
		case []bool:
			r = NewVector(value...)
		case *Vector:
			r = value
		default:
			panic(value)
		}
	}
	return
}

func (v *Vector) Prepend(value interface{}) (r typelib.List) {
	switch value := value.(type) {
	case bool:
		r = &Vector{ value: value }
	case []bool:
		if len(value) > 0 {
			r = NewVector(value...)
		}
	case *Vector:
		r = value
	default:
		panic(value)
	}

	if r == nil {
		r = v
	} else {
		v = r.Append(v).(*Vector)
	}
	return
}

func (v *Vector) Clone() (r typelib.List) {
	if v != nil {
		node := &Vector{ value: v.value }
		r = node
		for end := v.next; end != nil; {
			node.next = &Vector{ value: end.value }
			node = node.next
			end = end.next
		}
	} else {
		r = v
	}
	return
}

func (v *Vector) Each(f interface{}) {
	l := v
	switch f := f.(type) {
	case func(bool):
		for ; l != nil; l = l.next {
			f(l.value)
		}
	case func(int, bool):
		for i := 0; l != nil; l = l.next {
			f(i, l.value)
			i++
		}
	case func(interface{}, bool):
		for i := 0; l != nil; l = l.next {
			f(i, l.value)
			i++
		}
	case func(interface{}):
		for ; l != nil; l = l.next {
			f(l.value)
		}
	case func(int, interface{}):
		for i := 0; l != nil; l = l.next {
			f(i, l.value)
			i++
		}
	case func(interface{}, interface{}):
		for i := 0; l != nil; l = l.next {
			f(i, l.value)
			i++
		}
	case func(bool) bool:
		for ; l != nil; l = l.next {
			if !f(l.value) {
				break
			}
		}
	case func(int, bool) bool:
		for i := 0; l != nil; l = l.next {
			if f(i, l.value) {
				i++
			} else {
				break
			}
		}
	case func(interface{}, bool) bool:
		for i := 0; l != nil; l = l.next {
			if f(i, l.value) {
				i++
			} else {
				break
			}
		}
	case func(interface{}) bool:
		for ; l != nil; l = l.next {
			if !f(l.value) {
				break
			}
		}
	case func(int, interface{}) bool:
		for i := 0; l != nil; l = l.next {
			if f(i, l.value) {
				i++
			} else {
				break
			}
		}
	case func(interface{}, interface{}) bool:
		for i := 0; l != nil; l = l.next {
			if f(i, l.value) {
				i++
			} else {
				break
			}
		}
	}
}

func (v *Vector) Collect(f interface{}) (r typelib.List) {
	if l := v; l != nil {
		switch f := f.(type) {
		case func(bool) bool:
			end := &Vector{ value: f(l.value) }
			r = end
			for l = l.next; l != nil; l = l.next {
				end.next = &Vector{ value: f(l.value) }
				end = end.next
			}
		case func(int, bool) bool:
			end := &Vector{ value: f(0, l.value) }
			r = end
			l = l.next
			for i := 1; l != nil; l = l.next {
				end.next = &Vector{ value: f(i, l.value) }
				end = end.next
				i++
			}
		case func(interface{}, bool) bool:
			end := &Vector{ value: f(0, l.value) }
			r = end
			l = l.next
			for i := 1; l != nil; l = l.next {
				end.next = &Vector{ value: f(i, l.value) }
				end = end.next
				i++
			}
		}
	}
	return
}

func (v *Vector) Delete(f interface{}) (r typelib.List) {
	if l := v; l != nil {
		vector := new(Vector)
		end := vector
		switch f := f.(type) {
		case bool:
			for ; l != nil; l = l.next {
				if l.value != f {
					end.next = &Vector{ value: l.value }
					end = end.next
				}
			}
		case []bool:
			for _, o := range f {
				if l.value != o {
					end.next = &Vector{ value: l.value }
					end = end.next
				}
			}
		case *Vector:
			for o := f; o != nil; o = o.next {
				if l.value != o.value {
					end.next = &Vector{ value: l.value }
					end = end.next
				}
			}
		case func(bool) bool:
			for ; l != nil; l = l.next {
				if !f(l.value) {
					end.next = &Vector{ value: l.value }
					end = end.next
				}
			}
		case func(int, bool) bool:
			for i := 0; l != nil; l = l.next {
				if !f(i, l.value) {
					end.next = &Vector{ value: l.value }
					end = end.next
				}
				i++
			}
		case func(interface{}, bool) bool:
			for i := 0; l != nil; l = l.next {
				if !f(i, l.value) {
					end.next = &Vector{ value: l.value }
					end = end.next
				}
				i++
			}
		}
		r = vector.next
	} else {
		r = v
	}
	return
}

func (v *Vector) Reduce(f interface{}) (r typelib.List) {
	if l := v; l != nil {
		if f, ok := f.(func(bool, bool) bool); ok {
			x := l.value
			l = l.next
			for ; l != nil; l = l.next {
				x = f(x, l.value)
			}
			r = &Vector{ value: x }
		}
	}
	return
}

func (v *Vector) Reverse() (r typelib.List) {
	if v != nil {
		l := &Vector{ value: v.value }
		for end := v.next; end != nil; end = end.next {
			l = &Vector{ value: end.value, next: l }
		}
		r = l
	} else {
		r = v
	}
	return
}

func (v *Vector) Chain() (r *Chain) {
	r = new(Chain)
	if v != nil {
		r.start = typelib.NewNode(v.value)
		r.end = r.start
		r.length = 1
		for n := v.next; n != nil; n = n.next {
			r.end = r.end.Append(n.Content())
			r.length++
		}
	}
	return
}