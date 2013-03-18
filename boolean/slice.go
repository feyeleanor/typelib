package boolean

import (
	"fmt"
	"strings"
	"typelib/integer"
)

type	Slice	[]bool

func NewSlice(v ...bool) (r Slice) {
	if len(v) > 0 {
		r = Slice(v)
	} else {
		r = Slice{}
	}
	return
}

func (s Slice) String() (t string) {
	elements := []string{}
	for _, v := range s {
		elements = append(elements, fmt.Sprintf("%v", v))
	}
	return fmt.Sprintf("(%v)", strings.Join(elements, " "))
}

func (s Slice) equal(o Slice) (r bool) {
	if len(s) == len(o) {
		r = true
		for i, v := range s {
			if r = v == o[i]; !r {
				return
			}
		}
	}
	return
}

func (s Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case Slice:
		r = s.equal(o)
	case []bool:
		r = s.equal(o)
	default:
	}
	return
}

func (s Slice) Clone() interface{} {
	x := make(Slice, len(s))
	copy(x, s)
	return x
}

func (s Slice) Len() int {
	return len(s)
}

func (s Slice) Cap() int {
	return cap(s)
}

func (s Slice) At(i ...int) (r interface{}) {
	if len(i) == 1 {
		r = s[i[0]]
	} else {
		n := make(Slice, 0, cap(i))
		for _, v := range i {
			n = append(n, s[v])
		}
		r = n
	}
	return
}

func (s Slice) Select(f interface{}) interface{} {
	r := make(Slice, 0)
	switch f := f.(type) {
	case func(bool) bool:
		for _, v := range s {
			if f(v) {
				r = append(r, v)
			}
		}
	case func(interface{}) bool:
		for _, v := range s {
			if f(v) {
				r = append(r, v)
			}
		}
	}
	return r
}

func (s Slice) Find(v interface{}, n ...int) integer.Slice {
	r := make(integer.Slice, 0, 0)
	if len(n) > 0 {
		limit := n[0]
		switch v := v.(type) {
		case bool:
			for j, x := range s {
				if x == v {
					if len(r) >= limit {
						break
					} else {
						r = append(r, j)
					}
				}
			}
		case func(bool) bool:
			for j, x := range s {
				if v(x) {
					if len(r) >= limit {
						break
					} else {
						r = append(r, j)
					}
				}
			}
		case func(interface{}) bool:
			for j, x := range s {
				if v(x) {
					if len(r) >= limit {
						break
					} else {
						r = append(r, j)
					}
				}
			}
		}
	} else {
		switch v := v.(type) {
		case bool:
			for j, x := range s {
				if x == v {
					r = append(r, j)
				}
			}
		case func(bool) bool:
			for j, x := range s {
				if v(x) {
					r = append(r, j)
				}
			}
		case func(interface{}) bool:
			for j, x := range s {
				if v(x) {
					r = append(r, j)
				}
			}
		}
	}
	return r
}

func (s Slice) Set(i, v interface{}) bool {
	if i, ok := i.(int); ok {
		switch v := v.(type) {
		case bool:
			s[i] = v
		case []bool:
			copy(s[i:], v)
		case Slice:
			copy(s[i:], v)
		default:
			return false
		}
		return true
	}
	return false
}

func (s Slice) Clear(n ...int) bool {
	switch limit := len(n); limit {
	case 0:
		for limit--; limit > -1; limit-- {
			s[limit] = false
		}
	case 1:
		s[n[0]] = false
	case 2:
		i := n[0]
		x := i + n[1]
		if x > len(s) {
			x = len(s)
		}
		for x--; x >= i; x-- {
			s[x] = false
		}
	default:
		return false
	}
	return true
}

func (s Slice) Swap(n ...int) bool {
	switch len(n) {
	case 2:
		i, j := n[0], n[1]
		s[i], s[j] = s[j], s[i]
	case 3:
		i, j := n[0], n[1]
		limit := n[2]
		temp := make(Slice, limit)
		copy(temp, s[i:i + limit])
		copy(s[i:], s[j:j + limit])
		copy(s[j:], temp)
	default:
		return false
	}
	return true
}

func (s Slice) Merge(o, f interface{}) bool {
	function := f.(func(l, r bool) bool)
	switch o := o.(type) {
	case bool:
		for l := len(s) - 1; l > -1; l-- {
			s[l] = function(s[l], o)
		}
	case []bool:
		limit := len(s)
		if len(o) < limit {
			limit = len(o)
		}
		for limit--; limit > -1; limit-- {
			s[limit] = function(s[limit], o[limit])
		}
	case Slice:
		limit := len(s)
		if len(o) < limit {
			limit = len(o)
		}
		for limit--; limit > -1; limit-- {
			s[limit] = function(s[limit], o[limit])
		}
	default:
		return false
	}
	return true
}

func (s Slice) Reduce(f interface{}) interface{} {
	if l := len(s); l > 0 {
		f := f.(func(seed, value bool) bool)
		limit := len(s)
		r := s[0]
		for x := 1; x < limit; x++ {
			r = f(r, s[x])
		}
		return r
	}
	return false
}

