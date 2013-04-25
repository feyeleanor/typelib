package boolean

import(
	"testing"
	"typelib"
)

func TestMakeVector(t *testing.T) {
	ConfirmMakeVector := func(n int, r *Vector) {
		if x := MakeVector(n); !r.Equal(x) {
			t.Fatalf("MakeVector(%v) should be %v but is %v", n, r, x)
		}
	}

	ConfirmMakeVector(-1, nil)
	ConfirmMakeVector(0, nil)
	ConfirmMakeVector(1, NewVector(false))
	ConfirmMakeVector(2, NewVector(false, false))
}

func TestNewVector(t *testing.T) {
	ConfirmNewVector := func(l *Vector, r []bool) {
		x := l.Clone().(*Vector)
		for i, v := range r {
			if x.Content() != v {
				t.Fatalf("NewVector(%v...)[%v] should be %v but is %v", r, i, v, x.Content())
			}
			x = x.Tail().(*Vector)
		}
	}

	ConfirmNewVector(NewVector(), []bool{})
	ConfirmNewVector(NewVector(true), []bool{ true })
	ConfirmNewVector(NewVector(false), []bool{ false })
	ConfirmNewVector(NewVector(true, false), []bool{ true, false })
}

func TestVectorIsNil(t *testing.T) {
	ConfirmIsNil := func(l *Vector, r bool) {
		if l.IsNil() != r {
			t.Fatalf("%v.IsNil() should be %v but is %v", l, r, l.IsNil())
		}
	}

	ConfirmIsNil(nil, true)
	ConfirmIsNil(NewVector(), true)
	ConfirmIsNil(NewVector(true), false)
	ConfirmIsNil(NewVector(false), false)
}

func TestVectorContent(t *testing.T) {
	ConfirmContent := func(l *Vector, r interface{}) {
		if x := l.Content(); x != r {
			t.Fatalf("%v.Content() should be %v but is %v", l, r, x)
		}
	}

	ConfirmContent(NewVector(), nil)
	ConfirmContent(NewVector(true), true)
	ConfirmContent(NewVector(false), false)
	ConfirmContent(NewVector(true, false), true)
	ConfirmContent(NewVector(false, true), false)
}

func TestVectorTail(t *testing.T) {
	ConfirmTail := func(l, r *Vector) {
		if x := l.Tail(); !r.Equal(x) {
			t.Fatalf("%v.Tail() should be %v but is %v", l, r, x)
		}
	}

	ConfirmTail(NewVector(), nil)
	ConfirmTail(NewVector(true), nil)
	ConfirmTail(NewVector(true, false), NewVector(false))
	ConfirmTail(NewVector(true, false, true), NewVector(false, true))
}

func TestVectorAt(t *testing.T) {
	ConfirmAt := func(l *Vector, n int, v bool) {
		if x := l.At(n); x != v {
			t.Fatalf("%v.At(%v) should be %v but is %v", l, n, v, x)
		}
	}

	ConfirmAt(NewVector(true, false, true), 0, true)
	ConfirmAt(NewVector(true, false, true), 1, false)
	ConfirmAt(NewVector(true, false, true), 2, true)
}

func TestVectorStore(t *testing.T) {
	ConfirmSet := func(l *Vector, n int, v interface{}, r *Vector) {
		x := l.Clone().(*Vector)
		if x.Store(v, n); !x.Equal(r) {
			t.Fatalf("%v.Set(%v, %v) should be %v but is %v", l, n, v, r, x)
		}
	}

	ConfirmSet(NewVector(true, false, true), 0, false, NewVector(false, false, true))
	ConfirmSet(NewVector(true, false, true), 1, true, NewVector(true, true, true))
	ConfirmSet(NewVector(true, false, true), 2, false, NewVector(true, false, false))

	ConfirmSet(NewVector(true, false, true), 0, []bool{ false }, NewVector(false))
	ConfirmSet(NewVector(true, false, true), 1, []bool{ true }, NewVector(true, true))
	ConfirmSet(NewVector(true, false, true), 2, []bool{ false }, NewVector(true, false, false))

	ConfirmSet(NewVector(true, false, true), 0, []bool{ false, false }, NewVector(false, false))
	ConfirmSet(NewVector(true, false, true), 1, []bool{ true, true }, NewVector(true, true, true))
	ConfirmSet(NewVector(true, false, true), 2, []bool{ false, true }, NewVector(true, false, false, true))
}

