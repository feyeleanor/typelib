package boolean

import (
	"testing"
	"typelib/integer"
)

func TestSliceString(t *testing.T) {
	ConfirmString := func(s Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(Slice{}, "(boolean slice ())")
	ConfirmString(Slice{false}, "(boolean slice (false))")
	ConfirmString(Slice{false, true}, "(boolean slice (false true))")
}

func TestSliceEqual(t *testing.T) {
	ConfirmEqual := func(l, r Slice, ok bool) {
		if l.Equal(r) != ok {
			t.Fatalf("%v.Equal(%v) should be true", l, r)
		}

		x := []bool(r)
		if l.Equal(x) != ok {
			t.Fatalf("%v.Equal(%v) should be true", l, x)
		}
	}

	ConfirmEqual(Slice{}, Slice{}, true)
	ConfirmEqual(Slice{}, Slice{false}, false)
	ConfirmEqual(Slice{false}, Slice{}, false)
	ConfirmEqual(Slice{false}, Slice{false}, true)
	ConfirmEqual(Slice{false}, Slice{true}, false)
	ConfirmEqual(Slice{true}, Slice{false}, false)
	ConfirmEqual(Slice{true}, Slice{true}, true)

	ConfirmEqual(Slice{false, true}, Slice{true}, false)
	ConfirmEqual(Slice{true, false}, Slice{true}, false)
	ConfirmEqual(Slice{true, false}, Slice{false, true}, false)

	ConfirmEqual(Slice{false, false}, Slice{false, false}, true)
	ConfirmEqual(Slice{false, false}, Slice{false, true}, false)
	ConfirmEqual(Slice{false, false}, Slice{true, true}, false)
	ConfirmEqual(Slice{false, false}, Slice{true, false}, false)

	ConfirmEqual(Slice{false, true}, Slice{false, false}, false)
	ConfirmEqual(Slice{false, true}, Slice{false, true}, true)
	ConfirmEqual(Slice{false, true}, Slice{true, true}, false)
	ConfirmEqual(Slice{false, true}, Slice{true, false}, false)

	ConfirmEqual(Slice{true, true}, Slice{false, false}, false)
	ConfirmEqual(Slice{true, true}, Slice{false, true}, false)
	ConfirmEqual(Slice{true, true}, Slice{true, true}, true)
	ConfirmEqual(Slice{true, true}, Slice{true, false}, false)

	ConfirmEqual(Slice{true, false}, Slice{false, false}, false)
	ConfirmEqual(Slice{true, false}, Slice{false, true}, false)
	ConfirmEqual(Slice{true, false}, Slice{true, true}, false)
	ConfirmEqual(Slice{true, false}, Slice{true, false}, true)
}

func TestSliceClone(t *testing.T) {
	ConfirmClone := func(s, r Slice) {
		if x := s.Clone(); !r.Equal(x) {
			t.Fatalf("%v.Clone() should be %v but is %v", s, r, x)
		}
	}

	ConfirmClone(Slice{}, Slice{})
	ConfirmClone(Slice{true}, Slice{true})
	ConfirmClone(Slice{true, true}, Slice{true, true})
	ConfirmClone(Slice{true, false, true}, Slice{true, false, true})
}

func TestSliceLen(t *testing.T) {
	ConfirmLen := func(s Slice, r int) {
		if x := s.Len(); x != r {
			t.Fatalf("%v.Len() should be %v but is %v", s, r, x)
		}
	}

	ConfirmLen(Slice{}, 0)
	ConfirmLen(Slice{false}, 1)
	ConfirmLen(Slice{true}, 1)
	ConfirmLen(Slice{false, true}, 2)
}

func TestSliceCap(t *testing.T) {
	ConfirmCap := func(s Slice, r int) {
		if x := s.Cap(); x != r {
			t.Fatalf("%v.Cap() should be %v but is %v", s, r, x)
		}
	}

	ConfirmCap(make(Slice, 0, 0), 0)
	ConfirmCap(make(Slice, 0, 1), 1)
	ConfirmCap(make(Slice, 0, 2), 2)
}

func TestSliceAt(t *testing.T) {
	ConfirmAt := func(s Slice, i int, r bool) {
		if x := s.At(i); x != r {
			t.Fatalf("%v.At(%v) should be %v but is %v", s, i, r, x)
		}
	}

	ConfirmAt(Slice{true, false, true, false}, 0, true)
	ConfirmAt(Slice{true, false, true, false}, 1, false)
	ConfirmAt(Slice{true, false, true, false}, 2, true)
	ConfirmAt(Slice{true, false, true, false}, 3, false)

	ConfirmAtMany := func(s Slice, i []int, r interface{}) {
		switch x := s.At(i...).(type) {
		case bool:
			if x != r {
				t.Fatalf("%v.At(%v) should be %v but is %v", s, i, r, x)
			}
		case Slice:
			if !x.Equal(r) {
				t.Fatalf("%v.At(%v) should be %v but is %v", s, i, r, x)
			}
		}
	}

	ConfirmAtMany(Slice{true, false, true, false}, []int{0}, true)
	ConfirmAtMany(Slice{true, false, true, false}, []int{1}, false)
	ConfirmAtMany(Slice{true, false, true, false}, []int{2}, true)
	ConfirmAtMany(Slice{true, false, true, false}, []int{3}, false)
	ConfirmAtMany(Slice{true, false, true, false}, []int{0, 1}, Slice{true, false})
	ConfirmAtMany(Slice{true, false, true, false}, []int{1, 0}, Slice{false, true})
}

func TestSliceSelect(t *testing.T) {
	ConfirmSelect := func(s Slice, f interface{}, r Slice) {
		if x := s.Select(f).(Slice); !x.Equal(r) {
			t.Fatalf("%v.Select(%v) should be %v but is %v", s, f, r, x)
		}
	}

	ConfirmSelect(Slice{true, false, true, false, true}, func(x interface{}) bool { return x == true }, Slice{true, true, true})
	ConfirmSelect(Slice{true, false, true, false, true}, func(x interface{}) bool { return x == false }, Slice{false, false})

	ConfirmSelect(Slice{true, false, true, false, true}, func(x bool) bool { return x == true }, Slice{true, true, true})
	ConfirmSelect(Slice{true, false, true, false, true}, func(x bool) bool { return x == false }, Slice{false, false})
}

func TestSliceFind(t *testing.T) {
	ConfirmFind := func(s Slice, v interface{}, r integer.Slice) {
		if x := s.Find(v); !x.Equal(r) {
			t.Fatalf("%v.Find(%v) should be %v but is %v", s, v, r, x)
		}
	}

	ConfirmFind(Slice{true, false, true, false, true}, true, []int{0, 2, 4})
	ConfirmFind(Slice{true, false, true, false, true}, false, []int{1, 3})
	ConfirmFind(Slice{true, false, true, false, true}, func(x bool) bool { return x }, []int{0, 2, 4})
	ConfirmFind(Slice{true, false, true, false, true}, func(x bool) bool { return !x }, []int{1, 3})
	ConfirmFind(Slice{true, false, true, false, true}, func(x interface{}) bool { return x == true }, []int{0, 2, 4})
	ConfirmFind(Slice{true, false, true, false, true}, func(x interface{}) bool { return x == false }, []int{1, 3})

	ConfirmFindN := func(s Slice, v interface{}, n int, r integer.Slice) {
		if x := s.Find(v, n); !x.Equal(r) {
			t.Fatalf("%v.Find(%v, %v) should be %v but is %v", s, v, n, r, x)
		}
	}

	ConfirmFindN(Slice{true, false, true, false, true}, true, 2, []int{0, 2})
	ConfirmFindN(Slice{true, false, true, false, true}, false, 1, []int{1})
	ConfirmFindN(Slice{true, false, true, false, true}, func(x bool) bool { return x }, 2, []int{0, 2})
	ConfirmFindN(Slice{true, false, true, false, true}, func(x bool) bool { return !x }, 1, []int{1})
	ConfirmFindN(Slice{true, false, true, false, true}, func(x interface{}) bool { return x == true }, 2, []int{0, 2})
	ConfirmFindN(Slice{true, false, true, false, true}, func(x interface{}) bool { return x == false }, 1, []int{1})
}

func TestSliceSet(t *testing.T) {
	ConfirmSet := func(s Slice, i int, r bool) {
		x := s.Clone().(Slice)
		x.Set(i, r)
		if n := x.At(i); n != r {
			t.Fatalf("%v.Set(%v, %v) should be %v but is %v", s, i, r, r, n)
		}
	}

	ConfirmSet(Slice{false, false, false, false}, 0, true)
	ConfirmSet(Slice{true, false, false, false}, 0, false)
	ConfirmSet(Slice{false, false, false, false}, 1, true)
	ConfirmSet(Slice{false, true, false, false}, 1, false)
	ConfirmSet(Slice{false, false, false, false}, 2, true)
	ConfirmSet(Slice{false, false, true, false}, 2, false)
	ConfirmSet(Slice{false, false, false, false}, 3, true)
	ConfirmSet(Slice{false, false, false, true}, 3, false)

	ConfirmSetMany := func(s Slice, i int, v interface{}, r Slice) {
		x := s.Clone().(Slice)
		x.Set(i, v)
		if n := x[i:i + len(r)]; !n.Equal(r) {
			t.Fatalf("%v.Set(%v, %v) should be %v but is %v", s, i, v, r, n)
		}
	}

	ConfirmSetMany(Slice{false, false, false, false}, 0, Slice{true}, Slice{true, false, false, false})
	ConfirmSetMany(Slice{false, false, false, false}, 0, []bool{true}, Slice{true, false, false, false})
	ConfirmSetMany(Slice{false, false, false, false}, 0, Slice{true, true}, Slice{true, true, false, false})
	ConfirmSetMany(Slice{false, false, false, false}, 0, []bool{true, true}, Slice{true, true, false, false})
	ConfirmSetMany(Slice{false, false, false, false}, 0, Slice{false, true}, Slice{false, true, false, false})
	ConfirmSetMany(Slice{false, false, false, false}, 0, []bool{false, true}, Slice{false, true, false, false})
}

func TestSliceClear(t *testing.T) {
	ConfirmClear := func(s Slice, i, n int, r Slice) {
		x := s.Clone().(Slice)
		if x.Clear(i, n); !x.Equal(r) {
			t.Fatalf("%v.Clear(%v, %v) should be %v but is %v", s, i, n, r, x)
		}
	}

	ConfirmClear(Slice{true, true, true, true}, 0, 0, Slice{true, true, true, true})
	ConfirmClear(Slice{true, true, true, true}, 0, 1, Slice{false, true, true, true})
	ConfirmClear(Slice{true, true, true, true}, 0, 2, Slice{false, false, true, true})
	ConfirmClear(Slice{true, true, true, true}, 0, 3, Slice{false, false, false, true})
	ConfirmClear(Slice{true, true, true, true}, 0, 4, Slice{false, false, false, false})

	ConfirmClear(Slice{true, true, true, true}, 1, 2, Slice{true, false, false, true})
	ConfirmClear(Slice{true, true, true, true}, 3, 1, Slice{true, true, true, false})
	ConfirmClear(Slice{true, true, true, true}, 3, 0, Slice{true, true, true, true})
}

func TestSliceSwap(t *testing.T) {
	ConfirmSwap := func(s Slice, n []int, r Slice) {
		x := s.Clone().(Slice)
		if x.Swap(n...); !x.Equal(r) {
			t.Fatalf("%v.Swap(%v) should be %v but is %v", s, n, r, x)
		}
	}

	ConfirmSwap(Slice{true, false}, []int{0, 1}, Slice{false, true})
	ConfirmSwap(Slice{true, false, false}, []int{0, 1}, Slice{false, true, false})
	ConfirmSwap(Slice{true, false, true, false, true}, []int{0, 1, 2}, Slice{false, true, false, false, true})
	ConfirmSwap(Slice{true, true, true, false, false, false}, []int{0, 3, 3}, Slice{false, false, false, true, true, true})
}

func TestSliceCopy(t *testing.T) {
	ConfirmCopy := func(s Slice, destination, source, count int, r Slice) {
		x := s.Clone().(Slice)
		if x.Copy(destination, source, count); !x.Equal(r) {
			t.Fatalf("%v.Copy(%v, %v, %v) should be %v but is %v", s, destination, source, count, r, x)
		}
	}

	ConfirmCopy(Slice{}, 0, 0, 1, Slice{})
	ConfirmCopy(Slice{}, 1, 0, 1, Slice{})
	ConfirmCopy(Slice{}, 0, 1, 1, Slice{})

	ConfirmCopy(Slice{true, false, true, false, true}, 0, 0, 4, Slice{true, false, true, false, true})
	ConfirmCopy(Slice{true, false, true, false, true}, 4, 4, 4, Slice{true, false, true, false, true})
	ConfirmCopy(Slice{true, false, true, false, false}, 4, 0, 4, Slice{true, false, true, false, true})
	ConfirmCopy(Slice{true, false, true, false, false}, 5, 0, 4, Slice{true, false, true, false, false})
	ConfirmCopy(Slice{true, false, true, false, false}, 5, 5, 4, Slice{true, false, true, false, false})
	ConfirmCopy(Slice{true, false, true, false, false}, 3, 1, 4, Slice{true, false, true, false, true})
	ConfirmCopy(Slice{true, false, true, false, false}, 2, 4, 4, Slice{true, false, false, false, false})
}

func TestSliceMerge(t *testing.T) {
	ConfirmMerge := func(s, o Slice, f interface{}, r Slice) {
		x := s.Clone().(Slice)
		if x.Merge(o, f); !x.Equal(r) {
			t.Fatalf("%v.Merge(%v, %v) should be %v but is %v", s, o, f, r, x)
		}
	}

	ConfirmMergeAnd := func(s, o, r Slice) {
		ConfirmMerge(s, o, func(i, j bool) bool { return i && j }, r)
	}

	ConfirmMergeAnd(Slice{}, Slice{}, Slice{})
	ConfirmMergeAnd(Slice{true}, Slice{}, Slice{true})
	ConfirmMergeAnd(Slice{}, Slice{true}, Slice{})
	ConfirmMergeAnd(Slice{true}, Slice{true}, Slice{true})
	ConfirmMergeAnd(Slice{true, true}, Slice{true, true}, Slice{true, true})
	ConfirmMergeAnd(Slice{true, false}, Slice{true, true}, Slice{true, false})

	ConfirmMergeOr := func(s, o, r Slice) {
		ConfirmMerge(s, o, func(i, j bool) bool { return i || j }, r)
	}

	ConfirmMergeOr(Slice{}, Slice{}, Slice{})
	ConfirmMergeOr(Slice{true}, Slice{}, Slice{true})
	ConfirmMergeOr(Slice{}, Slice{true}, Slice{})
	ConfirmMergeOr(Slice{true}, Slice{true}, Slice{true})
	ConfirmMergeOr(Slice{true, true}, Slice{true, true}, Slice{true, true})
	ConfirmMergeOr(Slice{true, false}, Slice{true, true}, Slice{true, true})
	ConfirmMergeOr(Slice{false, false}, Slice{true, true}, Slice{true, true})
}

func TestSliceReduce(t *testing.T) {
	ConfirmReduce := func(s Slice, f interface{}, r bool) {
		if x := s.Reduce(f); x != r {
			t.Fatalf("%v.Reduce(%v) should be %v but is %v", s, f, r, x)
		}
	}

	ConfirmReduceAnd := func(s Slice, r bool) {
		ConfirmReduce(s, func(memo, value bool) bool { return memo && value }, r)
	}
	ConfirmReduceAnd(Slice{false}, false)
	ConfirmReduceAnd(Slice{false, true}, false)
	ConfirmReduceAnd(Slice{false, true, false}, false)
	ConfirmReduceAnd(Slice{true}, true)
	ConfirmReduceAnd(Slice{true, true}, true)
	ConfirmReduceAnd(Slice{true, true, true}, true)
	ConfirmReduceAnd(Slice{true, true, true, false}, false)

	ConfirmReduceOr := func(s Slice, r bool) {
		ConfirmReduce(s, func(memo, value bool) bool { return memo || value }, r)
	}
	ConfirmReduceOr(Slice{false}, false)
	ConfirmReduceOr(Slice{false, false}, false)
	ConfirmReduceOr(Slice{false, true}, true)
	ConfirmReduceOr(Slice{false, true, false}, true)
	ConfirmReduceOr(Slice{true}, true)
	ConfirmReduceOr(Slice{true, true}, true)
	ConfirmReduceOr(Slice{true, true, true}, true)
	ConfirmReduceOr(Slice{true, true, true, false}, true)
}

func TestSliceTrue(t *testing.T) {
	ConfirmTrue := func(s Slice, r bool) {
		if x := s.True(); x != r {
			t.Fatalf("%v.True() should be %v but is %v", s, r, x)
		}
	}

	ConfirmTrue(Slice{false}, false)
	ConfirmTrue(Slice{false, true}, false)
	ConfirmTrue(Slice{false, true, false}, false)
	ConfirmTrue(Slice{true}, true)
	ConfirmTrue(Slice{true, true}, true)
	ConfirmTrue(Slice{true, true, true}, true)
	ConfirmTrue(Slice{true, true, true, false}, false)
}

func TestSliceNegate(t *testing.T) {
	ConfirmNegate := func(s Slice, n []int, r Slice) {
		x := s.Clone().(Slice)
		if x.Negate(); !x.Equal(r) {
			t.Fatalf("%v.Negate() should be %v but is %v", s, r, x)
		}
	}

	ConfirmNegate(Slice{false}, nil, Slice{true})
	ConfirmNegate(Slice{true}, nil, Slice{false})
	ConfirmNegate(Slice{false, true}, nil, Slice{true, false})
	ConfirmNegate(Slice{false, true, false}, nil, Slice{true, false, true})
}

func TestSliceAnd(t *testing.T) {
	ConfirmAnd := func(i, j, r Slice) {
		x := i.Clone().(Slice)
		if x.And(j); !x.Equal(r) {
			t.Fatalf("%v.And(%v) should be %v but is %v", i, j, r, x)
		}

		x = i.Clone().(Slice)
		k := []bool(j)
		if x.And(k); !x.Equal(r) {
			t.Fatalf("%v.And(%v) should be %v but is %v", i, k, r, x)
		}
	}

	ConfirmAnd(Slice{}, Slice{}, Slice{})
	ConfirmAnd(Slice{}, Slice{true}, Slice{})
	ConfirmAnd(Slice{true}, Slice{}, Slice{true})

	ConfirmAnd(Slice{false}, Slice{false}, Slice{false})
	ConfirmAnd(Slice{false}, Slice{true}, Slice{false})
	ConfirmAnd(Slice{true}, Slice{false}, Slice{false})

	ConfirmAnd(Slice{false, false}, Slice{false, false}, Slice{false, false})
	ConfirmAnd(Slice{false, true}, Slice{true, false}, Slice{false, false})
	ConfirmAnd(Slice{true, false}, Slice{false, true}, Slice{false, false})
	ConfirmAnd(Slice{true, false}, Slice{true, false}, Slice{true, false})
	ConfirmAnd(Slice{false, true}, Slice{false, true}, Slice{false, true})
	ConfirmAnd(Slice{true, true}, Slice{true, true}, Slice{true, true})

	ConfirmAnd(Slice{true, true}, Slice{true}, Slice{true, true})
	ConfirmAnd(Slice{false, true}, Slice{true}, Slice{false, true})
}

func TestSliceOr(t *testing.T) {
	ConfirmOr := func(i, j, r Slice) {
		x := i.Clone().(Slice)
		if x.Or(j); !x.Equal(r) {
			t.Fatalf("%v.Or(%v) should be %v but is %v", i, j, r, x)
		}

		x = i.Clone().(Slice)
		k := []bool(j)
		if x.Or(k); !x.Equal(r) {
			t.Fatalf("%v.Or(%v) should be %v but is %v", i, k, r, x)
		}
	}

	ConfirmOr(Slice{}, Slice{}, Slice{})
	ConfirmOr(Slice{}, Slice{true}, Slice{})
	ConfirmOr(Slice{true}, Slice{}, Slice{true})

	ConfirmOr(Slice{false}, Slice{false}, Slice{false})
	ConfirmOr(Slice{false}, Slice{true}, Slice{true})
	ConfirmOr(Slice{true}, Slice{false}, Slice{true})

	ConfirmOr(Slice{false, false}, Slice{false, false}, Slice{false, false})
	ConfirmOr(Slice{false, true}, Slice{true, false}, Slice{true, true})
	ConfirmOr(Slice{true, false}, Slice{false, true}, Slice{true, true})
	ConfirmOr(Slice{true, false}, Slice{true, false}, Slice{true, false})
	ConfirmOr(Slice{false, true}, Slice{false, true}, Slice{false, true})
	ConfirmOr(Slice{true, true}, Slice{true, true}, Slice{true, true})

	ConfirmOr(Slice{true, true}, Slice{true}, Slice{true, true})
	ConfirmOr(Slice{false, true}, Slice{true}, Slice{true, true})
}

func TestSliceNot(t *testing.T) {
	ConfirmNot := func(i, j, r Slice) {
		x := i.Clone().(Slice)
		if x.Not(j); !x.Equal(r) {
			t.Fatalf("%v.Not(%v) should be %v but is %v", i, j, r, x)
		}

		x = i.Clone().(Slice)
		k := []bool(j)
		if x.Not(k); !x.Equal(r) {
			t.Fatalf("%v.Not(%v) should be %v but is %v", i, k, r, x)
		}
	}

	ConfirmNot(Slice{}, Slice{}, Slice{})
	ConfirmNot(Slice{}, Slice{true}, Slice{})
	ConfirmNot(Slice{true}, Slice{}, Slice{true})

	ConfirmNot(Slice{false}, Slice{false}, Slice{false})
	ConfirmNot(Slice{false}, Slice{true}, Slice{true})
	ConfirmNot(Slice{true}, Slice{false}, Slice{true})

	ConfirmNot(Slice{false, false}, Slice{false, false}, Slice{false, false})
	ConfirmNot(Slice{false, true}, Slice{true, false}, Slice{true, true})
	ConfirmNot(Slice{true, false}, Slice{false, true}, Slice{true, true})
	ConfirmNot(Slice{true, false}, Slice{true, false}, Slice{false, false})
	ConfirmNot(Slice{false, true}, Slice{false, true}, Slice{false, false})
	ConfirmNot(Slice{true, true}, Slice{true, true}, Slice{false, false})
	ConfirmNot(Slice{true, true}, Slice{false, false}, Slice{true, true})

	ConfirmNot(Slice{true, true}, Slice{true}, Slice{false, true})
	ConfirmNot(Slice{false, true}, Slice{true}, Slice{true, true})
}

func TestSliceRestrictTo(t *testing.T) {
	ConfirmRestrictTo := func(s Slice, i, j int, r Slice) {
		x := s.Clone().(Slice)
		if x.RestrictTo(i, j); !x.Equal(r) {
			t.Fatalf("%v.RestrictTo(%v, %v) should be %v but is %v", s, i, j, r, x)
		}
	}

	ConfirmRestrictTo(Slice{true, false, false, true}, 0, 1, Slice{true})
	ConfirmRestrictTo(Slice{true, false, false, true}, 0, 2, Slice{true, false})
	ConfirmRestrictTo(Slice{true, false, false, true}, 0, 3, Slice{true, false, false})
	ConfirmRestrictTo(Slice{true, false, false, true}, 0, 4, Slice{true, false, false, true})
	ConfirmRestrictTo(Slice{true, false, false, true}, 1, 2, Slice{false})
	ConfirmRestrictTo(Slice{true, false, false, true}, 1, 3, Slice{false, false})
}

func TestSliceCut(t *testing.T) {
	ConfirmCut := func(s Slice, i, j int, r Slice) {
		x := s.Clone().(Slice)
		if x.Cut(i, j); !x.Equal(r) {
			t.Fatalf("%v.Cut(%v, %v) should be %v but is %v", s, i, j, r, x)
		}
	}

	ConfirmCut(Slice{true, false, true, false, true, false}, 0, 1, Slice{false, true, false, true, false})
	ConfirmCut(Slice{true, false, true, false, true, false}, 1, 2, Slice{true, true, false, true, false})
	ConfirmCut(Slice{true, false, true, false, true, false}, 2, 3, Slice{true, false, false, true, false})
	ConfirmCut(Slice{true, false, true, false, true, false}, 3, 4, Slice{true, false, true, true, false})
	ConfirmCut(Slice{true, false, true, false, true, false}, 4, 5, Slice{true, false, true, false, false})
	ConfirmCut(Slice{true, false, true, false, true, false}, 5, 6, Slice{true, false, true, false, true})

	ConfirmCut(Slice{true, false, true, false, true, false}, -1, 1, Slice{false, true, false, true, false})
	ConfirmCut(Slice{true, false, true, false, true, false}, 0, 2, Slice{true, false, true, false})
	ConfirmCut(Slice{true, false, true, false, true, false}, 1, 3, Slice{true, false, true, false})
	ConfirmCut(Slice{true, false, true, false, true, false}, 2, 4, Slice{true, false, true, false})
	ConfirmCut(Slice{true, false, true, false, true, false}, 3, 5, Slice{true, false, true, false})
	ConfirmCut(Slice{true, false, true, false, true, false}, 4, 6, Slice{true, false, true, false})
	ConfirmCut(Slice{true, false, true, false, true, false}, 5, 7, Slice{true, false, true, false, true})
}

func TestSliceTrim(t *testing.T) {
	ConfirmTrim := func(s Slice, i, j int, r Slice) {
		x := s.Clone().(Slice)
		if x.Trim(i, j); !x.Equal(r) {
			t.Fatalf("%v.Trim(%v, %v) should be %v but is %v", s, i, j, r, x)
		}
	}

	ConfirmTrim(Slice{true, false, true}, 0, 1, Slice{true})
	ConfirmTrim(Slice{true, false, true}, 1, 2, Slice{false})
	ConfirmTrim(Slice{true, false, true}, 2, 3, Slice{true})

	ConfirmTrim(Slice{true, false, true}, -1, 1, Slice{true})
	ConfirmTrim(Slice{true, false, true}, 0, 2, Slice{true, false})
	ConfirmTrim(Slice{true, false, true}, 1, 3, Slice{false, true})
	ConfirmTrim(Slice{true, false, true}, 2, 4, Slice{true})
}

func TestSliceInsert(t *testing.T) {
	ConfirmInsert := func(s Slice, n int, v interface{}, r Slice) {
		x := s.Clone().(Slice)
		if x.Insert(n, v); !x.Equal(r) {
			t.Fatalf("%v.Insert(%v, %v) should be %v but is %v", s, n, v, r, x)
		}
	}

	ConfirmInsert(Slice{}, 0, false, Slice{false})
	ConfirmInsert(Slice{}, 0, Slice{false}, Slice{false})
	ConfirmInsert(Slice{}, 0, Slice{false, true}, Slice{false, true})

	ConfirmInsert(Slice{false}, 0, true, Slice{true, false})
	ConfirmInsert(Slice{false}, 1, true, Slice{false, true})
	ConfirmInsert(Slice{false}, 0, Slice{true}, Slice{true, false})
	ConfirmInsert(Slice{false}, 1, Slice{true}, Slice{false, true})

	ConfirmInsert(Slice{false, false, false}, 0, true, Slice{true, false, false, false})
	ConfirmInsert(Slice{false, false, false}, 1, true, Slice{false, true, false, false})
	ConfirmInsert(Slice{false, false, false}, 2, true, Slice{false, false, true, false})
	ConfirmInsert(Slice{false, false, false}, 3, true, Slice{false, false, false, true})

	ConfirmInsert(Slice{false, false, false}, 0, Slice{true, true}, Slice{true, true, false, false, false})
	ConfirmInsert(Slice{false, false, false}, 1, Slice{true, true}, Slice{false, true, true, false, false})
	ConfirmInsert(Slice{false, false, false}, 2, Slice{true, true}, Slice{false, false, true, true, false})
	ConfirmInsert(Slice{false, false, false}, 3, Slice{true, true}, Slice{false, false, false, true, true})
}

func TestSliceDelete(t *testing.T) {
	ConfirmDelete := func(s Slice, n []int, r Slice) {
		x := s.Clone().(Slice)
		if x.Delete(n...); !x.Equal(r) {
			t.Fatalf("%v.Delete(%v) should be %v but is %v", s, n, r, x)
		}
	}

	ConfirmDelete(Slice{true, false, true}, []int{-1}, Slice{true, false, true})
	ConfirmDelete(Slice{true, false, true}, []int{0}, Slice{false, true})
	ConfirmDelete(Slice{true, false, true}, []int{1}, Slice{true, true})
	ConfirmDelete(Slice{true, false, true}, []int{2}, Slice{true, false})
	ConfirmDelete(Slice{true, false, true}, []int{3}, Slice{true, false, true})

	ConfirmDelete(Slice{true, false, true}, []int{-1, 2}, Slice{true, false, true})
	ConfirmDelete(Slice{true, false, true}, []int{0, 2}, Slice{true})
	ConfirmDelete(Slice{true, false, true}, []int{1, 2}, Slice{true})
	ConfirmDelete(Slice{true, false, true}, []int{2, 2}, Slice{true, false, true})
}

func TestSliceDeleteIf(t *testing.T) {
	ConfirmDeleteIf := func(s Slice, f interface{}, r Slice) {
		x := s.Clone().(Slice)
		if x.DeleteIf(f); !x.Equal(r) {
			t.Fatalf("%v.DeleteIf(%v) should be %v but is %v", s, f, r, x)
		}
	}

	ConfirmDeleteIf(Slice{true, false, true}, false, Slice{true, true})
	ConfirmDeleteIf(Slice{true, false, true}, true, Slice{false})

	ConfirmDeleteIf(Slice{true, false, true}, func(x interface{}) bool { return x == false }, Slice{true, true})
	ConfirmDeleteIf(Slice{true, false, true}, func(x interface{}) bool { return x == true }, Slice{false})

	ConfirmDeleteIf(Slice{true, false, true}, func(x bool) bool { return x == false }, Slice{true, true})
	ConfirmDeleteIf(Slice{true, false, true}, func(x bool) bool { return x == true }, Slice{false})
}

func TestSliceKeepIf(t *testing.T) {
	ConfirmKeepIf := func(s Slice, f interface{}, r Slice) {
		x := s.Clone().(Slice)
		if x.KeepIf(f); !x.Equal(r) {
			t.Fatalf("%v.KeepIf(%v) should be %v but is %v", s, f, r, x)
		}
	}

	ConfirmKeepIf(Slice{true, false, true}, false, Slice{false})
	ConfirmKeepIf(Slice{true, false, true}, true, Slice{true, true})

	ConfirmKeepIf(Slice{true, false, true}, func(x interface{}) bool { return x == false }, Slice{false})
	ConfirmKeepIf(Slice{true, false, true}, func(x interface{}) bool { return x == true }, Slice{true, true})

	ConfirmKeepIf(Slice{true, false, true}, func(x bool) bool { return x == false }, Slice{false})
	ConfirmKeepIf(Slice{true, false, true}, func(x bool) bool { return x == true }, Slice{true, true})
}

func TestSliceEach(t *testing.T) {
	s := Slice{false, true, false, true, false, true}
	count := 0
	has_remainder := func(i int) bool {
		return (count % 2) == 1
	}

	s.Each(func(i interface{}) {
		if x := has_remainder(count); x != i {
			t.Fatalf("%v.Each(f(interface{})) element %v erroneously reported as %v", s, count, i)
		}
		count++
	})

	count = 0
	s.Each(func(index int, i interface{}) {
		switch x := has_remainder(count); {
		case index != count:
			t.Fatalf("%v.Each(f(int, interface{})) should have index %v but has %v", s, count, index)
		case x != i:
			t.Fatalf("%v.Each(f(int, interface{})) should contain %v but contains %v", s, x, i)
		}
		count++
	})

	count = 0
	s.Each(func(key, i interface{}) {
		switch x := has_remainder(count); {
		case key != count:
			t.Fatalf("%v.Each(f(interface{}, interface{})) should have key %v but has %v", s, count, key)
		case x != i:
			t.Fatalf("%v.Each(f(interface{}, interface{})) should contain %v but contains %v", s, x, i)
		}
		count++
	})

	count = 0
	s.Each(func(i bool) {
		if x := has_remainder(count); x != i {
			t.Fatalf("%v.Each(f(bool)) element %v erroneously reported as %v", s, count, i)
		}
		count++
	})

	count = 0
	s.Each(func(index int, i bool) {
		switch x := has_remainder(count); {
		case index != count:
			t.Fatalf("%v.Each(f(int, bool)) should have index %v but has %v", s, count, index)
		case x != i:
			t.Fatalf("%v.Each(f(int, bool)) should contain %v but contains %v", s, x, i)
		}
		count++
	})

	count = 0
	s.Each(func(key interface{}, i bool) {
		switch x := has_remainder(count); {
		case key != count:
			t.Fatalf("%v.Each(f(interface{}, bool)) should have key %v but has %v", s, count, key)
		case x != i:
			t.Fatalf("%v.Each(f(interface{}, bool)) should contain %v but contains %v", s, x, i)
		}
		count++
	})
}


func TestSliceReverseEach(t *testing.T) {
	s := Slice{false, true, false, true, false, true}
	count := len(s)
	has_remainder := func(i int) bool {
		return (count % 2) == 1
	}

	s.ReverseEach(func(i interface{}) {
		count--
		if x := has_remainder(count); x != i {
			t.Fatalf("%v.Each(f(interface{})) element %v erroneously reported as %v", s, count, i)
		}
	})

	count = len(s)
	s.ReverseEach(func(index int, i interface{}) {
		count--
		switch x := has_remainder(count); {
		case index != count:
			t.Fatalf("%v.Each(f(int, interface{})) should have index %v but has %v", s, count, index)
		case x != i:
			t.Fatalf("%v.Each(f(int, interface{})) should contain %v but contains %v", s, x, i)
		}
	})

	count = len(s)
	s.ReverseEach(func(key, i interface{}) {
		count--
		switch x := has_remainder(count); {
		case key != count:
			t.Fatalf("%v.Each(f(interface{}, interface{})) should have key %v but has %v", s, count, key)
		case x != i:
			t.Fatalf("%v.Each(f(interface{}, interface{})) should contain %v but contains %v", s, x, i)
		}
	})

	count = len(s)
	s.ReverseEach(func(i bool) {
		count--
		if x := has_remainder(count); x != i {
			t.Fatalf("%v.Each(f(bool)) element %v erroneously reported as %v", s, count, i)
		}
	})

	count = len(s)
	s.ReverseEach(func(index int, i bool) {
		count--
		switch x := has_remainder(count); {
		case index != count:
			t.Fatalf("%v.Each(f(int, bool)) should have index %v but has %v", s, count, index)
		case x != i:
			t.Fatalf("%v.Each(f(int, bool)) should contain %v but contains %v", s, x, i)
		}
	})

	count = len(s)
	s.ReverseEach(func(key interface{}, i bool) {
		count--
		switch x := has_remainder(count); {
		case key != count:
			t.Fatalf("%v.Each(f(interface{}, bool)) should have key %v but has %v", s, count, key)
		case x != i:
			t.Fatalf("%v.Each(f(interface{}, bool)) should contain %v but contains %v", s, x, i)
		}
	})
}

func TestSliceWhile(t *testing.T) {
	ConfirmLimit := func(s Slice, l int, f interface{}) {
		if count := s.While(f); count != l {
			t.Fatalf("%v.While() should have iterated %v times not %v times", s, l, count)
		}
	}

	s := Slice{true, true, true, true, false, true, true}
	limit := 4

	ConfirmLimit(s, limit, func(i interface{}) bool {
		return i.(bool)
	})

	ConfirmLimit(s, limit, func(index int, i interface{}) bool {
		return i.(bool)
	})

	ConfirmLimit(s, limit, func(index int, i interface{}) bool {
		return index != limit
	})

	ConfirmLimit(s, limit, func(key, i interface{}) bool {
		return i.(bool)
	})

	ConfirmLimit(s, limit, func(key, i interface{}) bool {
		return limit != key
	})

	ConfirmLimit(s, limit, func(i bool) bool {
		return i
	})

	ConfirmLimit(s, limit, func(index int, i bool) bool {
		return i
	})

	ConfirmLimit(s, limit, func(index int, i bool) bool {
		return index != limit
	})

	ConfirmLimit(s, limit, func(key interface{}, i bool) bool {
		return i
	})

	ConfirmLimit(s, limit, func(key interface{}, i bool) bool {
		return key.(int) != limit
	})
}

func TestSliceUntil(t *testing.T) {
	ConfirmLimit := func(s Slice, l int, f interface{}) {
		if count := s.Until(f); count != l {
			t.Fatalf("%v.Until() should have iterated %v times not %v times", s, l, count)
		}
	}

	s := Slice{false, false, false, false, true, false, false}
	limit := 4

	ConfirmLimit(s, limit, func(i interface{}) bool {
		return i.(bool)
	})

	ConfirmLimit(s, limit, func(index int, i interface{}) bool {
		return i.(bool)
	})

	ConfirmLimit(s, limit, func(index int, i interface{}) bool {
		return index == limit
	})

	ConfirmLimit(s, limit, func(key, i interface{}) bool {
		return i.(bool)
	})

	ConfirmLimit(s, limit, func(key, i interface{}) bool {
		return limit == key
	})

	ConfirmLimit(s, limit, func(i bool) bool {
		return i
	})

	ConfirmLimit(s, limit, func(index int, i bool) bool {
		return i
	})

	ConfirmLimit(s, limit, func(index int, i bool) bool {
		return index == limit
	})

	ConfirmLimit(s, limit, func(key interface{}, i bool) bool {
		return i
	})

	ConfirmLimit(s, limit, func(key interface{}, i bool) bool {
		return key.(int) == limit
	})
}

func TestSliceReverse(t *testing.T) {
	ConfirmReverse := func(s, r Slice) {
		x := s.Clone().(Slice)
		if x.Reverse(); !x.Equal(r) {
			t.Fatalf("%v.Reverse() should be %v but is %v", s, r, x)
		}
	}
	ConfirmReverse(Slice{}, Slice{})
	ConfirmReverse(Slice{true}, Slice{true})
	ConfirmReverse(Slice{true, false}, Slice{false, true})
	ConfirmReverse(Slice{true, false, false}, Slice{false, false, true})
	ConfirmReverse(Slice{true, false, true, false}, Slice{false, true, false, true})
}

func TestSliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s Slice, count int, r Slice) {
		x := s.Clone().(Slice)
		if x.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(Slice{}, 5, Slice{})
	ConfirmRepeat(Slice{true}, 1, Slice{true})
	ConfirmRepeat(Slice{true}, 2, Slice{true, true})
	ConfirmRepeat(Slice{false}, 3, Slice{false, false, false})
}