func (s Slice) True() (r bool) {
	r = true
	for i := len(s) - 1; r && i > -1; i-- {
		r = r && s[i]
	}
	return
}

func (s Slice) Negate(n ...int) bool {
	switch len(n) {
	case 0:
		for limit := len(s) - 1; limit > -1; limit-- {
			s[limit] = !s[limit]
		}
	case 1:
		s[n[0]] = !s[n[0]]
	case 2:
		i := n[0]
		limit := i + n[1]
		if limit > len(s) {
			limit = len(s)
		}
		for limit--; limit >= i; limit-- {
			s[limit] = !s[limit]
		}
	default:
		return false
	}
	return true
}

func (s Slice) And(o interface{}) bool {
	limit := len(s)
	switch o := o.(type) {
	case bool:
		for i := len(s) - 1; i > -1; i-- {
			s[i] = s[i] && o
		}
	case []bool:
		if limit > len(o) {
			limit = len(o)
		}
		for l := limit - 1; l > -1; l-- {
			s[l] = s[l] && o[l]
		}
	case Slice:
		if limit > len(o) {
			limit = len(o)
		}
		for l := limit - 1; l > -1; l-- {
			s[l] = s[l] && o[l]
		}
	default:
		return false
	}
	return true
}

func (s Slice) Or(o interface{}) bool {
	limit := len(s)
	switch o := o.(type) {
	case bool:
		for i := len(s) - 1; i > -1; i-- {
			s[i] = s[i] || o
		}
	case []bool:
		if limit > len(o) {
			limit = len(o)
		}
		for l := limit - 1; l > -1; l-- {
			s[l] = s[l] || o[l]
		}
	case Slice:
		if limit > len(o) {
			limit = len(o)
		}
		for l := limit - 1; l > -1; l-- {
			s[l] = s[l] || o[l]
		}
	default:
		return false
	}
	return true
}

func (s Slice) Not(o interface{}) bool {
	limit := len(s)
	switch o := o.(type) {
	case bool:
		for i := len(s) - 1; i > -1; i-- {
			s[i] = s[i] != o
		}
	case []bool:
		if limit > len(o) {
			limit = len(o)
		}
		for l := limit - 1; l > -1; l-- {
			s[l] = s[l] != o[l]
		}
	case Slice:
		if limit > len(o) {
			limit = len(o)
		}
		for l := limit - 1; l > -1; l-- {
			s[l] = s[l] != o[l]
		}
	default:
		return false
	}
	return true
}

func (s *Slice) RestrictTo(i, j int) {
	*s = (*s)[i:j]
}

func (s *Slice) Cut(i, j int) {
	a := *s
	l := len(a)
	if i < 0 {
		i = 0
	}
	if j > l {
		j = l
	}
	if j > i {
		l -= j - i
		copy(a[i:], a[j:])
		*s = a[:l]
	}
}

func (s *Slice) Trim(i, j int) {
	a := *s
	n := len(a)
	if i < 0 {
		i = 0
	}
	if j > n {
		j = n
	}
	if j > i {
		copy(a, a[i:j])
		*s = a[:j - i]
	}
}

func (s *Slice) Insert(i int, v interface{}) bool {
	switch v := v.(type) {
	case bool:
		l := s.Len() + 1
		n := make(Slice, l, l)
		copy(n, (*s)[:i])
		n[i] = v
		copy(n[i + 1:], (*s)[i:])
		*s = n
	case Slice:
		l := s.Len() + len(v)
		n := make(Slice, l, l)
		copy(n, (*s)[:i])
		copy(n[i:], v)
		copy(n[i + len(v):], (*s)[i:])
		*s = n
	case []bool:
		s.Insert(i, Slice(v))
	default:
		return false
	}
	return true
}

func (s *Slice) Delete(n... int) bool {
	a := *s
	limit := len(a)
	switch len(n) {
	case 1:
		i := n[0]
		if i > -1 && i < limit {
			copy(a[i:], a[i + 1:])
			*s = a[:limit - 1]
		}
	case 2:
		i, count := n[0], n[1]
		limit -= count
		if i > -1 && i <= limit {
			copy(a[i:], a[i + count:])
			*s = a[:limit]
		}
	default:
		return false
	}
	return true
}

func (s *Slice) DeleteIf(f interface{}) bool {
	a := *s
	p := 0
	switch f := f.(type) {
	case bool:
		for i, v := range a {
			if i != p {
				a[p] = v
			}
			if v != f {
				p++
			}
		}
	case func(bool) bool:
		for i, v := range a {
			if i != p {
				a[p] = v
			}
			if !f(v) {
				p++
			}
		}
	case func(interface{}) bool:
		for i, v := range a {
			if i != p {
				a[p] = v
			}
			if !f(v) {
				p++
			}
		}

	default:
		return false
	}
	*s = a[:p]
	return true
}