func TestVectorEnd(t *testing.T) {
	ConfirmEnd := func(l, r *Vector) {
		if x := l.End(); x != r {
			t.Fatalf("%v.End() should be %v but is %v", l, r, x)
		}
	}

	end := &Vector{ value: true }
	ConfirmEnd(end, end)
	ConfirmEnd(NewVector().Append(end).(*Vector), end)
	ConfirmEnd(NewVector(false).Append(end).(*Vector), end)
}

func TestVectorString(t *testing.T) {
	ConfirmString := func(s *Vector, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v.String() expected %v but produced %v", s, r, x)
		}
	}

	ConfirmString(NewVector(), "(list boolean ())")
	ConfirmString(NewVector(false), "(list boolean (false))")
	ConfirmString(NewVector(false, true), "(list boolean (false true))")
}

func TestVectorEqual(t *testing.T) {
	ConfirmEqual := func(l, r *Vector, ok bool) {
		if x := l.Equal(r); x != ok {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", l, r, ok, x)
		}
	}

	ConfirmEqual(NewVector(), NewVector(), true)
	ConfirmEqual(NewVector(), NewVector(false), false)
	ConfirmEqual(NewVector(false), NewVector(), false)
	ConfirmEqual(NewVector(false), NewVector(false), true)
	ConfirmEqual(NewVector(false), NewVector(true), false)
	ConfirmEqual(NewVector(true), NewVector(false), false)
	ConfirmEqual(NewVector(true), NewVector(true), true)

	ConfirmEqual(NewVector(false, true), NewVector(true), false)
	ConfirmEqual(NewVector(true, false), NewVector(true), false)
	ConfirmEqual(NewVector(true, false), NewVector(false, true), false)

	ConfirmEqual(NewVector(false, false), NewVector(false, false), true)
	ConfirmEqual(NewVector(false, false), NewVector(false, true), false)
	ConfirmEqual(NewVector(false, false), NewVector(true, true), false)
	ConfirmEqual(NewVector(false, false), NewVector(true, false), false)

	ConfirmEqual(NewVector(false, true), NewVector(false, false), false)
	ConfirmEqual(NewVector(false, true), NewVector(false, true), true)
	ConfirmEqual(NewVector(false, true), NewVector(true, true), false)
	ConfirmEqual(NewVector(false, true), NewVector(true, false), false)

	ConfirmEqual(NewVector(true, true), NewVector(false, false), false)
	ConfirmEqual(NewVector(true, true), NewVector(false, true), false)
	ConfirmEqual(NewVector(true, true), NewVector(true, true), true)
	ConfirmEqual(NewVector(true, true), NewVector(true, false), false)

	ConfirmEqual(NewVector(true, false), NewVector(false, false), false)
	ConfirmEqual(NewVector(true, false), NewVector(false, true), false)
	ConfirmEqual(NewVector(true, false), NewVector(true, true), false)
	ConfirmEqual(NewVector(true, false), NewVector(true, false), true)
}

func TestVectorLen(t *testing.T) {
	ConfirmLen := func(s *Vector, r int) {
		if x := s.Len(); x != r {
			t.Fatalf("%v.Len() should be %v but is %v", s, r, x)
		}
	}

	ConfirmLen(NewVector(), 0)
	ConfirmLen(NewVector(false), 1)
	ConfirmLen(NewVector(true), 1)
	ConfirmLen(NewVector(false, true), 2)
	ConfirmLen(NewVector(false, true, true), 3)
	ConfirmLen(NewVector(false, true, false), 3)
}

