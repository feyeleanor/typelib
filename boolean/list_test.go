package boolean

import(
	"testing"
)

func TestNewList(t *testing.T) {
	
}

func TestListString(t *testing.T) {
	ConfirmString := func(s List, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v.String() expected %v but produced %v", s, r, x)
		}
	}

	ConfirmString(NewList(), "(list boolean ())")
	ConfirmString(NewList(false), "(list boolean (false))")
	ConfirmString(NewList(false, true), "(list boolean (false true))")
}