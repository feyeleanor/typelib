package boolean

import(
	"fmt"
	"strings"
)

type SinglyLinkedList struct {
	*link
}

func Cons(n ...bool) (r SinglyLinkedList) {
	if length := len(n); length > 0 {
		r = SinglyLinkedList{ link: &link{ value: n[0] } }
		end := r.link
		for _, v := range n[1:] {
			end = end.Append(v)
		}
	}
	return
}

func (s SinglyLinkedList) IsNil() bool {
	return s.link == nil
}

func (s SinglyLinkedList) End() *link {
	if end := s.link; end != nil {
		for ; end.next != nil; end = end.next {}
		return end
	}
	return nil
}

func (s SinglyLinkedList) String() string {
	elements := []string{}
	s.Each(func(v bool) {
		elements = append(elements, fmt.Sprintf("%v", v))
	})
	return fmt.Sprintf("(cons-boolean (%v))", strings.Join(elements, " "))
}

func (s SinglyLinkedList) equal(o SinglyLinkedList) (r bool) {
	o_end := o.link
	for end := s.link; end != nil; {
		if o_end == nil || end.value != o_end.value {
			return false
		}
		end = end.next
		o_end = o_end.next
	}
	return o_end == nil
}

func (s SinglyLinkedList) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case SinglyLinkedList:
		r = s.equal(o)
	case *link:
		r = s.equal(SinglyLinkedList{ link: o })
	case link:
		r = s.equal(SinglyLinkedList{ link: &o })
	}
	return
}

func (s SinglyLinkedList) Len() (i int) {
	for end := s.link; end != nil; i++ {
		end = end.next
	}
	return
}

func (s *SinglyLinkedList) Append(v... bool) {
	if s.IsNil() {
		switch len(v) {
		case 0:
		case 1:
			s.link = &link{ value: v[0] }
		default:
			s.link = &link{ value: v[0], next: Cons(v[1:]...).link }
		}
	} else {
		if len(v) > 0 {
			end := s.link
			for ; end.next != nil; end = end.next {}
			end.next = Cons(v...).link
		}
	}
}

func (s *SinglyLinkedList) Prepend(v... bool) {
	if len(v) > 0 {
		if n := Cons(v...); s.IsNil() {
			s.link = n.link
		} else {
			n.End().next = s.link
			s.link = n.link
		}
	}
}


func (s SinglyLinkedList) Clone() (r SinglyLinkedList) {
	if s.link != nil {
		l := &link{ value: s.value }
		r.link = l
		for end := s.next; end != nil; {
			l.next = &link{ value: end.value }
			l = l.next
			end = end.next
		}
	}
	return
}

func (s SinglyLinkedList) Each(f interface{}) bool {
	if s.link != nil {
		l := s.link
		switch f := f.(type) {
		case func(bool):
			for ; l != nil; l = l.next {
				f(l.value)
			}
		case func(int, bool):
			for i := 0; l != nil; l = l.next {
				f(i, l.value)
				i++
			}
		case func(interface{}, bool):
			for i := 0; l != nil; l = l.next {
				f(i, l.value)
				i++
			}
		case func(interface{}):
			for ; l != nil; l = l.next {
				f(l.value)
			}
		case func(int, interface{}):
			for i := 0; l != nil; l = l.next {
				f(i, l.value)
				i++
			}
		case func(interface{}, interface{}):
			for i := 0; l != nil; l = l.next {
				f(i, l.value)
				i++
			}
		default:
			return false
		}
		return true
	}
	return false
}