func TestVectorAppend(t *testing.T) {
	ConfirmAppend := func(s *Vector, v interface{}, r *Vector) {
		x := s.Clone().(*Vector)
		if x = x.Append(v).(*Vector); !r.Equal(x) {
			t.Fatalf("%v.Append(%v) should be %v but is %v", s, v, r, x)
		}
	}

	ConfirmAppend(NewVector(), nil, NewVector())
	ConfirmAppend(NewVector(), true, NewVector(true))
	ConfirmAppend(NewVector(), false, NewVector(false))

	ConfirmAppend(NewVector(), []bool{}, NewVector())
	ConfirmAppend(NewVector(), []bool{ true }, NewVector(true))
	ConfirmAppend(NewVector(), []bool{ false }, NewVector(false))
	ConfirmAppend(NewVector(), []bool{ true, false }, NewVector(true, false))
	ConfirmAppend(NewVector(), []bool{ false, true }, NewVector(false, true))

	ConfirmAppend(NewVector(), NewVector(), NewVector())
	ConfirmAppend(NewVector(), NewVector(true), NewVector(true))
	ConfirmAppend(NewVector(), NewVector(false), NewVector(false))
	ConfirmAppend(NewVector(), NewVector(true, false), NewVector(true, false))
	ConfirmAppend(NewVector(), NewVector(false, true), NewVector(false, true))

	ConfirmAppend(NewVector(true), nil, NewVector(true))
	ConfirmAppend(NewVector(true), true, NewVector(true, true))
	ConfirmAppend(NewVector(true, false), true, NewVector(true, false, true))

	ConfirmAppend(NewVector(true), []bool{}, NewVector(true))
	ConfirmAppend(NewVector(true), []bool{ true }, NewVector(true, true))
	ConfirmAppend(NewVector(true, false), []bool{ true }, NewVector(true, false, true))
	ConfirmAppend(NewVector(true, false), []bool{ false, true }, NewVector(true, false, false, true))

	ConfirmAppend(NewVector(true), NewVector(), NewVector(true))
	ConfirmAppend(NewVector(true), NewVector(true), NewVector(true, true))
	ConfirmAppend(NewVector(true, false), NewVector(true), NewVector(true, false, true))
	ConfirmAppend(NewVector(true, false), NewVector(false, true), NewVector(true, false, false, true))
}

func TestVectorPrepend(t *testing.T) {
	ConfirmPrepend := func(s *Vector, v interface{}, r *Vector) {
		x := s.Clone().(*Vector)
		if x = x.Prepend(v).(*Vector); !r.Equal(x) {
			t.Fatalf("%v.Prepend(%v) should be %v but is %v", s, v, r, x)
		}
	}

	ConfirmPrepend(NewVector(), nil, NewVector())
	ConfirmPrepend(NewVector(), true, NewVector(true))
	ConfirmPrepend(NewVector(), false, NewVector(false))

	ConfirmPrepend(NewVector(), []bool{}, NewVector())
	ConfirmPrepend(NewVector(), []bool{ true }, NewVector(true))
	ConfirmPrepend(NewVector(), []bool{ false }, NewVector(false))
	ConfirmPrepend(NewVector(), []bool{ true, false }, NewVector(true, false))
	ConfirmPrepend(NewVector(), []bool{ false, true }, NewVector(false, true))

	ConfirmPrepend(NewVector(), NewVector(), NewVector())
	ConfirmPrepend(NewVector(), NewVector(true), NewVector(true))
	ConfirmPrepend(NewVector(), NewVector(false), NewVector(false))
	ConfirmPrepend(NewVector(), NewVector(true, false), NewVector(true, false))
	ConfirmPrepend(NewVector(), NewVector(false, true), NewVector(false, true))

	ConfirmPrepend(NewVector(true), nil, NewVector(true))
	ConfirmPrepend(NewVector(true), true, NewVector(true, true))
	ConfirmPrepend(NewVector(true), false, NewVector(false, true))

	ConfirmPrepend(NewVector(true), []bool{}, NewVector(true))
	ConfirmPrepend(NewVector(true), []bool{ true }, NewVector(true, true))
	ConfirmPrepend(NewVector(true, false), []bool{ true }, NewVector(true, true, false))
	ConfirmPrepend(NewVector(true, false), []bool{ false, true }, NewVector(false, true, true, false))


	ConfirmPrepend(NewVector(true), NewVector(), NewVector(true))
	ConfirmPrepend(NewVector(true), NewVector(true), NewVector(true, true))
	ConfirmPrepend(NewVector(true, false), NewVector(true), NewVector(true, true, false))
	ConfirmPrepend(NewVector(true, false), NewVector(false, true), NewVector(false, true, true, false))
}

