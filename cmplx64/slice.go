package cmplx64

import (
	"typelib/boolean"
	"fmt"
	"math/cmplx"
	"strings"
)

type	Slice	[]complex64

const	ZERO = complex64(0)

func (s Slice) String() string {
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
	case Slice:			r = s.equal(o)
	case []complex64:		r = s.equal(o)
	}
	return
}

func (s Slice) Clone() interface{} {
	x := make(Slice, len(s))
	copy(x, s)
	return x
}

func (s Slice) Merge(o, f interface{}) {
	function := f.(func(l, r complex64) complex64)
	switch o := o.(type) {
	case complex64:
		for l := len(s) - 1; l > -1; l-- {
			s[l] = function(s[l], o)
		}
	case []complex64:
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
		panic(o)
	}
}

func (s Slice) Reduce(f interface{}) interface{} {
	if l := len(s); l > 0 {
		f := f.(func(seed, value complex64) complex64)
		limit := len(s)
		r := s[0]
		for x := 1; x < limit; x++ {
			r = f(r, s[x])
		}
		return r
	}
	return ZERO
}

func (s Slice) Sum() interface{} {
	if l := len(s); l > 0 {
		r := s[0]
		for x := len(s) - 1; x > 0; x-- {
			r += s[x]
		}
		return r
	}
	return ZERO
}

func (s Slice) Product() interface{} {
	if l := len(s); l > 0 {
		r := s[0]
		for x := len(s) - 1; x > 0; x-- {
			r *= s[x]
		}
		return r
	}
	return ZERO
}

func (s Slice) DotProduct(o interface{}) (r interface{}) {
	limit := len(s)
	switch o := o.(type) {
	case []complex64:
		if limit > len(o) {
			limit = len(o)
		}
		x := make(Slice, limit)
		for limit--; limit > -1; limit-- {
			x[limit] = s[limit] * o[limit]
		}
		r = x.Sum()
	case Slice:
		if limit > len(o) {
			limit = len(o)
		}
		x := make(Slice, limit)
		for limit--; limit > -1; limit-- {
			x[limit] = s[limit] * o[limit]
		}
		r = x.Sum()
	default:
		panic(o)
	}
	return
}

func (s Slice) Negate() {
	for i := len(s) - 1; i > -1; i-- {
		s[i] = -s[i]
	}
}

func (s Slice) Increment() {
	for i := len(s) - 1; i > -1; i-- {
		s[i]++
	}
}

func (s Slice) Decrement() {
	for i := len(s) - 1; i > -1; i-- {
		s[i]--
	}
}

func (s Slice) Add(o interface{}) {
	limit := len(s)
	switch o := o.(type) {
	case complex64:
		for i := len(s) - 1; i > -1; i-- {
			s[i] += o
		}
	case []complex64:
		if limit > len(o) {
			limit = len(o)
		}
		for l := limit - 1; l > -1; l-- {
			s[l] += o[l]
		}
	case Slice:
		if limit > len(o) {
			limit = len(o)
		}
		for l := limit - 1; l > -1; l-- {
			s[l] += o[l]
		}
	default:
		panic(o)
	}
}

func (s Slice) Subtract(o interface{}) {
	limit := len(s)
	switch o := o.(type) {
	case complex64:
		for i := len(s) - 1; i > -1; i-- {
			s[i] -= o
		}
	case []complex64:
		if limit > len(o) {
			limit = len(o)
		}
		for l := limit - 1; l > -1; l-- {
			s[l] -= o[l]
		}
	case Slice:
		if limit > len(o) {
			limit = len(o)
		}
		for l := limit - 1; l > -1; l-- {
			s[l] -= o[l]
		}
	default:
		panic(o)
	}
}

func (s Slice) Multiply(o interface{}) {
	limit := len(s)
	switch o := o.(type) {
	case complex64:
		for i := len(s) - 1; i > -1; i-- {
			s[i] *= o
		}
	case []complex64:
		if limit > len(o) {
			limit = len(o)
		}
		for l := limit - 1; l > -1; l-- {
			s[l] *= o[l]
		}
	case Slice:
		if limit > len(o) {
			limit = len(o)
		}
		for l := limit - 1; l > -1; l-- {
			s[l] *= o[l]
		}
	default:
		panic(o)
	}
}

func (s Slice) Divide(o interface{}) {
	limit := len(s)
	switch o := o.(type) {
	case complex64:
		for i := len(s) - 1; i > -1; i-- {
			s[i] /= o
		}
	case []complex64:
		if limit > len(o) {
			limit = len(o)
		}
		for l := limit - 1; l > -1; l-- {
			s[l] /= o[l]
		}
	case Slice:
		if limit > len(o) {
			limit = len(o)
		}
		for l := limit - 1; l > -1; l-- {
			s[l] /= o[l]
		}
	default:
		panic(o)
	}
}

func (s Slice) IsInf() (r boolean.Slice) {
	if l := len(s); l > 0 {
		r = make(boolean.Slice, l)
		for i := l - 1; i > -1; i-- {
			r[i] = cmplx.IsInf(complex128(s[i]))
		}
	}
	return
}

func (s Slice) IsNaN() (r boolean.Slice) {
	if l := len(s); l > 0 {
		r = make(boolean.Slice, l)
		for i := l - 1; i > -1; i-- {
			r[i] = cmplx.IsNaN(complex128(s[i]))
		}
	}
	return
}