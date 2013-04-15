package typelib

type Equatable interface {
	Equal(o interface{}) bool
}

type Comparison int
const(
	LESS_THAN = Comparison(iota)
	EQUAL_TO
	GREATER_THAN
)

type Comparable interface {
	Compare(o interface{}) Comparison
}

type Cell interface {
	Content() interface{}
	Tail() Cell
}

type Offset int
const(
	PREVIOUS = Offset(iota)
	CURRENT
	NEXT
)