func TestSliceUniq(t *testing.T) {
	ConfirmUniq := func(s Slice, r Set) {
		if x := s.Uniq().(Set); !x.Equal(r) {
			t.Fatalf("%v.Uniq() should be %v but is %v", s, r, x)
		}
	}

	ConfirmUniq(Slice{false, false, false}, NewSet(false))
	ConfirmUniq(Slice{true, true, true}, NewSet(true))
	ConfirmUniq(Slice{false, true, false, true}, NewSet(false, true))
}

func TestSliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s Slice, l int, r Slice) {
		x := s.Clone().(Slice)
		y := &x
		switch y.Reallocate(l); {
		case y == nil:				t.Fatalf("%v.Reallocate(%v) created a nil value for Slice", s, l)
		case y.Len() != l:			t.Fatalf("%v.Reallocate(%v) length should be %v but is %v", s, l, r.Len(), y.Len())
		case !y.Equal(r):			t.Fatalf("%v.Reallocate(%v) should be %v but is %v", s, l, r, y)
		case x == nil:				t.Fatalf("%v.Reallocate(%v) created a nil value for Slice", s, l)
		case x.Len() != l:			t.Fatalf("%v.Reallocate(%v) length should be %v but is %v", s, l, r.Len(), x.Len())
		case !x.Equal(r):			t.Fatalf("%v.Reallocate(%v) should be %v but is %v", s, l, r, x)
		}
	}

	ConfirmReallocate(Slice{}, 0, Slice{})
	ConfirmReallocate(Slice{true, false, true, false}, 3, Slice{true, false, true})
	ConfirmReallocate(Slice{true, false, true, false}, 4, Slice{true, false, true, false})
	ConfirmReallocate(Slice{true, false, true, false}, 5, Slice{true, false, true, false, false})


	ConfirmReallocateCapacity := func(s Slice, l, c int, r Slice) {
		el := l
		if el > c {
			el = c
		}
		x := s.Clone().(Slice)
		y := &x
		switch y.Reallocate(l, c); {
		case y == nil:				t.Fatalf("%v.Reallocate(%v, %v) created a nil value for Slice", s, l, c)
		case y.Cap() != c:			t.Fatalf("%v.Reallocate(%v, %v) capacity should be %v but is %v", s, l, c, c, s.Cap())
		case y.Len() != el:			t.Fatalf("%v.Reallocate(%v, %v) length should be %v but is %v", s, l, c, el, s.Len())
		case !y.Equal(r):			t.Fatalf("%v.Reallocate(%v, %v) should be %v but is %v", s, l, c, r, y)
		case x == nil:				t.Fatalf("%v.Reallocate(%v, %v) created a nil value for Slice", s, l, c)
		case x.Cap() != c:			t.Fatalf("%v.Reallocate(%v, %v) capacity should be %v but is %v", s, l, c, c, s.Cap())
		case x.Len() != el:			t.Fatalf("%v.Reallocate(%v, %v) length should be %v but is %v", s, l, c, el, s.Len())
		case !x.Equal(r):			t.Fatalf("%v.Reallocate(%v, %v) should be %v but is %v", s, l, c, r, y)
		}
	}

	ConfirmReallocateCapacity(Slice{}, 0, 10, make(Slice, 0, 10))
	ConfirmReallocateCapacity(Slice{true, true, true, true, true}, 3, 10, Slice{true, true, true})
	ConfirmReallocateCapacity(Slice{true, true, true, true, true}, 5, 10, Slice{true, true, true, true, true})
	ConfirmReallocateCapacity(Slice{true, true, true, true, true}, 10, 10, Slice{true, true, true, true, true, false, false, false, false, false})
	ConfirmReallocateCapacity(Slice{true, true, true, true, true}, 1, 3, Slice{true})
	ConfirmReallocateCapacity(Slice{true, true, true, true, true}, 5, 5, Slice{true, true, true, true, true})
	ConfirmReallocateCapacity(Slice{true, true, true, true, true, true, true, true, true, true}, 10, 5, Slice{true, true, true, true, true})
}