func TestVectorClone(t *testing.T) {
	ConfirmClone := func(s, r *Vector) {
		if x := s.Clone().(*Vector); !r.Equal(x) {
			t.Fatalf("%v.Clone() should be %v but is %v", s, r, x)
		}
	}

	ConfirmClone(NewVector(), NewVector())
	ConfirmClone(NewVector(true), NewVector(true))
	ConfirmClone(NewVector(true, true), NewVector(true, true))
	ConfirmClone(NewVector(true, false, true), NewVector(true, false, true))
}

func TestVectorEach(t *testing.T) {
	slice := []bool{ true, false, false, true, false, true, true }
	list := NewVector(slice...)
	count := 0
	list.Each(func(v bool) {
		if v != slice[count] {
			t.Fatalf("%v.Each() element %v should be %v but is %v", list, count, slice[count], v)
		}
		count++
	})

	count = 0
	list.Each(func(v interface{}) {
		if v != slice[count] {
			t.Fatalf("%v.Each() element %v should be %v but is %v", list, count, slice[count], v)
		}
		count++
	})

	list.Each(func(i int, v bool) {
		if v != slice[i] {
			t.Fatalf("%v.Each() element %v should be %v but is %v", list, i, slice[i], v)
		}
	})

	list.Each(func(i int, v interface{}) {
		if v != slice[i] {
			t.Fatalf("%v.Each() element %v should be %v but is %v", list, i, slice[i], v)
		}
	})

	list.Each(func(key interface{}, v bool) {
		i := key.(int)
		if v != slice[i] {
			t.Fatalf("%v.Each() element %v should be %v but is %v", list, i, slice[i], v)
		}
	})

	list.Each(func(key, v interface{}) {
		i := key.(int)
		if v != slice[i] {
			t.Fatalf("%v.Each() element %v should be %v but is %v", list, i, slice[i], v)
		}
	})

	ConfirmEachPredicated := func(s *Vector, f interface{}, r int) {
		count = 0
		if s.Each(f); count != r {
			t.Fatalf("%v.Each(%v) should execute %v times but actually executed %v times", s, f, r, count)
		}
	}

	ConfirmEachPredicated(list, func(v bool) bool {
		count++
		return count < 5
	}, 5)

	ConfirmEachPredicated(list, func(v interface{}) bool {
		count++
		return count < 5
	}, 5)

	ConfirmEachPredicated(list, func(i int, v bool) bool {
		count = i
		return count < 5
	}, 5)

	ConfirmEachPredicated(list, func(i int, v interface{}) bool {
		count = i
		return count < 5
	}, 5)

	ConfirmEachPredicated(list, func(key interface{}, v bool) bool {
		count = key.(int)
		return count < 5
	}, 5)

	ConfirmEachPredicated(list, func(key, v interface{}) bool {
		count = key.(int)
		return count < 5
	}, 5)
}

func TestVectorCollect(t *testing.T) {
	ConfirmCollect := func(s *Vector, f interface{}, r *Vector) {
		if x := s.Collect(f); !x.Equal(r) {
			t.Fatalf("%v.Collect(%v) should be %v but is %v", s, f, r, x)
		}
	}

	list := NewVector(true, false, false, true, false, true, true)
	ConfirmCollect(list, func(v bool) bool { return v }, list)
	ConfirmCollect(list, func(i int, v bool) bool { return v }, list)
	ConfirmCollect(list, func(key interface{}, v bool) bool { return v }, list)
}

