package slices

import (
	"reflect"
	"testing"
	"unsafe"
)

func isNil(i any) bool {
	if i == nil {
		return true
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Ptr,
		reflect.UnsafePointer,
		reflect.Interface,
		reflect.Slice:
		return v.IsNil()
	}

	return false
}

func assertNil[T any](t *testing.T, actual T) {
	if !isNil(actual) {
		t.Errorf("Test %s: Expected value to be nil, Received `%v` (type %v)", t.Name(), actual, reflect.TypeOf(actual))
	}
}

func assertNotNil[T any](t *testing.T, actual T) {
	if isNil(actual) {
		t.Errorf("Test %s: Expected value to not be nil, Received `%v` (type %v)", t.Name(), actual, reflect.TypeOf(actual))
	}
}

func assertEqual[T any](t *testing.T, expected, actual T) {
	if (isNil(expected) && isNil(actual)) || reflect.DeepEqual(expected, actual) {
		return
	}

	t.Errorf("Test %s: Expected `%v` (type %v), Received `%v` (type %v)", t.Name(), expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual))
}

func assertPanic(t *testing.T, expected any, f func()) {
	defer func() {
		if r := recover(); r == nil || r != expected {
			t.Errorf("Test %s: Expected Panic `%v` (type %v), Received Panic `%v` (type %v)", t.Name(), expected, reflect.TypeOf(expected), r, reflect.TypeOf(r))
		}
	}()
	f()
}

func TestIndex(t *testing.T) {
	tests := []struct {
		s    []int
		e, i int
	}{
		{s: nil, e: 0, i: -1},
		{s: []int{}, e: 0, i: -1},
		{s: []int{1, 2, 3}, e: 2, i: 1},
		{s: []int{1, 2, 2, 3}, e: 2, i: 1},
		{s: []int{1, 2, 3, 2}, e: 2, i: 1},
	}

	for _, test := range tests {
		assertEqual(t, test.i, Index(test.s, test.e))
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		s []int
		e int
		c bool
	}{
		{s: nil, e: 0, c: false},
		{s: []int{}, e: 0, c: false},
		{s: []int{1, 2, 3}, e: 2, c: true},
		{s: []int{1, 2, 2, 3}, e: 2, c: true},
		{s: []int{1, 2, 3, 2}, e: 2, c: true},
	}

	for _, test := range tests {
		assertEqual(t, test.c, Contains(test.s, test.e))
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		s, e []int
	}{
		{s: []int{1, 2, 3, 4, 5}, e: []int{3, 4, 5}},
	}

	for _, test := range tests {
		gt2 := Filter(test.s, func(i int) bool { return i > 2 })
		assertEqual(t, test.e, gt2)
		assertEqual(t, cap(test.e), cap(gt2))
		if unsafe.Pointer(&test.s[0]) == unsafe.Pointer(&gt2[0]) {
			t.Errorf("Test %s: Expected s1 and s2 to not be the same slice", t.Name())
		}
	}
}

func TestFilterInPlace(t *testing.T) {
	tests := []struct {
		s, e []int
	}{
		{s: []int{1, 2, 3, 4, 5}, e: []int{3, 4, 5}},
	}

	for _, test := range tests {
		gt2 := FilterInPlace(test.s, func(i int) bool { return i > 2 })
		assertEqual(t, test.e, gt2)
		assertEqual(t, cap(test.e), cap(gt2))
		assertEqual(t, unsafe.Pointer(&test.s[0]), unsafe.Pointer(&gt2[0]))
	}
}

func TestFind(t *testing.T) {
	type Person struct {
		firstname, lastname string
	}

	tests := []struct {
		s   []*Person
		fn  func(p *Person) bool
		e   *Person
		err error
	}{
		{
			s: []*Person{
				{
					firstname: "Grace",
					lastname:  "Hoper",
				},
				{
					firstname: "Jacob",
					lastname:  "Bernoulli",
				},
				{
					firstname: "Johann",
					lastname:  "Bernoulli",
				},
			},
			fn: func(p *Person) bool {
				return p.lastname == "Hoper"
			},
			e: &Person{
				firstname: "Grace",
				lastname:  "Hoper",
			},
			err: nil,
		},
		{
			s: []*Person{
				{
					firstname: "Grace",
					lastname:  "Hoper",
				},
				{
					firstname: "Jacob",
					lastname:  "Bernoulli",
				},
				{
					firstname: "Johann",
					lastname:  "Bernoulli",
				},
			},
			fn: func(p *Person) bool {
				return p.lastname == "Bernoulli"
			},
			e: &Person{
				firstname: "Jacob",
				lastname:  "Bernoulli",
			},
			err: nil,
		},
		{
			s: []*Person{
				{
					firstname: "Jacob",
					lastname:  "Bernoulli",
				},
				{
					firstname: "Johann",
					lastname:  "Bernoulli",
				},
			},
			fn: func(p *Person) bool {
				return p.lastname == "Hoper"
			},
			e:   nil,
			err: errElementNotFound,
		},
	}

	for _, test := range tests {
		p, err := Find(test.s, test.fn)
		assertEqual(t, test.e, p)
		assertEqual(t, test.err, err)
	}
}

