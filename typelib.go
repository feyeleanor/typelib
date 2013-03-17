package typelib

type Equatable interface {
	Equal(o interface{}) bool
}

type Comparison int
const(
	LESS_THAN = iota
	EQUAL
	GREATER_THAN
)

type Comparable interface {
	Compare(o interface{}) Comparison
}