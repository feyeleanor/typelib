package boolean

import(
	"testing"
)

func TestMakeChain(t *testing.T) {
	ConfirmMakeChain := func(n int, r *Chain) {
		if x := MakeChain(n); !r.Equal(x) {
			t.Fatalf("MakeChain(%v) should be %v but is %v", n, r, x)
		}
	}

	ConfirmMakeChain(-1, nil)
	ConfirmMakeChain(0, nil)
	ConfirmMakeChain(1, NewChain(false))
	ConfirmMakeChain(2, NewChain(false, false))
}

func TestNewChain(t *testing.T) {
	ConfirmNewChain := func(l *Chain, r []bool) {
		x := l.start
		for i, v := range r {
			if x.Content() != v {
				t.Fatalf("NewChain(%v...)[%v] should be %v but is %v", r, i, v, x.Content())
			}
			x = x.Next()
		}
	}

	ConfirmNewChain(NewChain(), []bool{})
	ConfirmNewChain(NewChain(true), []bool{ true })
	ConfirmNewChain(NewChain(false), []bool{ false })
	ConfirmNewChain(NewChain(true, false), []bool{ true, false })
}

func TestChainEqual(t *testing.T) {
	ConfirmEqual := func(l, r *Chain, ok bool) {
		if x := l.Equal(r); x != ok {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", l, r, ok, x)
		}
	}

	ConfirmEqual(NewChain(), NewChain(), true)
	ConfirmEqual(NewChain(), NewChain(false), false)
	ConfirmEqual(NewChain(false), NewChain(), false)
	ConfirmEqual(NewChain(false), NewChain(false), true)
	ConfirmEqual(NewChain(false), NewChain(true), false)
	ConfirmEqual(NewChain(true), NewChain(false), false)
	ConfirmEqual(NewChain(true), NewChain(true), true)

	ConfirmEqual(NewChain(false, true), NewChain(true), false)
	ConfirmEqual(NewChain(true, false), NewChain(true), false)
	ConfirmEqual(NewChain(true, false), NewChain(false, true), false)

	ConfirmEqual(NewChain(false, false), NewChain(false, false), true)
	ConfirmEqual(NewChain(false, false), NewChain(false, true), false)
	ConfirmEqual(NewChain(false, false), NewChain(true, true), false)
	ConfirmEqual(NewChain(false, false), NewChain(true, false), false)

	ConfirmEqual(NewChain(false, true), NewChain(false, false), false)
	ConfirmEqual(NewChain(false, true), NewChain(false, true), true)
	ConfirmEqual(NewChain(false, true), NewChain(true, true), false)
	ConfirmEqual(NewChain(false, true), NewChain(true, false), false)

	ConfirmEqual(NewChain(true, true), NewChain(false, false), false)
	ConfirmEqual(NewChain(true, true), NewChain(false, true), false)
	ConfirmEqual(NewChain(true, true), NewChain(true, true), true)
	ConfirmEqual(NewChain(true, true), NewChain(true, false), false)

	ConfirmEqual(NewChain(true, false), NewChain(false, false), false)
	ConfirmEqual(NewChain(true, false), NewChain(false, true), false)
	ConfirmEqual(NewChain(true, false), NewChain(true, true), false)
	ConfirmEqual(NewChain(true, false), NewChain(true, false), true)
}
/*func TestDoublyLinkedListString(t *testing.T) {
	ConfirmString := func(s DoublyLinkedList, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v.String() expected %v but produced %v", s, r, x)
		}
	}

	ConfirmString(NewDoublyLinkedList(), "(list boolean ())")
	ConfirmString(NewDoublyLinkedList(false), "(list boolean (false))")
	ConfirmString(NewDoublyLinkedList(false, true), "(list boolean (false true))")
}*/