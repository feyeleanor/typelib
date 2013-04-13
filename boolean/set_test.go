package boolean

import (
	"testing"
)

func TestSetString(t *testing.T) {
	ConfirmString := func(s Set, r string) {
		if v := s.String(); r != v {
			t.Errorf("%v.String() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmString(NewSet(), "(set boolean ())")
	ConfirmString(NewSet(false, true), "(set boolean (false true))")
	ConfirmString(NewSet(true, false), "(set boolean (false true))")
}

func TestSetMember(t *testing.T) {
	ConfirmMember := func(s Set, x, r bool) {
		if v := s.Member(x); r != v {
			t.Errorf("%v.Member(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmMember(NewSet(), false, false)
	ConfirmMember(NewSet(), true, false)
	ConfirmMember(NewSet(false), false, true)
	ConfirmMember(NewSet(false), true, false)
	ConfirmMember(NewSet(true), false, false)
	ConfirmMember(NewSet(true), true, true)
	ConfirmMember(NewSet(false, true), false, true)
	ConfirmMember(NewSet(false, true), true, true)
}

func TestSetEqual(t *testing.T) {
	ConfirmEqual := func(s, x Set, r bool) {
		if v := s.Equal(x); v != r {
			t.Errorf("%v.Equal(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmEqual(NewSet(), NewSet(), true)
	ConfirmEqual(NewSet(false), NewSet(), false)
	ConfirmEqual(NewSet(), NewSet(false), false)
	ConfirmEqual(NewSet(false), NewSet(false), true)
	ConfirmEqual(NewSet(false, false), NewSet(false), true)
	ConfirmEqual(NewSet(false), NewSet(false, false), true)
	ConfirmEqual(NewSet(false, true), NewSet(false, false), false)
	ConfirmEqual(NewSet(false, true), NewSet(false, true), true)
	ConfirmEqual(NewSet(false, true), NewSet(true, false), true)
	ConfirmEqual(NewSet(false, true), NewSet(true, true), false)
}