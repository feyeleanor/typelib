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
	ConfirmAt := func(l *List, n int, v bool) {
		if x := l.Set(n, v); x.value != v {
			t.Fatalf("%v.Set(%v) should be %v but is %v", l, n, v, x)
		}
	}

	ConfirmAt(NewList(true, false, true), 0, false)
	ConfirmAt(NewList(true, false, true), 1, true)
	ConfirmAt(NewList(true, false, true), 2, false)
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