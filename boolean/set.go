package boolean

import(
	"fmt"
	"strings"
)

type _map	map[bool] bool

func (m _map) Len() int {
	return len(m)
}

func (m _map) Member(i interface{}) (r bool) {
	if i, ok := i.(bool); ok {
		r = m[i]
	}
	return
}

func (m _map) Include(v interface{}) {
	switch v := v.(type) {
	case []bool:
		for i := len(v) - 1; i > -1; i-- {
			m[v[i]] = true
		}
	case bool:
		m[v] = true
	default:
		panic(v)
	}
}

func (m _map) Each(f interface{}) {
	switch f := f.(type) {
	case func(bool):
		for k, v := range m {
			if v {
				f(k)
			}
		}
	case func(interface{}):
		for k, v := range m {
			if v {
				f(k)
			}
		}
	}
}

func (m _map) String() string {
	elements := []string{}
	m.Each(func(v bool) {
		elements = append(elements, fmt.Sprintf("%v", v))
	})
	return strings.Join(elements, " ")
}


type Set struct {
	_map
}

func NewSet(v... bool) (r Set) {
	r._map = make(_map)
	r.Include(v)
	return
}

func (s Set) String() string {
	return fmt.Sprintf("(boolean set (%v))", s._map.String())
}

func (s Set) Empty() Set {
	return NewSet()
}

func (s Set) Equal(o interface{}) (r bool) {
	if o, ok := o.(Set); ok {
		if r = s.Len() == o.Len(); r {
			s.Each(func(v bool) {
				if !o.Member(v) {
					r = false
				}
			})
		}
	}
	return
}