package cmplx64

import (
//	"../boolean"
	"math/cmplx"
	"testing"
)

func TestSliceString(t *testing.T) {
	ConfirmString := func(s Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(Slice{}, "()")
	ConfirmString(Slice{0}, "((0+0i))")
	ConfirmString(Slice{0, 1}, "((0+0i) (1+0i))")
	ConfirmString(Slice{0, 1i}, "((0+0i) (0+1i))")
}

func TestSliceEqual(t *testing.T) {
	ConfirmEqual := func(l Slice, r interface{}, ok bool) {
		if l.Equal(r) != ok {
			t.Fatalf("%v.Equal(%v) should be true", l, r)
		}

		x := []complex64(r.(Slice))
		if l.Equal(x) != ok {
			t.Fatalf("%v.Equal(%v) should be true", l, x)
		}
	}

	ConfirmEqual(Slice{}, Slice{}, true)
	ConfirmEqual(Slice{}, Slice{0}, false)
	ConfirmEqual(Slice{0}, Slice{}, false)
	ConfirmEqual(Slice{0}, Slice{0}, true)
	ConfirmEqual(Slice{0}, Slice{1}, false)
	ConfirmEqual(Slice{1}, Slice{0}, false)
	ConfirmEqual(Slice{1}, Slice{1}, true)

	ConfirmEqual(Slice{0, 1}, Slice{1}, false)
	ConfirmEqual(Slice{1, 0}, Slice{1}, false)
	ConfirmEqual(Slice{1, 0}, Slice{0, 1}, false)

	ConfirmEqual(Slice{0, 0}, Slice{0, 0}, true)
	ConfirmEqual(Slice{0, 0}, Slice{0, 1}, false)
	ConfirmEqual(Slice{0, 0}, Slice{1, 1}, false)
	ConfirmEqual(Slice{0, 0}, Slice{1, 0}, false)

	ConfirmEqual(Slice{0, 1}, Slice{0, 0}, false)
	ConfirmEqual(Slice{0, 1}, Slice{0, 1}, true)
	ConfirmEqual(Slice{0, 1}, Slice{1, 1}, false)
	ConfirmEqual(Slice{0, 1}, Slice{1, 0}, false)

	ConfirmEqual(Slice{1, 1}, Slice{0, 0}, false)
	ConfirmEqual(Slice{1, 1}, Slice{0, 1}, false)
	ConfirmEqual(Slice{1, 1}, Slice{1, 1}, true)
	ConfirmEqual(Slice{1, 1}, Slice{1, 0}, false)

	ConfirmEqual(Slice{1, 0}, Slice{0, 0}, false)
	ConfirmEqual(Slice{1, 0}, Slice{0, 1}, false)
	ConfirmEqual(Slice{1, 0}, Slice{1, 1}, false)
	ConfirmEqual(Slice{1, 0}, Slice{1, 0}, true)

	infinity := complex64(cmplx.Inf())
	nan := complex64(cmplx.NaN())
	ConfirmEqual(Slice{infinity}, Slice{infinity}, true)
	ConfirmEqual(Slice{infinity}, Slice{nan}, false)
	ConfirmEqual(Slice{nan}, Slice{infinity}, false)
	ConfirmEqual(Slice{nan}, Slice{nan}, false)
}

func TestSliceClone(t *testing.T) {
	ConfirmClone := func(s, r Slice) {
		if x := s.Clone(); !r.Equal(x) {
			t.Fatalf("%v.Clone() should be %v but is %v", s, r, x)
		}
	}

	ConfirmClone(Slice{}, Slice{})
	ConfirmClone(Slice{1i}, Slice{1i})
	ConfirmClone(Slice{1i, 2}, Slice{1i, 2})
}

func TestSliceMerge(t *testing.T) {
	ConfirmMerge := func(s, o Slice, f interface{}, r Slice) {
		x := s.Clone().(Slice)
		if x.Merge(o, f); !x.Equal(r) {
			t.Fatalf("%v.Merge(%v, %v) should be %v but is %v", s, o, f, r, x)
		}
	}

	ConfirmMergeAdd := func(s, o, r Slice) {
		ConfirmMerge(s, o, func(i, j complex64) complex64 { return i + j }, r)
	}

	ConfirmMergeAdd(Slice{}, Slice{}, Slice{})
	ConfirmMergeAdd(Slice{1}, Slice{}, Slice{1})
	ConfirmMergeAdd(Slice{}, Slice{1}, Slice{})
	ConfirmMergeAdd(Slice{1}, Slice{1}, Slice{2})
	ConfirmMergeAdd(Slice{1, 2}, Slice{2, 4}, Slice{3, 6})

	ConfirmMergeMultiply := func(s, o, r Slice) {
		ConfirmMerge(s, o, func(i, j complex64) complex64 { return i * j }, r)
	}

	ConfirmMergeMultiply(Slice{}, Slice{}, Slice{})
	ConfirmMergeMultiply(Slice{1}, Slice{}, Slice{1})
	ConfirmMergeMultiply(Slice{}, Slice{1}, Slice{})
	ConfirmMergeMultiply(Slice{1}, Slice{1}, Slice{1})
	ConfirmMergeMultiply(Slice{10}, Slice{10}, Slice{100})
	ConfirmMergeMultiply(Slice{1, 2, 4}, Slice{2, 4, 8}, Slice{2, 8, 32})
}

func TestSliceReduce(t *testing.T) {
	ConfirmReduce := func(s Slice, f interface{}, r complex64) {
		if x := s.Reduce(f); x != r {
			t.Fatalf("%v.Reduce(%v) should be %v but is %v", s, f, r, x)
		}
	}

	ConfirmReduceSum := func(s Slice, r complex64) {
		ConfirmReduce(s, func(memo, value complex64) complex64 { return memo + value }, r)
	}
	ConfirmReduceSum(Slice{0}, 0)
	ConfirmReduceSum(Slice{0, 1}, 1)
	ConfirmReduceSum(Slice{0, 1, 2}, 3)
	ConfirmReduceSum(Slice{0, 2, 1}, 3)
	ConfirmReduceSum(Slice{0, 1i}, 0+1i)
	ConfirmReduceSum(Slice{0, 1i, 2i}, 0+3i)

	ConfirmReduceProduct := func(s Slice, r complex64) {
		ConfirmReduce(s, func(memo, value complex64) complex64 { return memo * value }, r)
	}
	ConfirmReduceProduct(Slice{0}, 0)
	ConfirmReduceProduct(Slice{1}, 1)
	ConfirmReduceProduct(Slice{1, 2}, 2)
	ConfirmReduceProduct(Slice{1, 2, 3}, 6)
	ConfirmReduceProduct(Slice{1, 1i}, 0+1i)
	ConfirmReduceProduct(Slice{1, 1i, 2+2i}, -2+2i)
}

func TestSliceSum(t *testing.T) {
	ConfirmSum := func(s Slice, r complex64) {
		if x := s.Sum(); x != r {
			t.Fatalf("%v.Sum() should be %v but is %v", s, r, x)
		}
	}

	ConfirmSum(Slice{0}, 0)
	ConfirmSum(Slice{0, 1}, 1)
	ConfirmSum(Slice{0, 1, 2}, 3)
	ConfirmSum(Slice{0, 2, 1}, 3)
	ConfirmSum(Slice{0, 1i}, 0+1i)
	ConfirmSum(Slice{0, 1i, 2i}, 0+3i)
}

func TestSliceProduct(t *testing.T) {
	ConfirmProduct := func(s Slice, r complex64) {
		if x := s.Product(); x != r {
			t.Fatalf("%v.Product() should be %v but is %v", s, r, x)
		}
	}

	ConfirmProduct(Slice{0}, 0)
	ConfirmProduct(Slice{0, 1}, 0)
	ConfirmProduct(Slice{1, 2}, 2)
	ConfirmProduct(Slice{2, 1}, 2)
	ConfirmProduct(Slice{1i}, 0+1i)
	ConfirmProduct(Slice{1i, 2i}, -2)
	ConfirmProduct(Slice{1i, 2i, 3i}, -0-6i)
	ConfirmProduct(Slice{1i, 2i, 1+3i}, -2-6i)
}

func TestSliceDotProduct(t *testing.T) {
	ConfirmDotProduct := func(i, j Slice, r complex64) {
		if x := i.DotProduct(j); x != r {
			t.Fatalf("%v.DotProduct(%v) should be %v but is %v", i, j, r, x)
		}

		k := []complex64(j)
		if x := i.DotProduct(k); x != r {
			t.Fatalf("%v.DotProduct(%v) should be %v but is %v", i, k, r, x)
		}
	}

	ConfirmDotProduct(Slice{}, Slice{1}, 0)
	ConfirmDotProduct(Slice{1}, Slice{}, 0)

	ConfirmDotProduct(Slice{0}, Slice{0}, 0)
	ConfirmDotProduct(Slice{0}, Slice{1}, 0)
	ConfirmDotProduct(Slice{1}, Slice{0}, 0)

	ConfirmDotProduct(Slice{0, 1}, Slice{1, 0}, 0)
	ConfirmDotProduct(Slice{1, 0}, Slice{0, 1}, 0)
	ConfirmDotProduct(Slice{1, 0}, Slice{1, 0}, 1)
	ConfirmDotProduct(Slice{0, 1}, Slice{0, 1}, 1)
	ConfirmDotProduct(Slice{1, 1}, Slice{1, 1}, 2)

	ConfirmDotProduct(Slice{1, 1}, Slice{1}, 1)
	ConfirmDotProduct(Slice{1, 2}, Slice{1}, 1)
	ConfirmDotProduct(Slice{2, 1}, Slice{1}, 2)

	ConfirmDotProduct(Slice{1}, Slice{1, 1}, 1)
	ConfirmDotProduct(Slice{1}, Slice{2, 1}, 2)

	ConfirmDotProduct(Slice{1, 1}, Slice{1, 2}, 3)
	ConfirmDotProduct(Slice{1, 2}, Slice{1, 1}, 3)

	ConfirmDotProduct(Slice{1, 2}, Slice{3, 4}, 11)

	ConfirmDotProduct(Slice{1i}, Slice{1}, 1i)
	ConfirmDotProduct(Slice{1i, 2i}, Slice{1, 1}, 3i)
	ConfirmDotProduct(Slice{1i, 1+2i}, Slice{1, 1}, 1+3i)
}

func TestSliceNegate(t *testing.T) {
	ConfirmNegate := func(s, r Slice) {
		x := s.Clone().(Slice)
		if x.Negate(); !x.Equal(r) {
			t.Fatalf("%v.Negate() should be %v but is %v", s, r, x)
		}
	}

	ConfirmNegate(Slice{0}, Slice{0})
	ConfirmNegate(Slice{1}, Slice{-1})
	ConfirmNegate(Slice{0, 1}, Slice{0, -1})
	ConfirmNegate(Slice{0, 1, 2}, Slice{0, -1, -2})
	ConfirmNegate(Slice{0, 2, 1}, Slice{0, -2, -1})
	ConfirmNegate(Slice{0, 1i}, Slice{0, -1i})
	ConfirmNegate(Slice{0, 1i, 2i}, Slice{0, -1i, -2i})
}

func TestSliceIncrement(t *testing.T) {
	ConfirmIncrement := func(s, r Slice) {
		x := s.Clone().(Slice)
		if x.Increment(); !x.Equal(r) {
			t.Fatalf("%v.Increment() should be %v but is %v", s, r, x)
		}
	}

	ConfirmIncrement(Slice{0}, Slice{1})
	ConfirmIncrement(Slice{1}, Slice{2})
	ConfirmIncrement(Slice{0, 1}, Slice{1, 2})
	ConfirmIncrement(Slice{0, 1, 2}, Slice{1, 2, 3})
	ConfirmIncrement(Slice{0, 2, 1}, Slice{1, 3, 2})
	ConfirmIncrement(Slice{0, 1i}, Slice{1, 1+1i})
	ConfirmIncrement(Slice{0, 1i, 2i}, Slice{1, 1+1i, 1+2i})
}

func TestSliceDecrement(t *testing.T) {
	ConfirmDecrement := func(s, r Slice) {
		x := s.Clone().(Slice)
		if x.Decrement(); !x.Equal(r) {
			t.Fatalf("%v.Decrement() should be %v but is %v", s, r, x)
		}
	}

	ConfirmDecrement(Slice{0}, Slice{-1})
	ConfirmDecrement(Slice{1}, Slice{0})
	ConfirmDecrement(Slice{0, 1}, Slice{-1, 0})
	ConfirmDecrement(Slice{0, 1, 2}, Slice{-1, 0, 1})
	ConfirmDecrement(Slice{0, 2, 1}, Slice{-1, 1, 0})
	ConfirmDecrement(Slice{0, 1i}, Slice{-1, -1+1i})
	ConfirmDecrement(Slice{0, 1i, 2i}, Slice{-1, -1+1i, -1+2i})
}

func TestSliceAdd(t *testing.T) {
	ConfirmAdd := func(i, j, r Slice) {
		x := i.Clone().(Slice)
		if x.Add(j); !x.Equal(r) {
			t.Fatalf("%v.Add(%v) should be %v but is %v", i, j, r, x)
		}

		x = i.Clone().(Slice)
		k := []complex64(j)
		if x.Add(k); !x.Equal(r) {
			t.Fatalf("%v.Add(%v) should be %v but is %v", i, k, r, x)
		}
	}

	ConfirmAdd(Slice{}, Slice{}, Slice{})
	ConfirmAdd(Slice{}, Slice{1}, Slice{})
	ConfirmAdd(Slice{1}, Slice{}, Slice{1})

	ConfirmAdd(Slice{0}, Slice{0}, Slice{0})
	ConfirmAdd(Slice{0}, Slice{1}, Slice{1})
	ConfirmAdd(Slice{1}, Slice{0}, Slice{1})

	ConfirmAdd(Slice{0, 1}, Slice{1, 0}, Slice{1, 1})
	ConfirmAdd(Slice{1, 0}, Slice{0, 1}, Slice{1, 1})
	ConfirmAdd(Slice{1, 0}, Slice{1, 0}, Slice{2, 0})
	ConfirmAdd(Slice{0, 1}, Slice{0, 1}, Slice{0, 2})
	ConfirmAdd(Slice{1, 1}, Slice{1, 1}, Slice{2, 2})

	ConfirmAdd(Slice{1, 1}, Slice{1}, Slice{2, 1})
	ConfirmAdd(Slice{1, 2}, Slice{1}, Slice{2, 2})
	ConfirmAdd(Slice{2, 1}, Slice{1}, Slice{3, 1})

	ConfirmAdd(Slice{1}, Slice{1, 1}, Slice{2})
	ConfirmAdd(Slice{1}, Slice{2, 1}, Slice{3})

	ConfirmAdd(Slice{1, 1}, Slice{1, 2}, Slice{2, 3})
	ConfirmAdd(Slice{1, 2}, Slice{1, 1}, Slice{2, 3})

	ConfirmAdd(Slice{1i}, Slice{1}, Slice{1+1i})
	ConfirmAdd(Slice{1i, 2i}, Slice{1, 1}, Slice{1+1i, 1+2i})
	ConfirmAdd(Slice{1i, 1+2i}, Slice{1, 1}, Slice{1+1i, 2+2i})
}

func TestSliceSubtract(t *testing.T) {
	ConfirmSubtract := func(i, j, r Slice) {
		x := i.Clone().(Slice)
		if x.Subtract(j); !x.Equal(r) {
			t.Fatalf("%v.Subtract(%v) should be %v but is %v", i, j, r, x)
		}

		x = i.Clone().(Slice)
		k := []complex64(j)
		if x.Subtract(k); !x.Equal(r) {
			t.Fatalf("%v.Subtract(%v) should be %v but is %v", i, k, r, x)
		}
	}

	ConfirmSubtract(Slice{}, Slice{}, Slice{})
	ConfirmSubtract(Slice{}, Slice{1}, Slice{})
	ConfirmSubtract(Slice{1}, Slice{}, Slice{1})

	ConfirmSubtract(Slice{0}, Slice{0}, Slice{0})
	ConfirmSubtract(Slice{0}, Slice{1}, Slice{-1})
	ConfirmSubtract(Slice{1}, Slice{0}, Slice{1})

	ConfirmSubtract(Slice{0, 1}, Slice{1, 0}, Slice{-1, 1})
	ConfirmSubtract(Slice{1, 0}, Slice{0, 1}, Slice{1, -1})
	ConfirmSubtract(Slice{1, 0}, Slice{1, 0}, Slice{0, 0})
	ConfirmSubtract(Slice{0, 1}, Slice{0, 1}, Slice{0, 0})
	ConfirmSubtract(Slice{1, 1}, Slice{1, 1}, Slice{0, 0})

	ConfirmSubtract(Slice{1, 1}, Slice{1}, Slice{0, 1})
	ConfirmSubtract(Slice{1, 2}, Slice{1}, Slice{0, 2})
	ConfirmSubtract(Slice{2, 1}, Slice{1}, Slice{1, 1})

	ConfirmSubtract(Slice{1}, Slice{1, 1}, Slice{0})
	ConfirmSubtract(Slice{1}, Slice{2, 1}, Slice{-1})

	ConfirmSubtract(Slice{1, 1}, Slice{1, 2}, Slice{0, -1})
	ConfirmSubtract(Slice{1, 2}, Slice{1, 1}, Slice{0, 1})

	ConfirmSubtract(Slice{1i}, Slice{1}, Slice{-1+1i})
	ConfirmSubtract(Slice{1i, 2i}, Slice{1, 1}, Slice{-1+1i, -1+2i})
	ConfirmSubtract(Slice{1i, 1+2i}, Slice{1, 1}, Slice{-1+1i, 0+2i})
}

func TestSliceMultiply(t *testing.T) {
	ConfirmMultiply := func(i, j, r Slice) {
		x := i.Clone().(Slice)
		if x.Multiply(j); !x.Equal(r) {
			t.Fatalf("%v.Multiply(%v) should be %v but is %v", i, j, r, x)
		}

		x = i.Clone().(Slice)
		k := []complex64(j)
		if x.Multiply(k); !x.Equal(r) {
			t.Fatalf("%v.Multiply(%v) should be %v but is %v", i, k, r, x)
		}
	}

	ConfirmMultiply(Slice{}, Slice{}, Slice{})
	ConfirmMultiply(Slice{}, Slice{1}, Slice{})
	ConfirmMultiply(Slice{1}, Slice{}, Slice{1})

	ConfirmMultiply(Slice{0}, Slice{0}, Slice{0})
	ConfirmMultiply(Slice{0}, Slice{1}, Slice{0})
	ConfirmMultiply(Slice{1}, Slice{0}, Slice{0})
	ConfirmMultiply(Slice{1}, Slice{1}, Slice{1})
	ConfirmMultiply(Slice{10}, Slice{10}, Slice{100})
	ConfirmMultiply(Slice{10+10i}, Slice{10+10i}, Slice{200i})

	ConfirmMultiply(Slice{0, 1}, Slice{1, 0}, Slice{0, 0})
	ConfirmMultiply(Slice{1, 0}, Slice{0, 1}, Slice{0, 0})
	ConfirmMultiply(Slice{1, 0}, Slice{1, 0}, Slice{1, 0})
	ConfirmMultiply(Slice{0, 1}, Slice{0, 1}, Slice{0, 1})
	ConfirmMultiply(Slice{1, 1}, Slice{1, 1}, Slice{1, 1})
	ConfirmMultiply(Slice{10, 10}, Slice{1, 10}, Slice{10, 100})
	ConfirmMultiply(Slice{10, 10+10i}, Slice{1+1i, 10}, Slice{10+10i, 100+100i})

	ConfirmMultiply(Slice{10, 1}, Slice{1}, Slice{10, 1})
	ConfirmMultiply(Slice{10, 2}, Slice{1}, Slice{10, 2})
	ConfirmMultiply(Slice{10+10i, 1}, Slice{1}, Slice{10+10i, 1})

	ConfirmMultiply(Slice{10}, Slice{1, 1}, Slice{10})
	ConfirmMultiply(Slice{10}, Slice{2, 1}, Slice{20})

	ConfirmMultiply(Slice{1i}, Slice{1}, Slice{1i})
	ConfirmMultiply(Slice{1i, 2i}, Slice{1, 1}, Slice{1i, 2i})
	ConfirmMultiply(Slice{1i, 1+2i}, Slice{1, 1}, Slice{1i, 1+2i})

	ConfirmMultiply(Slice{1i}, Slice{1i}, Slice{-1})
	ConfirmMultiply(Slice{1i, 2i}, Slice{1i, 1i}, Slice{-1, -2})
	ConfirmMultiply(Slice{1i, 1+2i}, Slice{10i, 10i}, Slice{-10, -20+10i})
}

func TestSliceDivide(t *testing.T) {
	ConfirmDivide := func(i, j, r Slice) {
		x := i.Clone().(Slice)
		if x.Divide(j); !x.Equal(r) {
			t.Fatalf("%v.Divide(%v) should be %v but is %v", i, j, r, x)
		}

		x = i.Clone().(Slice)
		k := []complex64(j)
		if x.Divide(k); !x.Equal(r) {
			t.Fatalf("%v.Divide(%v) should be %v but is %v", i, k, r, x)
		}
	}

	ConfirmDivide(Slice{}, Slice{}, Slice{})
	ConfirmDivide(Slice{}, Slice{1}, Slice{})
	ConfirmDivide(Slice{1}, Slice{}, Slice{1})

	ConfirmDivide(Slice{0}, Slice{1}, Slice{0})
	ConfirmDivide(Slice{1, 1}, Slice{1, 1}, Slice{1, 1})

	ConfirmDivide(Slice{1, 1}, Slice{2}, Slice{0.5, 1})
	ConfirmDivide(Slice{1, 2}, Slice{2}, Slice{0.5, 2})
	ConfirmDivide(Slice{2, 1}, Slice{2}, Slice{1, 1})

	ConfirmDivide(Slice{1}, Slice{1, 1}, Slice{1})
	ConfirmDivide(Slice{1}, Slice{2, 1}, Slice{0.5})

	ConfirmDivide(Slice{1, 1}, Slice{1, 2}, Slice{1, 0.5})
	ConfirmDivide(Slice{1, 2}, Slice{1, 1}, Slice{1, 2})

	ConfirmDivide(Slice{1i}, Slice{1}, Slice{1i})
	ConfirmDivide(Slice{1i, 2i}, Slice{1, 1}, Slice{1i, 2i})
	ConfirmDivide(Slice{1i, 1+2i}, Slice{1, 1}, Slice{1i, 1+2i})
	ConfirmDivide(Slice{1i, 2i}, Slice{2, 2}, Slice{0.5i, 1i})
	ConfirmDivide(Slice{1i, 1+2i}, Slice{2, 2}, Slice{0.5i, 0.5+1i})
}

func TestSliceIsInf(t *testing.T) {
	ConfirmDivideIsInf := func(i, j Slice, r []bool) {
		x := i.Clone().(Slice)
		x.Divide(j)
		if b := x.IsInf(); !b.Equal(r) {
			t.Fatalf("%v.Divide(%v).IsInf() should be %v [%T] but is %v [%T]", i, j, r, r, b, b)
		}
	}

	ConfirmDivideIsInf(Slice{1}, Slice{0}, []bool{true})
	ConfirmDivideIsInf(Slice{0, 1}, Slice{1, 0}, []bool{false, true})
	ConfirmDivideIsInf(Slice{1, 0}, Slice{0, 1}, []bool{true, false})
}

func TestSliceIsNaN(t *testing.T) {
	ConfirmDivideIsNaN := func(i, j Slice, r []bool) {
		x := i.Clone().(Slice)
		x.Divide(j)
		if b := x.IsNaN(); !b.Equal(r) {
			t.Fatalf("%v.Divide(%v).IsNan() should be %v [%T] but is %v [%T]", i, j, r, r, b, b)
		}
	}

	ConfirmDivideIsNaN(Slice{1, 0}, Slice{1, 0}, []bool{false, true})
	ConfirmDivideIsNaN(Slice{0, 1}, Slice{0, 1}, []bool{true, false})
}