func TestFindLast(t *testing.T) {
	type Person struct {
		firstname, lastname string
	}

	tests := []struct {
		s   []*Person
		fn  func(p *Person) bool
		e   *Person
		err error
	}{
		{
			s: []*Person{
				{
					firstname: "Grace",
					lastname:  "Hoper",
				},
				{
					firstname: "Jacob",
					lastname:  "Bernoulli",
				},
				{
					firstname: "Johann",
					lastname:  "Bernoulli",
				},
			},
			fn: func(p *Person) bool {
				return p.lastname == "Hoper"
			},
			e: &Person{
				firstname: "Grace",
				lastname:  "Hoper",
			},
			err: nil,
		},
		{
			s: []*Person{
				{
					firstname: "Grace",
					lastname:  "Hoper",
				},
				{
					firstname: "Jacob",
					lastname:  "Bernoulli",
				},
				{
					firstname: "Johann",
					lastname:  "Bernoulli",
				},
			},
			fn: func(p *Person) bool {
				return p.lastname == "Bernoulli"
			},
			e: &Person{
				firstname: "Johann",
				lastname:  "Bernoulli",
			},
			err: nil,
		},
		{
			s: []*Person{
				{
					firstname: "Jacob",
					lastname:  "Bernoulli",
				},
				{
					firstname: "Johann",
					lastname:  "Bernoulli",
				},
			},
			fn: func(p *Person) bool {
				return p.lastname == "Hoper"
			},
			e:   nil,
			err: errElementNotFound,
		},
	}

	for _, test := range tests {
		p, err := FindLast(test.s, test.fn)
		assertEqual(t, test.e, p)
		assertEqual(t, test.err, err)
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		s, e []int
	}{
		{s: []int{1, 2, 3, 4, 5}, e: []int{2, 4, 6, 8, 10}},
	}

	for _, test := range tests {
		times2 := Map(test.s, func(i int) int { return i * 2 })
		assertEqual(t, test.e, times2)
		assertEqual(t, cap(test.e), cap(times2))
		if unsafe.Pointer(&test.s[0]) == unsafe.Pointer(&times2[0]) {
			t.Errorf("Test %s: Expected s1 and s2 to not be the same slice", t.Name())
		}
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		s   []int
		sum int
	}{
		{s: []int{1}, sum: 1},
		{s: []int{1, 2, 3, 4, 5}, sum: 15},
	}

	for _, test := range tests {
		assertEqual(t, test.sum, Reduce(test.s, func(acc, i int) int { return acc + i }))
	}

	assertPanic(t, errEmptySlice, func() { Reduce(nil, func(acc, i int) int { return acc + i }) })
	assertPanic(t, errEmptySlice, func() { Reduce([]int{}, func(acc, i int) int { return acc + i }) })
}

func TestAll(t *testing.T) {
	tests := []struct {
		s []int
		e bool
	}{
		{s: nil, e: false},
		{s: []int{}, e: false},
		{s: []int{1, 2, 3, 4, 5}, e: false},
		{s: []int{2, 4, 6, 8, 10}, e: true},
	}

	for _, test := range tests {
		assertEqual(t, test.e, All(test.s, func(v int) bool { return v%2 == 0 }))
	}
}

func TestAny(t *testing.T) {
	tests := []struct {
		s []int
		e bool
	}{
		{s: nil, e: false},
		{s: []int{}, e: false},
		{s: []int{1, 2, 3, 4, 5}, e: true},
		{s: []int{1, 3, 5, 7, 9}, e: false},
	}

	for _, test := range tests {
		assertEqual(t, test.e, Any(test.s, func(v int) bool { return v%2 == 0 }))
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		s []int
		e uint
	}{
		{s: nil, e: uint(0)},
		{s: []int{}, e: uint(0)},
		{s: []int{1, 2, 3, 4, 5}, e: uint(2)},
		{s: []int{1, 3, 5, 7, 9}, e: uint(0)},
	}

	for _, test := range tests {
		assertEqual(t, test.e, Count(test.s, func(v int) bool { return v%2 == 0 }))
	}
}