func TestSliceExpand(t *testing.T) {
	ConfirmExpand := func(s Slice, n int, r Slice) {
		c := s.Cap()
		x := s.Clone().(Slice)
		y := &x
		switch y.Expand(n); {
		case y.Len() != r.Len():	t.Fatalf("%v.Extend(%v) len should be %v but is %v", s, n, r.Len(), y.Len())
		case y.Cap() != c + n:		t.Fatalf("%v.Extend(%v) cap should be %v but is %v", s, n, c + n, y.Cap())
		case !y.Equal(r):			t.Fatalf("%v.Extend(%v) should be %v but is %v", s, n, r, y)
		case x.Len() != r.Len():	t.Fatalf("%v.Extend(%v) len should be %v but is %v", s, n, r.Len(), x.Len())
		case x.Cap() != c + n:		t.Fatalf("%v.Extend(%v) cap should be %v but is %v", s, n, c + n, x.Cap())
		case !x.Equal(r):			t.Fatalf("%v.Extend(%v) should be %v but is %v", s, n, r, x)
		}
	}

	ConfirmExpand(Slice{}, 1, Slice{false})
	ConfirmExpand(Slice{}, 2, Slice{false, false})
	ConfirmExpand(Slice{true}, 2, Slice{true, false, false})
	ConfirmExpand(Slice{true, true}, 2, Slice{true, true, false, false})

	ConfirmExpandInline := func(s Slice, i, l int, r Slice) {
		c := s.Cap()
		x := s.Clone().(Slice)
		y := &x
		switch y.Expand(i, l); {
		case y.Len() != r.Len():	t.Fatalf("%v.Extend(%v, %v) len should be %v but is %v", s, i, l, r.Len(), y.Len())
		case y.Cap() != c + l:		t.Fatalf("%v.Extend(%v, %v) cap should be %v but is %v", s, i, l, c + l, y.Cap())
		case !y.Equal(r):			t.Fatalf("%v.Extend(%v, %v) should be %v but is %v", s, i, l, r, y)
		case x.Len() != r.Len():	t.Fatalf("%v.Extend(%v, %v) len should be %v but is %v", s, i, l, r.Len(), x.Len())
		case x.Cap() != c + l:		t.Fatalf("%v.Extend(%v, %v) cap should be %v but is %v", s, i, l, c + l, x.Cap())
		case !x.Equal(r):			t.Fatalf("%v.Extend(%v, %v) should be %v but is %v", s, i, l, r, x)
		}
	}

	ConfirmExpandInline(Slice{}, 0, 1, Slice{false})
	ConfirmExpandInline(Slice{}, 1, 1, Slice{false})
	ConfirmExpandInline(Slice{}, 1, 2, Slice{false, false})
	ConfirmExpandInline(Slice{}, 2, 3, Slice{false, false, false})
	ConfirmExpandInline(Slice{true, true}, 1, 3, Slice{true, false, false, false, true})
}

