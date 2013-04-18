package boolean

import(
	"fmt"
	"strings"
)

type List struct {
	value	bool
	next	*List
}

func MakeList(n int) (r *List) {
	if n > 0 {
		for ; n > 0; n-- {
			r = &List{ next: r }
		}
	}
	return
}

func NewList(n ...bool) (r *List) {
	if length := len(n); length > 0 {
		r = &List{ value: n[0] }
		end := r
		for _, v := range n[1:] {
			end.next = &List{ value: v }
			end = end.next
		}
	}
	return
}

func (s *List) At(i interface{}) (r *List) {
	x := i.(int)
	for r = s; x > 0 && r.next != nil; x-- {
		r = r.next
	}
	return
}

func (s *List) Set(i, v interface{}) (r *List) {
	x := i.(int)
	for r = s; x > 0 && r.next != nil; x-- {
		r = r.next
	}
	if r != nil {
		switch v := v.(type) {
		case bool:
			r.value = v
		case []bool:
			switch len(v) {
			case 0:
			case 1:
				r.value = v[0]
				r.next = nil
			default:
				r.value = v[0]
				r.next = NewList(v[1:]...)
			}
		case *List:
		default:
			panic(v)
		}
	}
	return
}

func (s *List) End() (r *List) {
	if r = s; r != nil {
		for ; r.next != nil; r = r.next {}
	}
	return
}

func (s *List) String() string {
	elements := []string{}
	s.Each(func(v bool) {
		elements = append(elements, fmt.Sprintf("%v", v))
	})
	return fmt.Sprintf("(list boolean (%v))", strings.Join(elements, " "))
}

func (s *List) equal(o *List) (r bool) {
	o_end := o
	for end := s; end != nil; {
		if o_end == nil || end.value != o_end.value {
			return false
		}
		end = end.next
		o_end = o_end.next
	}
	return o_end == nil
}

func (s *List) Equal(o interface{}) (r bool) {
	if o, ok := o.(*List); ok {
		r = s.equal(o)
	}
	return
}

func (s *List) Len() (i int) {
	for end := s; end != nil; i++ {
		end = end.next
	}
	return
}

func (s *List) Append(v interface{}) (r *List) {
	if end := s.End(); end != nil {
		switch v := v.(type) {
		case bool:
			end.next = &List{ value: v }
		case []bool:
			end.next = NewList(v...)
		case *List:
			end.next = v
		}
		r = s
	} else {
		switch v := v.(type) {
		case bool:
			r = &List{ value: v }
		case []bool:
			r = NewList(v...)
		case *List:
			r = v
		}
	}
	return
}

func (s *List) Prepend(v interface{}) (r *List) {
	switch v := v.(type) {
	case bool:
		r = &List{ value: v }
	case []bool:
		r = NewList(v...)
	case *List:
		r = v
	}

	if r == nil {
		r = s
	} else {
		s = r.Append(s)
	}
	return
}

func (s *List) Clone() (r *List) {
	if s != nil {
		r = &List{ value: s.value }
		l := r
		for end := s.next; end != nil; {
			l.next = &List{ value: end.value }
			l = l.next
			end = end.next
		}
	}
	return
}

func (s *List) Each(f interface{}) {
	l := s
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

func (s *List) Collect(f interface{}) (r *List) {
	if l := s; l != nil {
		switch f := f.(type) {
		case func(bool) bool:
			r = &List{ value: f(l.value) }
			end := r
			for l = l.next; l != nil; l = l.next {
				end.next = &List{ value: f(l.value) }
				end = end.next
			}
		case func(int, bool) bool:
			r = &List{ value: f(0, l.value) }
			end := r
			l = l.next
			for i := 1; l != nil; l = l.next {
				end.next = &List{ value: f(i, l.value) }
				end = end.next
				i++
			}
		case func(interface{}, bool) bool:
			r = &List{ value: f(0, l.value) }
			end := r
			l = l.next
			for i := 1; l != nil; l = l.next {
				end.next = &List{ value: f(i, l.value) }
				end = end.next
				i++
			}
		}
	}
	return
}

func (s *List) Delete(f interface{}) (r *List) {
	if l := s; l != nil {
		r = &List{ value: false }
		end := r
		switch f := f.(type) {
		case bool:
			for ; l != nil; l = l.next {
				if l.value != f {
					end.next = &List{ value: l.value }
					end = end.next
				}
			}
		case []bool:
			for _, o := range f {
				if l.value != o {
					end.next = &List{ value: l.value }
					end = end.next
				}
			}
		case *List:
			for o := f; o != nil; o = o.next {
				if l.value != o.value {
					end.next = &List{ value: l.value }
					end = end.next
				}
			}
		case func(bool) bool:
			for ; l != nil; l = l.next {
				if !f(l.value) {
					end.next = &List{ value: l.value }
					end = end.next
				}
			}
		case func(int, bool) bool:
			for i := 0; l != nil; l = l.next {
				if !f(i, l.value) {
					end.next = &List{ value: l.value }
					end = end.next
				}
				i++
			}
		case func(interface{}, bool) bool:
			for i := 0; l != nil; l = l.next {
				if !f(i, l.value) {
					end.next = &List{ value: l.value }
					end = end.next
				}
				i++
			}
		}
		r = r.next
	}
	return
}

func (s *List) Reduce(f interface{}) (r bool) {
	if l := s; l != nil {
		if f, ok := f.(func(bool, bool) bool); ok {
			r = l.value
			l = l.next
			for ; l != nil; l = l.next {
				r = f(r, l.value)
			}
		}
	}
	return
}

func (s *List) Reverse() (l *List) {
	if s != nil {
		l = &List{ value: s.value }
		for end := s.next; end != nil; end = end.next {
			l = &List{ value: end.value, next: l }
		}
	}
	return
}