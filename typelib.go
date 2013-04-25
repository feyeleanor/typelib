package typelib

type Equatable interface {
	Equal(o interface{}) bool
}

type Referenceable interface {
	IsNil() bool
}

type Enumerable interface {
	Each(f interface{})
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