func TestSliceAppend(t *testing.T) {
	ConfirmAppend := func(s Slice, v interface{}, r Slice) {
		x := s.Clone().(Slice)
		if x.Append(v); !x.Equal(r) {
			t.Fatalf("%v.Append(%v) should be %v but is %v", s, v, r, x)
		}
	}

	ConfirmAppend(Slice{}, true, Slice{true})
	ConfirmAppend(Slice{false}, true, Slice{false, true})

	ConfirmAppend(Slice{}, Slice{true}, Slice{true})
	ConfirmAppend(Slice{}, Slice{false, true}, Slice{false, true})
	ConfirmAppend(Slice{true, true, true}, Slice{false, false}, Slice{true, true, true, false, false})

	ConfirmAppend(Slice{}, []bool{true}, Slice{true})
	ConfirmAppend(Slice{}, []bool{false, true}, Slice{false, true})
	ConfirmAppend(Slice{true, true, true}, []bool{false, false}, Slice{true, true, true, false, false})
}

func TestSlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s Slice, v interface{}, r Slice) {
		x := s.Clone().(Slice)
		if x.Prepend(v); !x.Equal(r) {
			t.Fatalf("%v.Prepend(%v) should be %v but is %v", s, v, r, x)
		}
	}

	ConfirmPrepend(Slice{}, false, Slice{false})
	ConfirmPrepend(Slice{false}, true, Slice{true, false})

	ConfirmPrepend(Slice{}, Slice{false}, Slice{false})
	ConfirmPrepend(Slice{}, Slice{false, true}, Slice{false, true})
	ConfirmPrepend(Slice{false, false}, Slice{true}, Slice{true, false, false})
	ConfirmPrepend(Slice{false, false}, Slice{true, true}, Slice{true, true, false, false})

	ConfirmPrepend(Slice{}, []bool{false}, Slice{false})
	ConfirmPrepend(Slice{}, []bool{false, true}, Slice{false, true})
	ConfirmPrepend(Slice{false, false}, []bool{true}, Slice{true, false, false})
	ConfirmPrepend(Slice{false, false}, []bool{true, true}, Slice{true, true, false, false})
}