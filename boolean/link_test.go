package boolean

import(
	"testing"
)

func TestLinkString(t *testing.T) {
	ConfirmString := func(l link, r string) {
		if x := l.String(); x != r {
			t.Fatalf("%v.String() expected %v but produced %v", l, r, x)
		}
	}

	ConfirmString(link{ value: false }, "false")
	ConfirmString(link{ value: true }, "true")
}

func TestLinkAppend(t *testing.T) {
	ConfirmAppend := func(l *link, v bool, r *link) {
		switch x := l.Append(v); {
		case x == nil:
			t.Fatalf("%v.Append(%v) failed to create a tail node", l, v)
		case x.value != v:
			t.Fatalf("%v.Append(%v) tail node value is %v", l, v, x.value)
		case x.value != r.value:
			t.Fatalf("%v.Append(%v) returned node value is %v but should be %v", l, v, x, r)
		}

	}

	ConfirmAppend((*link)(nil), true, &link{ value: true })
	ConfirmAppend(&link{ value: false}, true, &link{ value: true })
}