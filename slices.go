package slices

import "errors"

// number is a constraint that permits any numeric type: any type
// that supports the operators + - * / %.
type number interface {
	/* Signed */ ~int | ~int8 | ~int16 | ~int32 | ~int64 | /* Unsigned */ ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | /* Float */ ~float32 | ~float64 | /* Complex */ ~complex64 | ~complex128
}

var (
	errEmptySlice      = errors.New("slices: empty slice")
	errElementNotFound = errors.New("slices: no such element")
)

// Index returns the index of the first occurrence of v in e,
// or -1 if not present.
func Index[E comparable](s []E, v E) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}

// Contains reports whether v is present in e.
func Contains[E comparable](s []E, v E) bool {
	return Index(s, v) >= 0
}

// Filter executes the function fn to each element of the slice e
// returning a newly allocated slice of all elements for which the
// function fn returns true.
func Filter[E any](s []E, fn func(e E) bool) []E {
	r := make([]E, 0, len(s))
	for _, e := range s {
		if fn(e) {
			r = append(r, e)
		}
	}
	return r[:len(r):len(r)]
}

// FilterInPlace executes the function fn to each element of the slice e
// returning a slice of all elements for which the function fn returns true.
//
// It modifies the underlying array of slice e. Thus, this method should only
// be used if the passed slice e is not used afterwards!
func FilterInPlace[E any](s []E, fn func(e E) bool) []E {
	n := 0
	for _, e := range s {
		if fn(e) {
			s[n] = e
			n++
		}
	}
	return s[:n:n]
}

// Map applies the function fn to each element of the slice e.
// It returns a newly allocated slice with same length as e where
// each element is the result of calling the function fn on successive
// elements of the slice.
func Map[E1, E2 any](s []E1, fn func(e E1) E2) []E2 {
	r := make([]E2, len(s))
	for i, e := range s {
		r[i] = fn(e)
	}
	return r
}

// Reduce computes the reduction of the function fn across the
// elements of the slice e.
//
// If the slice is empty, Reduce will panic; if it has only one element,
// it returns that element.
func Reduce[E any](s []E, fn func(acc, e E) E) E {
	if len(s) == 0 {
		panic(errEmptySlice)
	}
	acc := s[0]
	for _, e := range s[1:] {
		acc = fn(acc, e)
	}
	return acc
}

// All returns true if the evaluation of the predicate function fn
// returns true for all elements of the slice e.
func All[E any](s []E, fn func(e E) bool) bool {
	if len(s) == 0 {
		return false
	}
	for _, e := range s {
		if !fn(e) {
			return false
		}
	}
	return true
}

// Any returns true if the evaluation of the predicate function fn
// returns true for at least one element of the slice e.
func Any[E any](s []E, fn func(e E) bool) bool {
	for _, e := range s {
		if fn(e) {
			return true
		}
	}
	return false
}

// Count returns an integer value indicating how many elements
// of the slice e yield true for the predicate function fn.
func Count[E any](s []E, fn func(e E) bool) uint {
	i := uint(0)
	for _, e := range s {
		if fn(e) {
			i++
		}
	}
	return i
}

// AssociateBy returns a map from the elements of the slice e as values
// with the key retrieved by applying the given function fn.
func AssociateBy[E any, K comparable](s []E, fn func(e E) K) map[K]E {
	m := make(map[K]E, len(s))
	for _, e := range s {
		m[fn(e)] = e
	}
	return m
}

// AssociateWith returns a map from the elements of the slice e as keys
// with the value retrieved by applying the given function fn.
func AssociateWith[K comparable, V any](s []K, fn func(key K) V) map[K]V {
	m := make(map[K]V, len(s))
	for _, k := range s {
		m[k] = fn(k)
	}
	return m
}

// Flatten returns a single slice of all elements from all slices in the given s.
func Flatten[E any](s [][]E) []E {
	n := SumOf(s, func(e []E) int { return len(e) })
	r := make([]E, n)
	i := 0
	for _, e := range s {
		for _, e := range e {
			r[i] = e
			i++
		}
	}
	return r
}

// Chunked returns a slice of slices, each with the size n containing the
// elements of the original slice e.
func Chunked[E any](s []E, n int) [][]E {
	c := len(s) / n
	if len(s)%n != 0 {
		c++
	}
	r := make([][]E, c)
	for i := 0; i < c; i++ {
		m := n
		if i == c-1 {
			m = len(s) - i*n
		}
		r[i] = make([]E, m)
		for j := range r[i] {
			r[i][j] = s[i*n+j]
		}
	}
	return r
}

// Unique returns the unique elements of a slice.
func Unique[E comparable](s []E) []E {
	r := make([]E, 0, len(s))
	for _, v := range s {
		if !Contains(r, v) {
			r = append(r, v)
		}
	}
	return r[:len(r):len(r)]
}

// UniqueInPlace returns the unique elements of a slice.
//
// It modifies the underlying array of slice e. Thus, this method should only
// be used if the passed slice e is not used afterwards!
func UniqueInPlace[E comparable](s []E) []E {
	n := 0
	for _, v := range s {
		if !Contains(s[:n], v) {
			s[n] = v
			n++
		}
	}
	return s[:n]
}

// UniqueBy returns a slice containing only elements from of slice e
// having unique keys returned by the given selector function fn.
func UniqueBy[E1 any, E2 comparable](s []E1, fn func(e E1) E2) []E1 {
	r := make([]E1, 0, len(s))
	k := make([]E2, 0, len(s))
	for _, v := range s {
		if key := fn(v); !Contains(k, key) {
			k = append(k, key)
			r = append(r, v)
		}
	}
	return r[:len(r):len(r)]
}

// UniqueByInPlace returns a slice containing only elements from of slice e
// having unique keys returned by the given selector function fn.
//
// It modifies the underlying array of slice e. Thus, this method should only
// be used if the passed slice e is not used afterwards!
func UniqueByInPlace[E1 any, E2 comparable](s []E1, fn func(e E1) E2) []E1 {
	n := 0
	k := make([]E2, 0, len(s))
	for _, v := range s {
		if key := fn(v); !Contains(k[:n], key) {
			k = append(k, key)
			s[n] = v
			n++
		}
	}
	return s[:n:n]
}

// Find returns the first element in the slice for which the
// function fn returns true or nil if no such element was found.
func Find[E any](s []E, fn func(e E) bool) (zeroValue E, _ error) {
	for _, e := range s {
		if fn(e) {
			return e, nil
		}
	}
	return zeroValue, errElementNotFound
}

// FindLast returns the last element in the slice for which the
// function fn returns true or nil if no such element was found.
func FindLast[E any](s []E, fn func(e E) bool) (zeroValue E, _ error) {
	for i := len(s) - 1; i >= 0; i-- {
		e := s[i]
		if fn(e) {
			return e, nil
		}
	}
	return zeroValue, errElementNotFound
}

// GroupBy
// Partition

// Intersect
// Distinct

// SumOf returns the sum of all values produced by applying the function fn
// to each element of the slice e.
func SumOf[E any, N number](s []E, fn func(e E) N) N {
	var n N
	for _, e := range s {
		n += fn(e)
	}
	return n
}

// MinOf
// MaxOf

// Reversed
// ReversedInPlace