func TestVectorDelete(t *testing.T) {
	ConfirmDelete := func(s *Vector, f interface{}, r *Vector) {
		if x := s.Delete(f); !x.Equal(r) {
			t.Fatalf("%v.Delete(%v) should be %v but is %v", s, f, r, x)
		}
	}

	ConfirmDelete(NewVector(true, false, true, false), true, NewVector(false, false))
	ConfirmDelete(NewVector(true, false, true, false), false, NewVector(true, true))

	ConfirmDelete(NewVector(true, true, true, true), []bool{false, true, true, true}, NewVector(true))
	ConfirmDelete(NewVector(true, true, true, true), []bool{false, false, true, true}, NewVector(true, true))
	ConfirmDelete(NewVector(true, true, true, true), []bool{false, false, false, true}, NewVector(true, true, true))
	ConfirmDelete(NewVector(true, true, true, true), []bool{false, false, false, false}, NewVector(true, true, true, true))

	ConfirmDelete(NewVector(true, true, true, true), NewVector(false, true, true, true), NewVector(true))
	ConfirmDelete(NewVector(true, true, true, true), NewVector(false, false, true, true), NewVector(true, true))
	ConfirmDelete(NewVector(true, true, true, true), NewVector(false, false, false, true), NewVector(true, true, true))
	ConfirmDelete(NewVector(true, true, true, true), NewVector(false, false, false, false), NewVector(true, true, true, true))

	ConfirmDelete(NewVector(false, true, true, true), func(v bool) bool { return !v }, NewVector(true, true, true))
	ConfirmDelete(NewVector(false, false, true, true), func(v bool) bool { return !v }, NewVector(true, true))
	ConfirmDelete(NewVector(false, false, false, true), func(v bool) bool { return !v }, NewVector(true))
	ConfirmDelete(NewVector(false, false, false, false), func(v bool) bool { return !v }, NewVector())

	ConfirmDelete(NewVector(false, true, true, true), func(v bool) bool { return v }, NewVector(false))
	ConfirmDelete(NewVector(false, false, true, true), func(v bool) bool { return v }, NewVector(false, false))
	ConfirmDelete(NewVector(false, false, false, true), func(v bool) bool { return v }, NewVector(false, false, false))
	ConfirmDelete(NewVector(false, false, false, false), func(v bool) bool { return v }, NewVector(false, false, false, false))

	ConfirmDelete(NewVector(false, true, true, true), func(i int, v bool) bool { return !v }, NewVector(true, true, true))
	ConfirmDelete(NewVector(false, false, true, true), func(i int, v bool) bool { return !v }, NewVector(true, true))
	ConfirmDelete(NewVector(false, false, false, true), func(i int, v bool) bool { return !v }, NewVector(true))
	ConfirmDelete(NewVector(false, false, false, false), func(i int, v bool) bool { return !v }, NewVector())

	ConfirmDelete(NewVector(false, true, true, true), func(i int, v bool) bool { return v }, NewVector(false))
	ConfirmDelete(NewVector(false, false, true, true), func(i int, v bool) bool { return v }, NewVector(false, false))
	ConfirmDelete(NewVector(false, false, false, true), func(i int, v bool) bool { return v }, NewVector(false, false, false))
	ConfirmDelete(NewVector(false, false, false, false), func(i int, v bool) bool { return v }, NewVector(false, false, false, false))

	ConfirmDelete(NewVector(false, true, true, true), func(i int, v bool) bool { return i > 2 }, NewVector(false, true, true))
	ConfirmDelete(NewVector(false, true, true, true), func(i int, v bool) bool { return i > 1 }, NewVector(false, true))
	ConfirmDelete(NewVector(false, true, true, true), func(i int, v bool) bool { return i > 0 }, NewVector(false))
	ConfirmDelete(NewVector(false, true, true, true), func(i int, v bool) bool { return true }, NewVector())

	ConfirmDelete(NewVector(false, true, true, true), func(i interface{}, v bool) bool { return !v }, NewVector(true, true, true))
	ConfirmDelete(NewVector(false, false, true, true), func(i interface{}, v bool) bool { return !v }, NewVector(true, true))
	ConfirmDelete(NewVector(false, false, false, true), func(i interface{}, v bool) bool { return !v }, NewVector(true))
	ConfirmDelete(NewVector(false, false, false, false), func(i interface{}, v bool) bool { return !v }, NewVector())

	ConfirmDelete(NewVector(false, true, true, true), func(i interface{}, v bool) bool { return v }, NewVector(false))
	ConfirmDelete(NewVector(false, false, true, true), func(i interface{}, v bool) bool { return v }, NewVector(false, false))
	ConfirmDelete(NewVector(false, false, false, true), func(i interface{}, v bool) bool { return v }, NewVector(false, false, false))
	ConfirmDelete(NewVector(false, false, false, false), func(i interface{}, v bool) bool { return v }, NewVector(false, false, false, false))

	ConfirmDelete(NewVector(false, true, true, true), func(i interface{}, v bool) bool { return i.(int) > 2 }, NewVector(false, true, true))
	ConfirmDelete(NewVector(false, true, true, true), func(i interface{}, v bool) bool { return i.(int) > 1 }, NewVector(false, true))
	ConfirmDelete(NewVector(false, true, true, true), func(i interface{}, v bool) bool { return i.(int) > 0 }, NewVector(false))
	ConfirmDelete(NewVector(false, true, true, true), func(i interface{}, v bool) bool { return true }, NewVector())
}