func (s *Slice) KeepIf(f interface{}) bool {
	a := *s
	p := 0
	switch f := f.(type) {
	case bool:
		for i, v := range a {
			if i != p {
				a[p] = v
			}
			if v == f {
				p++
			}
		}
	case func(bool) bool:
		for i, v := range a {
			if i != p {
				a[p] = v
			}
			if f(v) {
				p++
			}
		}
	case func(interface{}) bool:
		for i, v := range a {
			if i != p {
				a[p] = v
			}
			if f(v) {
				p++
			}
		}
	default:
		return false
	}
	*s = a[:p]
	return true
}

func (s Slice) Each(f interface{}) bool {
	switch f := f.(type) {
	case func(bool):
		for _, v := range s {
			f(v)
		}
	case func(int, bool):
		for i, v := range s {
			f(i, v)
		}
	case func(interface{}, bool):
		for i, v := range s {
			f(i, v)
		}
	case func(interface{}):
		for _, v := range s {
			f(v)
		}
	case func(int, interface{}):
		for i, v := range s {
			f(i, v)
		}
	case func(interface{}, interface{}):
		for i, v := range s {
			f(i, v)
		}
	default:
		return false
	}
	return true
}

func (s Slice) ReverseEach(f interface{}) bool {
	switch f := f.(type) {
	case func(bool):
		for i := len(s) - 1; i > -1; i-- {
			f(s[i])
		}
	case func(int, bool):
		for i := len(s) - 1; i > -1; i-- {
			f(i, s[i])
		}
	case func(interface{}, bool):
		for i := len(s) - 1; i > -1; i-- {
			f(i, s[i])
		}
	case func(interface{}):
		for i := len(s) - 1; i > -1; i-- {
			f(s[i])
		}
	case func(int, interface{}):
		for i := len(s) - 1; i > -1; i-- {
			f(i, s[i])
		}
	case func(interface{}, interface{}):
		for i := len(s) - 1; i > -1; i-- {
			f(i, s[i])
		}
	default:
		return false
	}
	return true
}

func (s Slice) While(f interface{}) int {
	switch f := f.(type) {
	case func(interface{}) bool:
		for i, v := range s {
			if !f(v) {
				return i
			}
		}
	case func(bool) bool:
		for i, v := range s {
			if !f(v) {
				return i
			}
		}
	case func(int, interface{}) bool:
		for i, v := range s {
			if !f(i, v) {
				return i
			}
		}
	case func(int, bool) bool:
		for i, v := range s {
			if !f(i, v) {
				return i
			}
		}
	case func(interface{}, interface{}) bool:
		for i, v := range s {
			if !f(i, v) {
				return i
			}
		}
	case func(interface{}, bool) bool:
		for i, v := range s {
			if !f(i, v) {
				return i
			}
		}
	default:
		return 0
	}
	return len(s)
}

func (s Slice) Until(f interface{}) int {
	switch f := f.(type) {
	case func(interface{}) bool:
		for i, v := range s {
			if f(v) {
				return i
			}
		}
	case func(bool) bool:
		for i, v := range s {
			if f(v) {
				return i
			}
		}
	case func(int, interface{}) bool:
		for i, v := range s {
			if f(i, v) {
				return i
			}
		}
	case func(int, bool) bool:
		for i, v := range s {
			if f(i, v) {
				return i
			}
		}
	case func(interface{}, interface{}) bool:
		for i, v := range s {
			if f(i, v) {
				return i
			}
		}
	case func(interface{}, bool) bool:
		for i, v := range s {
			if f(i, v) {
				return i
			}
		}
	default:
		return 0
	}
	return len(s)
}

func (s Slice) Reverse() {
	end := len(s) - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s *Slice) Repeat(count int) {
	if count > 1 {
		a := *s
		length := len(a) * count
		capacity := cap(a)
		if capacity < length {
			capacity = length
		}
		destination := make(Slice, length, capacity)
		for start, end := 0, len(a); count > 0; count-- {
			copy(destination[start:end], a)
			start = end
			end += len(a)
		}
		*s = destination
	}
}

func (s Slice) Uniq() interface{} {
	return NewSet(s...)
}

func (s *Slice) Reallocate(n ...int) bool {
	switch len(n) {
	case 1:
		length := n[0]
		if length > cap(*s) {
			x := make(Slice, length)
			copy(x, *s)
			*s = x
		} else {
			*s = (*s)[:length]
		}
	case 2:
		length, capacity := n[0], n[1]
		switch {
		case length > capacity:
			s.Reallocate(capacity, capacity)
		case capacity != cap(*s):
			x := make(Slice, length, capacity)
			copy(x, *s)
			*s = x
		default:
			*s = (*s)[:length]
		}
	default:
		return false
	}
	return true
}

func (s *Slice) Extend(n int) bool {
	if n > 0 {
		c := cap(*s)
		l := len(*s) + n
		if l > c {
			c = l
		}
		return s.Reallocate(l, c)
	}
	return false
}