func TestAssociateBy(t *testing.T) {
	type Person struct {
		firstname, lastname string
	}

	tests := []struct {
		s []*Person
		e map[string]*Person
	}{
		{
			s: []*Person{
				{
					firstname: "Grace",
					lastname:  "Hoper",
				},
				{
					firstname: "Jacob",
					lastname:  "Bernoulli",
				},
				{
					firstname: "Johann",
					lastname:  "Bernoulli",
				},
			},
			e: map[string]*Person{
				"Hoper":     {"Grace", "Hoper"},
				"Bernoulli": {"Johann", "Bernoulli"},
			},
		},
	}

	for _, test := range tests {
		assertEqual(t, test.e, AssociateBy(test.s, func(p *Person) string { return p.lastname }))
	}
}

func TestAssociateWith(t *testing.T) {
	type Person struct {
		firstname, lastname string
	}

	s := []string{"Hopper", "Bernoulli"}
	actual := AssociateWith(s, func(lastname string) *Person {
		switch lastname {
		case "Hopper":
			return &Person{"Grace", "Hoper"}
		case "Bernoulli":
			return &Person{"Johann", "Bernoulli"}
		default:
			return nil
		}
	})
	expected := map[string]*Person{
		"Hopper":    {"Grace", "Hoper"},
		"Bernoulli": {"Johann", "Bernoulli"},
	}
	assertEqual(t, expected, actual)
}

func TestGroupBy(t *testing.T) {
	type Person struct {
		firstname, lastname string
	}

	tests := []struct {
		s []*Person
		e map[string][]*Person
	}{
		{
			s: []*Person{
				{
					firstname: "Grace",
					lastname:  "Hoper",
				},
				{
					firstname: "Jacob",
					lastname:  "Bernoulli",
				},
				{
					firstname: "Johann",
					lastname:  "Bernoulli",
				},
			},
			e: map[string][]*Person{
				"Hoper":     {{"Grace", "Hoper"}},
				"Bernoulli": {{"Jacob", "Bernoulli"}, {"Johann", "Bernoulli"}},
			},
		},
	}

	for _, test := range tests {
		assertEqual(t, test.e, GroupBy(test.s, func(p *Person) string { return p.lastname }))
	}
}

