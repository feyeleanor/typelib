package typelib

type Container interface {
	At(x... interface{}) interface{}
	Store(value interface{}, x... interface{})
}

type List interface {
	Equatable
	Referenceable
	Enumerable
	Container
	String() string
	Len() int
	Append(v interface{}) List
	Prepend(v interface{}) List
	Clone() List
	Collect(f interface{}) List
	Delete(f interface{}) List
	Reduce(f interface{}) List
	Reverse() List
}