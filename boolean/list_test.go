package boolean

import(
	"testing"
)

func TestMakeList(t *testing.T) {
	ConfirmMakeList := func(n int) {
		x := MakeList(n)
		i := 0
		for ; i < n; i++ {
			if x.value {
				t.Fatalf("MakeList(%v)[%v] should be false but is true", n, i)
			}
			x = x.next
		}
		if x != nil {
			t.Fatalf("MakeList(%v) should contain %v elements but must contain more", n, i)
		}
	}

	ConfirmMakeList(0)
	ConfirmMakeList(1)
	ConfirmMakeList(-1)
	ConfirmMakeList(10)
}

func TestNewList(t *testing.T) {
	ConfirmNewList := func(l *List, r []bool) {
		x := l.Clone()
		for i, v := range r {
			if x.value != v {
				t.Fatalf("NewList(%v...)[%v] should be %v but is %v", r, i, v, x.value)
			}
			x = x.next
		}
	}

	ConfirmNewList(NewList(), []bool{})
	ConfirmNewList(NewList(true), []bool{ true })
	ConfirmNewList(NewList(false), []bool{ false })
	ConfirmNewList(NewList(true, false), []bool{ true, false })
}

func TestListAt(t *testing.T) {
	ConfirmAt := func(l *List, n int, v bool) {
		if x := l.At(n); x.value != v {
			t.Fatalf("%v.At(%v) should be %v but is %v", l, n, v, x)
		}
	}

	ConfirmAt(NewList(true, false, true), 0, true)
	ConfirmAt(NewList(true, false, true), 1, false)
	ConfirmAt(NewList(true, false, true), 2, true)
}

func TestListSet(t *testing.T) {
	ConfirmSet := func(l *List, n int, v bool, r *List) {
		x := l.Clone()
		if x.Set(n, v); !x.Equal(r) {
			t.Fatalf("%v.Set(%v) should be %v but is %v", l, n, r, x)
		}
	}

	ConfirmSet(NewList(true, false, true), 0, false, NewList(false, false, true))
	ConfirmSet(NewList(true, false, true), 1, true, NewList(true, true, true))
	ConfirmSet(NewList(true, false, true), 2, false, NewList(true, false, false))
}

func TestListEnd(t *testing.T) {
	ConfirmEnd := func(l, r *List) {
		if x := l.End(); x != r {
			t.Fatalf("%v.End() should be %v but is %v", l, r, x)
		}
	}

	end := &List{ value: true }
	ConfirmEnd(end, end)
	ConfirmEnd(NewList().Append(end), end)
	ConfirmEnd(NewList(false).Append(end), end)
}

func TestListString(t *testing.T) {
	ConfirmString := func(s *List, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v.String() expected %v but produced %v", s, r, x)
		}
	}

	ConfirmString(NewList(), "(list boolean ())")
	ConfirmString(NewList(false), "(list boolean (false))")
	ConfirmString(NewList(false, true), "(list boolean (false true))")
}

func TestListEqual(t *testing.T) {
	ConfirmEqual := func(l, r *List, ok bool) {
		if x := l.Equal(r); x != ok {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", l, r, ok, x)
		}
	}

	ConfirmEqual(NewList(), NewList(), true)
	ConfirmEqual(NewList(), NewList(false), false)
	ConfirmEqual(NewList(false), NewList(), false)
	ConfirmEqual(NewList(false), NewList(false), true)
	ConfirmEqual(NewList(false), NewList(true), false)
	ConfirmEqual(NewList(true), NewList(false), false)
	ConfirmEqual(NewList(true), NewList(true), true)

	ConfirmEqual(NewList(false, true), NewList(true), false)
	ConfirmEqual(NewList(true, false), NewList(true), false)
	ConfirmEqual(NewList(true, false), NewList(false, true), false)

	ConfirmEqual(NewList(false, false), NewList(false, false), true)
	ConfirmEqual(NewList(false, false), NewList(false, true), false)
	ConfirmEqual(NewList(false, false), NewList(true, true), false)
	ConfirmEqual(NewList(false, false), NewList(true, false), false)

	ConfirmEqual(NewList(false, true), NewList(false, false), false)
	ConfirmEqual(NewList(false, true), NewList(false, true), true)
	ConfirmEqual(NewList(false, true), NewList(true, true), false)
	ConfirmEqual(NewList(false, true), NewList(true, false), false)

	ConfirmEqual(NewList(true, true), NewList(false, false), false)
	ConfirmEqual(NewList(true, true), NewList(false, true), false)
	ConfirmEqual(NewList(true, true), NewList(true, true), true)
	ConfirmEqual(NewList(true, true), NewList(true, false), false)

	ConfirmEqual(NewList(true, false), NewList(false, false), false)
	ConfirmEqual(NewList(true, false), NewList(false, true), false)
	ConfirmEqual(NewList(true, false), NewList(true, true), false)
	ConfirmEqual(NewList(true, false), NewList(true, false), true)
}

