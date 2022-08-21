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

func assertNil(t *testing.T, actual any) {
	if !isNil(actual) {
		t.Errorf("Test %s: Expected value to be nil, Received `%v` (type %v)", t.Name(), actual, reflect.TypeOf(actual))
	}
}

func assertNotNil(t *testing.T, actual any) {
	if isNil(actual) {
		t.Errorf("Test %s: Expected value to not be nil, Received `%v` (type %v)", t.Name(), actual, reflect.TypeOf(actual))
	}
}

func assertEqual(t *testing.T, expected, actual any) {
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

func TestFilter(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := Filter(s1, func(i int) bool { return i > 2 })
	assertEqual(t, []int{3, 4, 5}, s2)
	assertEqual(t, []int{1, 2, 3, 4, 5}, s1)
	if unsafe.Pointer(&s1[0]) == unsafe.Pointer(&s2[0]) {
		t.Errorf("Test %s: Expected s1 and s2 to not be the same slice", t.Name())
	}
}

func TestFilterInPlace(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := FilterInPlace(s1, func(i int) bool { return i > 2 })
	assertEqual(t, []int{3, 4, 5}, s2)
	assertEqual(t, []int{3, 4, 5, 4, 5}, s1)
	assertEqual(t, unsafe.Pointer(&s1[0]), unsafe.Pointer(&s2[0]))
}

func TestMap(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := Map(s1, func(i int) int { return i * 2 })
	assertEqual(t, []int{2, 4, 6, 8, 10}, s2)
	assertEqual(t, []int{1, 2, 3, 4, 5}, s1)
	if unsafe.Pointer(&s1[0]) == unsafe.Pointer(&s2[0]) {
		t.Errorf("Test %s: Expected s1 and s2 to not be the same slice", t.Name())
	}
}

func TestReduce(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := Reduce(s1, func(acc, i int) int { return acc + i })
	assertEqual(t, 15, s2)

	s1 = []int{1}
	s2 = Reduce(s1, func(acc, i int) int { return acc + i })
	assertEqual(t, 1, s2)

	assertPanic(t, errEmptySlice, func() {
		Reduce([]int{}, func(acc, i int) int { return acc + i })
	})
}
