package boolean

import(
	"testing"
)

func TestCons(t *testing.T) {
	ConfirmCons := func(l SinglyLinkedList, r []bool) {
		x := l.Clone()
		for i, v := range r {
			if x.value != v {
				t.Fatalf("Cons(%v...)[%v] should be %v but is %v", r, i, v, x.value)
			}
			x.link = x.next
		}
	}

	ConfirmCons(Cons(), []bool{})
	ConfirmCons(Cons(true), []bool{ true })
	ConfirmCons(Cons(false), []bool{ false })
	ConfirmCons(Cons(true, false), []bool{ true, false })
}

func TestSinglyLinkedListIsNil(t *testing.T) {
	ConfirmIsNil := func(l SinglyLinkedList, r bool) {
		if x := l.IsNil(); x != r {
			t.Fatalf("%v.IsNil() should be %v but is %v", l, r, x)
		}
	}

	ConfirmIsNil(SinglyLinkedList{}, true)
	ConfirmIsNil(Cons(), true)
	ConfirmIsNil(SinglyLinkedList{ link: &link{ value: false } }, false)
	ConfirmIsNil(SinglyLinkedList{ link: &link{ value: true } }, false)
	ConfirmIsNil(Cons(true), false)
	ConfirmIsNil(Cons(false), false)
}

func TestSinglyLinkedListEnd(t *testing.T) {
	ConfirmEnd := func(l SinglyLinkedList, r *link) {
		if x := l.End(); x != r {
			t.Fatalf("%v.End() should be %v but is %v", l, r, x)
		}
	}

	end := &link{ value: true }
	ConfirmEnd(SinglyLinkedList{ link: end }, end)
	ConfirmEnd(SinglyLinkedList{ link: &link{ next: end } }, end)
	ConfirmEnd(SinglyLinkedList{ link: &link{ value: true , next: end } }, end)
}

func TestSinglyLinkedListString(t *testing.T) {
	ConfirmString := func(s SinglyLinkedList, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v.String() expected %v but produced %v", s, r, x)
		}
	}

	ConfirmString(Cons(), "(cons-boolean ())")
	ConfirmString(Cons(false), "(cons-boolean (false))")
	ConfirmString(Cons(false, true), "(cons-boolean (false true))")
}

func TestSinglyLinkedListEqual(t *testing.T) {
	ConfirmEqual := func(l, r SinglyLinkedList, ok bool) {
		if x := l.Equal(r); x != ok {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", l, r, ok, x)
		}
	}

	ConfirmEqual(Cons(), Cons(), true)
	ConfirmEqual(Cons(), Cons(false), false)
	ConfirmEqual(Cons(false), Cons(), false)
	ConfirmEqual(Cons(false), Cons(false), true)
	ConfirmEqual(Cons(false), Cons(true), false)
	ConfirmEqual(Cons(true), Cons(false), false)
	ConfirmEqual(Cons(true), Cons(true), true)

	ConfirmEqual(Cons(false, true), Cons(true), false)
	ConfirmEqual(Cons(true, false), Cons(true), false)
	ConfirmEqual(Cons(true, false), Cons(false, true), false)

	ConfirmEqual(Cons(false, false), Cons(false, false), true)
	ConfirmEqual(Cons(false, false), Cons(false, true), false)
	ConfirmEqual(Cons(false, false), Cons(true, true), false)
	ConfirmEqual(Cons(false, false), Cons(true, false), false)

	ConfirmEqual(Cons(false, true), Cons(false, false), false)
	ConfirmEqual(Cons(false, true), Cons(false, true), true)
	ConfirmEqual(Cons(false, true), Cons(true, true), false)
	ConfirmEqual(Cons(false, true), Cons(true, false), false)

	ConfirmEqual(Cons(true, true), Cons(false, false), false)
	ConfirmEqual(Cons(true, true), Cons(false, true), false)
	ConfirmEqual(Cons(true, true), Cons(true, true), true)
	ConfirmEqual(Cons(true, true), Cons(true, false), false)

	ConfirmEqual(Cons(true, false), Cons(false, false), false)
	ConfirmEqual(Cons(true, false), Cons(false, true), false)
	ConfirmEqual(Cons(true, false), Cons(true, true), false)
	ConfirmEqual(Cons(true, false), Cons(true, false), true)
}