func TestListLen(t *testing.T) {
	ConfirmLen := func(s *List, r int) {
		if x := s.Len(); x != r {
			t.Fatalf("%v.Len() should be %v but is %v", s, r, x)
		}
	}

	ConfirmLen(NewList(), 0)
	ConfirmLen(NewList(false), 1)
	ConfirmLen(NewList(true), 1)
	ConfirmLen(NewList(false, true), 2)
	ConfirmLen(NewList(false, true, true), 3)
	ConfirmLen(NewList(false, true, false), 3)
}

func TestListAppend(t *testing.T) {
	ConfirmAppend := func(s *List, v interface{}, r *List) {
		x := s.Clone()
		if x = x.Append(v); !r.Equal(x) {
			t.Fatalf("%v.Append(%v) should be %v but is %v", s, v, r, x)
		}
	}

	ConfirmAppend(NewList(), nil, NewList())
	ConfirmAppend(NewList(), true, NewList(true))
	ConfirmAppend(NewList(), false, NewList(false))

	ConfirmAppend(NewList(), []bool{}, NewList())
	ConfirmAppend(NewList(), []bool{ true }, NewList(true))
	ConfirmAppend(NewList(), []bool{ false }, NewList(false))
	ConfirmAppend(NewList(), []bool{ true, false }, NewList(true, false))
	ConfirmAppend(NewList(), []bool{ false, true }, NewList(false, true))

	ConfirmAppend(NewList(), NewList(), NewList())
	ConfirmAppend(NewList(), NewList(true), NewList(true))
	ConfirmAppend(NewList(), NewList(false), NewList(false))
	ConfirmAppend(NewList(), NewList(true, false), NewList(true, false))
	ConfirmAppend(NewList(), NewList(false, true), NewList(false, true))

	ConfirmAppend(NewList(true), nil, NewList(true))
	ConfirmAppend(NewList(true), true, NewList(true, true))
	ConfirmAppend(NewList(true, false), true, NewList(true, false, true))

	ConfirmAppend(NewList(true), []bool{}, NewList(true))
	ConfirmAppend(NewList(true), []bool{ true }, NewList(true, true))
	ConfirmAppend(NewList(true, false), []bool{ true }, NewList(true, false, true))
	ConfirmAppend(NewList(true, false), []bool{ false, true }, NewList(true, false, false, true))

	ConfirmAppend(NewList(true), NewList(), NewList(true))
	ConfirmAppend(NewList(true), NewList(true), NewList(true, true))
	ConfirmAppend(NewList(true, false), NewList(true), NewList(true, false, true))
	ConfirmAppend(NewList(true, false), NewList(false, true), NewList(true, false, false, true))
}

func TestListPrepend(t *testing.T) {
	ConfirmPrepend := func(s *List, v interface{}, r *List) {
		x := s.Clone()
		if x = x.Prepend(v); !r.Equal(x) {
			t.Fatalf("%v.Prepend(%v) should be %v but is %v", s, v, r, x)
		}
	}

	ConfirmPrepend(NewList(), nil, NewList())
	ConfirmPrepend(NewList(), true, NewList(true))
	ConfirmPrepend(NewList(), false, NewList(false))

	ConfirmPrepend(NewList(), []bool{}, NewList())
	ConfirmPrepend(NewList(), []bool{ true }, NewList(true))
	ConfirmPrepend(NewList(), []bool{ false }, NewList(false))
	ConfirmPrepend(NewList(), []bool{ true, false }, NewList(true, false))
	ConfirmPrepend(NewList(), []bool{ false, true }, NewList(false, true))

	ConfirmPrepend(NewList(), NewList(), NewList())
	ConfirmPrepend(NewList(), NewList(true), NewList(true))
	ConfirmPrepend(NewList(), NewList(false), NewList(false))
	ConfirmPrepend(NewList(), NewList(true, false), NewList(true, false))
	ConfirmPrepend(NewList(), NewList(false, true), NewList(false, true))

	ConfirmPrepend(NewList(true), nil, NewList(true))
	ConfirmPrepend(NewList(true), true, NewList(true, true))
	ConfirmPrepend(NewList(true), false, NewList(false, true))

	ConfirmPrepend(NewList(true), []bool{}, NewList(true))
	ConfirmPrepend(NewList(true), []bool{ true }, NewList(true, true))
	ConfirmPrepend(NewList(true, false), []bool{ true }, NewList(true, true, false))
	ConfirmPrepend(NewList(true, false), []bool{ false, true }, NewList(false, true, true, false))


	ConfirmPrepend(NewList(true), NewList(), NewList(true))
	ConfirmPrepend(NewList(true), NewList(true), NewList(true, true))
	ConfirmPrepend(NewList(true, false), NewList(true), NewList(true, true, false))
	ConfirmPrepend(NewList(true, false), NewList(false, true), NewList(false, true, true, false))
}

