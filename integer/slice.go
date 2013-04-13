package integer

import (
	"fmt"
	"strings"
)

type	Slice	[]int

const	ZERO = int(0)

func (s Slice) String() string {
	elements := []string{}
	for _, v := range s {
		elements = append(elements, fmt.Sprintf("%v", v))
	}
	return fmt.Sprintf("(%v)", strings.Join(elements, " "))
}

func (s Slice) equal(o Slice) (r bool) {
	if len(s) == len(o) {
		r = true
		for i, v := range s {
			if r = v == o[i]; !r {
				return
			}
		}
	}
	return
}

func (s Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case Slice:
		r = s.equal(o)
	case []int:
		r = s.equal(o)
	}
	return
}