func TestPartition(t *testing.T) {
	type Person struct {
		firstname, lastname string
	}

	tests := []struct {
		s, e1, e2 []*Person
	}{
		{
			s: []*Person{
				{
					firstname: "Grace",
					lastname:  "Hoper",
				},
				{
					firstname: "Jacob",
					lastname:  "Bernoulli",
				},
				{
					firstname: "Johann",
					lastname:  "Bernoulli",
				},
			},
			e1: []*Person{
				{
					firstname: "Grace",
					lastname:  "Hoper",
				},
			},
			e2: []*Person{
				{
					firstname: "Jacob",
					lastname:  "Bernoulli",
				},
				{
					firstname: "Johann",
					lastname:  "Bernoulli",
				},
			},
		},
	}

	for _, test := range tests {
		e1, e2 := Partition(test.s, func(p *Person) bool { return p.lastname == "Hoper" })
		assertEqual(t, test.e1, e1)
		assertEqual(t, test.e2, e2)
	}
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		s [][]int
		e []int
	}{
		{
			s: [][]int{{1, 2}, {3, 4}, {5}},
			e: []int{1, 2, 3, 4, 5},
		},
		{
			s: [][]int{{1, 2, 3, 4}},
			e: []int{1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		assertEqual(t, test.e, Flatten(test.s))
	}
}

func TestChunked(t *testing.T) {
	tests := []struct {
		s []int
		n int
		e [][]int
	}{
		{
			s: []int{1, 2, 3, 4, 5},
			n: 2,
			e: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			s: []int{1, 2, 3, 4},
			n: 4,
			e: [][]int{{1, 2, 3, 4}},
		},
	}

	for _, test := range tests {
		assertEqual(t, test.e, Chunked(test.s, test.n))
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		s, e []int
	}{
		{s: []int{1, 2, 2, 3, 3, 3}, e: []int{1, 2, 3}},
	}

	for _, test := range tests {
		unique := Unique(test.s)
		assertEqual(t, test.e, unique)
		if unsafe.Pointer(&test.s[0]) == unsafe.Pointer(&unique[0]) {
			t.Errorf("Test %s: Expected s1 and s2 to not be the same slice", t.Name())
		}
	}
}

func TestUniqueInPlace(t *testing.T) {
	tests := []struct {
		s, e []int
	}{
		{s: []int{1, 2, 2, 3, 3, 3}, e: []int{1, 2, 3}},
	}

	for _, test := range tests {
		unique := UniqueInPlace(test.s)
		assertEqual(t, test.e, unique)
		assertEqual(t, unsafe.Pointer(&test.s[0]), unsafe.Pointer(&unique[0]))
	}
}

func TestUniqueBy(t *testing.T) {
	type Person struct {
		firstname, lastname string
	}

	tests := []struct {
		s, e []*Person
	}{
		{
			s: []*Person{
				{
					firstname: "Grace",
					lastname:  "Hoper",
				},
				{
					firstname: "Jacob",
					lastname:  "Bernoulli",
				},
				{
					firstname: "Johann",
					lastname:  "Bernoulli",
				},
			},
			e: []*Person{
				{
					firstname: "Grace",
					lastname:  "Hoper",
				},
				{
					firstname: "Jacob",
					lastname:  "Bernoulli",
				},
			},
		},
	}

	for _, test := range tests {
		unique := UniqueBy(test.s, func(p *Person) string {
			return p.lastname
		})
		assertEqual(t, test.e, unique)
		assertEqual(t, cap(test.e), cap(unique))
		if unsafe.Pointer(&test.s[0]) == unsafe.Pointer(&unique[0]) {
			t.Errorf("Test %s: Expected s1 and s2 to not be the same slice", t.Name())
		}
	}
}

func TestUniqueByInPlace(t *testing.T) {
	type Person struct {
		firstname, lastname string
	}

	tests := []struct {
		s, e []*Person
	}{
		{
			s: []*Person{
				{
					firstname: "Grace",
					lastname:  "Hoper",
				},
				{
					firstname: "Jacob",
					lastname:  "Bernoulli",
				},
				{
					firstname: "Johann",
					lastname:  "Bernoulli",
				},
			},
			e: []*Person{
				{
					firstname: "Grace",
					lastname:  "Hoper",
				},
				{
					firstname: "Jacob",
					lastname:  "Bernoulli",
				},
			},
		},
	}

	for _, test := range tests {
		unique := UniqueByInPlace(test.s, func(p *Person) string {
			return p.lastname
		})
		assertEqual(t, test.e, unique)
		assertEqual(t, cap(test.e), cap(unique))
		assertEqual(t, unsafe.Pointer(&test.s[0]), unsafe.Pointer(&unique[0]))
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		s1, s2, e []int
	}{
		{s1: []int{}, s2: []int{}, e: []int{}},
		{s1: []int{1}, s2: []int{}, e: []int{}},
		{s1: []int{}, s2: []int{1}, e: []int{}},
		{s1: []int{1}, s2: []int{1}, e: []int{1}},
		{s1: []int{1, 2}, s2: []int{1}, e: []int{1}},
	}

	for _, test := range tests {
		assertEqual(t, test.e, Intersect(test.s1, test.s2))
	}
}

func TestDistinct(t *testing.T) {
	tests := []struct {
		s1, s2, e []int
	}{
		{s1: []int{}, s2: []int{}, e: []int{}},
		{s1: []int{1}, s2: []int{}, e: []int{1}},
		{s1: []int{}, s2: []int{1}, e: []int{1}},
		{s1: []int{1}, s2: []int{1}, e: []int{}},
		{s1: []int{1, 2}, s2: []int{1}, e: []int{2}},
	}

	for _, test := range tests {
		assertEqual(t, test.e, Distinct(test.s1, test.s2))
	}
}

func TestSumOf(t *testing.T) {
	tests := []struct {
		s   []int
		sum int
	}{
		{s: nil, sum: 0},
		{s: []int{}, sum: 0},
		{s: []int{1}, sum: 1},
		{s: []int{1, 2, 3, 4, 5}, sum: 15},
	}

	for _, test := range tests {
		assertEqual(t, test.sum, SumOf(test.s, func(i int) int { return i }))
	}
}

func TestMinOf(t *testing.T) {
	tests := []struct {
		s   []int
		min int
	}{
		{s: []int{1}, min: 1},
		{s: []int{1, 2, 3, 4, 5}, min: 1},
	}

	for _, test := range tests {
		assertEqual(t, test.min, MinOf(test.s, func(i int) int { return i }))
	}

	assertPanic(t, errEmptySlice, func() { MinOf(nil, func(i int) int { return i }) })
	assertPanic(t, errEmptySlice, func() { MinOf([]int{}, func(i int) int { return i }) })
}

func TestMaxOf(t *testing.T) {
	tests := []struct {
		s   []int
		max int
	}{
		{s: []int{1}, max: 1},
		{s: []int{1, 2, 3, 4, 5}, max: 5},
	}

	for _, test := range tests {
		assertEqual(t, test.max, MaxOf(test.s, func(i int) int { return i }))
	}

	assertPanic(t, errEmptySlice, func() { MinOf(nil, func(i int) int { return i }) })
	assertPanic(t, errEmptySlice, func() { MinOf([]int{}, func(i int) int { return i }) })
}