func TestListClone(t *testing.T) {
	ConfirmClone := func(s, r *List) {
		if x := s.Clone(); !r.Equal(x) {
			t.Fatalf("%v.Clone() should be %v but is %v", s, r, x)
		}
	}

	ConfirmClone(NewList(), NewList())
	ConfirmClone(NewList(true), NewList(true))
	ConfirmClone(NewList(true, true), NewList(true, true))
	ConfirmClone(NewList(true, false, true), NewList(true, false, true))
}

func TestListEach(t *testing.T) {
	slice := []bool{ true, false, false, true, false, true, true }
	list := NewList(slice...)
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

	ConfirmEachPredicated := func(s *List, f interface{}, r int) {
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

func TestListCollect(t *testing.T) {
	ConfirmCollect := func(s *List, f interface{}, r *List) {
		if x := s.Collect(f); !x.Equal(r) {
			t.Fatalf("%v.Collect(%v) should be %v but is %v", s, f, r, x)
		}
	}

	list := NewList(true, false, false, true, false, true, true)
	ConfirmCollect(list, func(v bool) bool { return v }, list)
	ConfirmCollect(list, func(i int, v bool) bool { return v }, list)
	ConfirmCollect(list, func(key interface{}, v bool) bool { return v }, list)
}

func TestListDelete(t *testing.T) {
	ConfirmDelete := func(s *List, f interface{}, r *List) {
		if x := s.Delete(f); !x.Equal(r) {
			t.Fatalf("%v.Delete(%v) should be %v but is %v", s, f, r, x)
		}
	}

	ConfirmDelete(NewList(true, false, true, false), true, NewList(false, false))
	ConfirmDelete(NewList(true, false, true, false), false, NewList(true, true))

	ConfirmDelete(NewList(true, true, true, true), []bool{false, true, true, true}, NewList(true))
	ConfirmDelete(NewList(true, true, true, true), []bool{false, false, true, true}, NewList(true, true))
	ConfirmDelete(NewList(true, true, true, true), []bool{false, false, false, true}, NewList(true, true, true))
	ConfirmDelete(NewList(true, true, true, true), []bool{false, false, false, false}, NewList(true, true, true, true))

	ConfirmDelete(NewList(true, true, true, true), NewList(false, true, true, true), NewList(true))
	ConfirmDelete(NewList(true, true, true, true), NewList(false, false, true, true), NewList(true, true))
	ConfirmDelete(NewList(true, true, true, true), NewList(false, false, false, true), NewList(true, true, true))
	ConfirmDelete(NewList(true, true, true, true), NewList(false, false, false, false), NewList(true, true, true, true))

	ConfirmDelete(NewList(false, true, true, true), func(v bool) bool { return !v }, NewList(true, true, true))
	ConfirmDelete(NewList(false, false, true, true), func(v bool) bool { return !v }, NewList(true, true))
	ConfirmDelete(NewList(false, false, false, true), func(v bool) bool { return !v }, NewList(true))
	ConfirmDelete(NewList(false, false, false, false), func(v bool) bool { return !v }, NewList())

	ConfirmDelete(NewList(false, true, true, true), func(v bool) bool { return v }, NewList(false))
	ConfirmDelete(NewList(false, false, true, true), func(v bool) bool { return v }, NewList(false, false))
	ConfirmDelete(NewList(false, false, false, true), func(v bool) bool { return v }, NewList(false, false, false))
	ConfirmDelete(NewList(false, false, false, false), func(v bool) bool { return v }, NewList(false, false, false, false))

	ConfirmDelete(NewList(false, true, true, true), func(i int, v bool) bool { return !v }, NewList(true, true, true))
	ConfirmDelete(NewList(false, false, true, true), func(i int, v bool) bool { return !v }, NewList(true, true))
	ConfirmDelete(NewList(false, false, false, true), func(i int, v bool) bool { return !v }, NewList(true))
	ConfirmDelete(NewList(false, false, false, false), func(i int, v bool) bool { return !v }, NewList())

	ConfirmDelete(NewList(false, true, true, true), func(i int, v bool) bool { return v }, NewList(false))
	ConfirmDelete(NewList(false, false, true, true), func(i int, v bool) bool { return v }, NewList(false, false))
	ConfirmDelete(NewList(false, false, false, true), func(i int, v bool) bool { return v }, NewList(false, false, false))
	ConfirmDelete(NewList(false, false, false, false), func(i int, v bool) bool { return v }, NewList(false, false, false, false))

	ConfirmDelete(NewList(false, true, true, true), func(i int, v bool) bool { return i > 2 }, NewList(false, true, true))
	ConfirmDelete(NewList(false, true, true, true), func(i int, v bool) bool { return i > 1 }, NewList(false, true))
	ConfirmDelete(NewList(false, true, true, true), func(i int, v bool) bool { return i > 0 }, NewList(false))
	ConfirmDelete(NewList(false, true, true, true), func(i int, v bool) bool { return true }, NewList())

	ConfirmDelete(NewList(false, true, true, true), func(i interface{}, v bool) bool { return !v }, NewList(true, true, true))
	ConfirmDelete(NewList(false, false, true, true), func(i interface{}, v bool) bool { return !v }, NewList(true, true))
	ConfirmDelete(NewList(false, false, false, true), func(i interface{}, v bool) bool { return !v }, NewList(true))
	ConfirmDelete(NewList(false, false, false, false), func(i interface{}, v bool) bool { return !v }, NewList())

	ConfirmDelete(NewList(false, true, true, true), func(i interface{}, v bool) bool { return v }, NewList(false))
	ConfirmDelete(NewList(false, false, true, true), func(i interface{}, v bool) bool { return v }, NewList(false, false))
	ConfirmDelete(NewList(false, false, false, true), func(i interface{}, v bool) bool { return v }, NewList(false, false, false))
	ConfirmDelete(NewList(false, false, false, false), func(i interface{}, v bool) bool { return v }, NewList(false, false, false, false))

	ConfirmDelete(NewList(false, true, true, true), func(i interface{}, v bool) bool { return i.(int) > 2 }, NewList(false, true, true))
	ConfirmDelete(NewList(false, true, true, true), func(i interface{}, v bool) bool { return i.(int) > 1 }, NewList(false, true))
	ConfirmDelete(NewList(false, true, true, true), func(i interface{}, v bool) bool { return i.(int) > 0 }, NewList(false))
	ConfirmDelete(NewList(false, true, true, true), func(i interface{}, v bool) bool { return true }, NewList())
}

func TestListReduce(t *testing.T) {
	ConfirmReduce := func(s *List, f func(bool, bool) bool, seed interface{}, r bool) {
		var l	*List
		if seed, ok := seed.(bool); ok {
			l = s.Prepend(seed)
		} else {
			l = s
		}
		if x := l.Reduce(f); x != r {
			t.Fatalf("%v.Reduce(%v) with seed %v should be %v but is %v", s, f, seed, r, x)
		}
	}

	ConfirmReduce(NewList(true, true, true), func(seed bool, v bool) bool { return seed && v }, nil, true)
	ConfirmReduce(NewList(false, true, true), func(seed bool, v bool) bool { return seed && v }, nil, false)
	ConfirmReduce(NewList(true, false, true), func(seed bool, v bool) bool { return seed && v }, nil, false)
	ConfirmReduce(NewList(true, true, false), func(seed bool, v bool) bool { return seed && v }, nil, false)

	ConfirmReduce(NewList(true, true, true), func(seed bool, v bool) bool { return seed || v }, nil, true)
	ConfirmReduce(NewList(false, true, true), func(seed bool, v bool) bool { return seed || v }, nil, true)
	ConfirmReduce(NewList(true, false, true), func(seed bool, v bool) bool { return seed || v }, nil, true)
	ConfirmReduce(NewList(true, true, false), func(seed bool, v bool) bool { return seed || v }, nil, true)

	ConfirmReduce(NewList(true, true, true), func(seed bool, v bool) bool { return seed && v }, false, false)
	ConfirmReduce(NewList(false, true, true), func(seed bool, v bool) bool { return seed && v }, false, false)
	ConfirmReduce(NewList(true, false, true), func(seed bool, v bool) bool { return seed && v }, false, false)
	ConfirmReduce(NewList(true, true, false), func(seed bool, v bool) bool { return seed && v }, false, false)

	ConfirmReduce(NewList(true, true, true), func(seed bool, v bool) bool { return seed || v }, false, true)
	ConfirmReduce(NewList(false, true, true), func(seed bool, v bool) bool { return seed || v }, false, true)
	ConfirmReduce(NewList(true, false, true), func(seed bool, v bool) bool { return seed || v }, false, true)
	ConfirmReduce(NewList(true, true, false), func(seed bool, v bool) bool { return seed || v }, false, true)
}

func TestListReverse(t *testing.T) {
	ConfirmReverse := func(s, r *List) {
		if x := s.Reverse(); !r.Equal(x) {
			t.Fatalf("%v.Reverse() should be %v but is %v", s, r, x)
		}
	}

	ConfirmReverse(NewList(), NewList())
	ConfirmReverse(NewList(true), NewList(true))
	ConfirmReverse(NewList(true, false), NewList(false, true))
	ConfirmReverse(NewList(true, false, false), NewList(false, false, true))
}