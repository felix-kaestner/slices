package slices

import "errors"

var errEmptySlice = errors.New("slices: empty slice")

// Filter executes the function fn to each element of the slice s
// returning a newly allocated slice of all elements for which the
// function fn returns true.
func Filter[E any](s []E, fn func(e E) bool) []E {
	r := make([]E, 0, len(s))
	for _, e := range s {
		if fn(e) {
			r = append(r, e)
		}
	}
	return r
}

// Filter executes the function fn to each element of the slice s
// returning a slice of all elements for which the function fn returns true.
//
// It modifies the underlying array of slice s. Thus, this method should only
// be used if the passed slice s is not used afterwards!
func FilterInPlace[E any](s []E, fn func(e E) bool) []E {
	n := 0
	for _, e := range s {
		if fn(e) {
			s[n] = e
			n++
		}
	}
	return s[:n]
}

// Map applies the function fn to each element of the slice s.
// It returns a newly allocated slice with same length as s where
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
// elements of the slice s.
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