func TestVectorReduce(t *testing.T) {
	ConfirmReduce := func(s *Vector, f func(bool, bool) bool, seed interface{}, r *Vector) {
		var l	typelib.List
		if seed, ok := seed.(bool); ok {
			l = s.Prepend(seed)
		} else {
			l = s
		}
		if x := l.Reduce(f); !x.Equal(r) {
			t.Fatalf("%v.Reduce(%v) with seed %v should be %v but is %v", s, f, seed, r, x)
		}
	}

	ConfirmReduce(NewVector(true, true, true), func(seed bool, v bool) bool { return seed && v }, nil, NewVector(true))
	ConfirmReduce(NewVector(false, true, true), func(seed bool, v bool) bool { return seed && v }, nil, NewVector(false))
	ConfirmReduce(NewVector(true, false, true), func(seed bool, v bool) bool { return seed && v }, nil, NewVector(false))
	ConfirmReduce(NewVector(true, true, false), func(seed bool, v bool) bool { return seed && v }, nil, NewVector(false))

	ConfirmReduce(NewVector(true, true, true), func(seed bool, v bool) bool { return seed || v }, nil, NewVector(true))
	ConfirmReduce(NewVector(false, true, true), func(seed bool, v bool) bool { return seed || v }, nil, NewVector(true))
	ConfirmReduce(NewVector(true, false, true), func(seed bool, v bool) bool { return seed || v }, nil, NewVector(true))
	ConfirmReduce(NewVector(true, true, false), func(seed bool, v bool) bool { return seed || v }, nil, NewVector(true))

	ConfirmReduce(NewVector(true, true, true), func(seed bool, v bool) bool { return seed && v }, false, NewVector(false))
	ConfirmReduce(NewVector(false, true, true), func(seed bool, v bool) bool { return seed && v }, false, NewVector(false))
	ConfirmReduce(NewVector(true, false, true), func(seed bool, v bool) bool { return seed && v }, false, NewVector(false))
	ConfirmReduce(NewVector(true, true, false), func(seed bool, v bool) bool { return seed && v }, false, NewVector(false))

	ConfirmReduce(NewVector(true, true, true), func(seed bool, v bool) bool { return seed || v }, false, NewVector(true))
	ConfirmReduce(NewVector(false, true, true), func(seed bool, v bool) bool { return seed || v }, false, NewVector(true))
	ConfirmReduce(NewVector(true, false, true), func(seed bool, v bool) bool { return seed || v }, false, NewVector(true))
	ConfirmReduce(NewVector(true, true, false), func(seed bool, v bool) bool { return seed || v }, false, NewVector(true))
}

func TestVectorReverse(t *testing.T) {
	ConfirmReverse := func(s, r *Vector) {
		if x := s.Reverse(); !r.Equal(x) {
			t.Fatalf("%v.Reverse() should be %v but is %v", s, r, x)
		}
	}

	ConfirmReverse(NewVector(), NewVector())
	ConfirmReverse(NewVector(true), NewVector(true))
	ConfirmReverse(NewVector(true, false), NewVector(false, true))
	ConfirmReverse(NewVector(true, false, false), NewVector(false, false, true))
}