func TestSinglyLinkedListLen(t *testing.T) {
	ConfirmLen := func(s SinglyLinkedList, r int) {
		if x := s.Len(); x != r {
			t.Fatalf("%v.Len() should be %v but is %v", s, r, x)
		}
	}

	ConfirmLen(Cons(), 0)
	ConfirmLen(Cons(false), 1)
	ConfirmLen(Cons(true), 1)
	ConfirmLen(Cons(false, true), 2)
	ConfirmLen(Cons(false, true, true), 3)
	ConfirmLen(Cons(false, true, false), 3)
}

func TestSinglyLinkedListAppend(t *testing.T) {
	ConfirmAppend := func(s SinglyLinkedList, v []bool, r SinglyLinkedList) {
		x := s.Clone()
		if x.Append(v...); !r.Equal(x) {
			t.Fatalf("%v.Append(%v) should be %v but is %v", s, v, r, x)
		}
	}

	ConfirmAppend(Cons(), []bool{}, Cons())
	ConfirmAppend(Cons(), []bool{ true }, Cons(true))
	ConfirmAppend(Cons(), []bool{ false }, Cons(false))
	ConfirmAppend(Cons(), []bool{ true, false }, Cons(true, false))
	ConfirmAppend(Cons(), []bool{ false, true }, Cons(false, true))

	ConfirmAppend(Cons(true), []bool{}, Cons(true))
	ConfirmAppend(Cons(true), []bool{ true }, Cons(true, true))
	ConfirmAppend(Cons(true, false), []bool{ true }, Cons(true, false, true))
	ConfirmAppend(Cons(true, false), []bool{ false, true }, Cons(true, false, false, true))
}

func TestSinglyLinkedListPrepend(t *testing.T) {
	ConfirmPrepend := func(s SinglyLinkedList, v []bool, r SinglyLinkedList) {
		x := s.Clone()
		if x.Prepend(v...); !r.Equal(x) {
			t.Fatalf("%v.Prepend(%v) should be %v but is %v", s, v, r, x)
		}
	}

	ConfirmPrepend(Cons(), []bool{}, Cons())
	ConfirmPrepend(Cons(), []bool{ true }, Cons(true))
	ConfirmPrepend(Cons(), []bool{ false }, Cons(false))
	ConfirmPrepend(Cons(), []bool{ true, false }, Cons(true, false))
	ConfirmPrepend(Cons(), []bool{ false, true }, Cons(false, true))

	ConfirmPrepend(Cons(true), []bool{}, Cons(true))
	ConfirmPrepend(Cons(true), []bool{ true }, Cons(true, true))
	ConfirmPrepend(Cons(true, false), []bool{ true }, Cons(true, true, false))
	ConfirmPrepend(Cons(true, false), []bool{ false, true }, Cons(false, true, true, false))
}

func TestSinglyLinkedListClone(t *testing.T) {
	ConfirmClone := func(s, r SinglyLinkedList) {
		if x := s.Clone(); !r.Equal(x) {
			t.Fatalf("%v.Clone() should be %v but is %v", s, r, x)
		}
	}

	ConfirmClone(Cons(), Cons())
	ConfirmClone(Cons(true), Cons(true))
	ConfirmClone(Cons(true, true), Cons(true, true))
	ConfirmClone(Cons(true, false, true), Cons(true, false, true))
}

func TestSinglyLinkedListEach(t *testing.T) {
	slice := []bool{ true, false, false, true, false, true, true }
	list := Cons(slice...)
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