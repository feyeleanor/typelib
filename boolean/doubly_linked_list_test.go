package boolean

import(
	"testing"
)

func TestNewDoublyLinkedList(t *testing.T) {
	
}

func TestDoublyLinkedListString(t *testing.T) {
	ConfirmString := func(s DoublyLinkedList, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v.String() expected %v but produced %v", s, r, x)
		}
	}

	ConfirmString(NewDoublyLinkedList(), "(list boolean ())")
	ConfirmString(NewDoublyLinkedList(false), "(list boolean (false))")
	ConfirmString(NewDoublyLinkedList(false, true), "(list boolean (false true))")
}