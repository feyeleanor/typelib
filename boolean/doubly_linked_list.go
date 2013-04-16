package boolean

import(
	"fmt"
	"strings"
)

type DoublyLinkedList struct {
	start 		*node
	end			*node
	length		int
}

func NewDoublyLinkedList(n ...bool) (r DoublyLinkedList) {
	if length := len(n); length > 0 {
		r.start = &node{ value: n[0] }
		r.end = r.start
		for _, v := range n[1:] {
			r.end = r.end.Append(v)
		}
		r.length = length
	}
	return
}

func (l DoublyLinkedList) String() string {
	elements := []string{}
	l.Each(func(v bool) {
		elements = append(elements, fmt.Sprintf("%v", v))
	})
	return fmt.Sprintf("(list boolean (%v))", strings.Join(elements, " "))
}

func (l DoublyLinkedList) Each(f interface{}) bool {
	if l.length > 0 {
		n := l.start
		switch f := f.(type) {
		case func(bool):
			for ; n != nil; n = n.tail {
				f(n.value)
			}
		case func(int, bool):
			for i := 0; n != nil; n = n.tail {
				f(i, n.value)
				i++
			}
		case func(interface{}, bool):
			for i := 0; n != nil; n = n.tail {
				f(i, n.value)
				i++
			}
		case func(interface{}):
			for ; n != nil; n = n.tail {
				f(n.value)
			}
		case func(int, interface{}):
			for i := 0; n != nil; n = n.tail {
				f(i, n.value)
				i++
			}
		case func(interface{}, interface{}):
			for i := 0; n != nil; n = n.tail {
				f(i, n.value)
				i++
			}
		default:
			return false
		}
		return true
	}
	return false
}

func (l DoublyLinkedList) ReverseEach(f interface{}) bool {
	if l.length > 0 {
		n := l.end
		switch f := f.(type) {
		case func(bool):
			for ; n.head != nil; {
				f(n.value)
				n = n.head
			}
		case func(int, bool):
			for i := 0; n.head != nil; i++ {
				f(i, n.value)
				n = n.head
			}
		case func(interface{}, bool):
			for i := 0; n.head != nil; i++ {
				f(i, n.value)
				n = n.head
			}
		case func(interface{}):
			for ; n.head != nil; {
				f(n.value)
				n = n.head
			}
		case func(int, interface{}):
			for i := 0; n.head != nil; i++ {
				f(i, n.value)
				n = n.head
			}
		case func(interface{}, interface{}):
			for i := 0; n.head != nil; i++ {
				f(i, n.value)
				n = n.head
			}
		default:
			return false
		}
		return true
	}
	return false
}