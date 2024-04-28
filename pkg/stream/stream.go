// Package stream contains a generic Stream type that can be used to filter, map and reduce slices of any type
package stream

// Stream is a generic type that holds a slice of any type
type Stream[T any] struct {
	stream []T
}

// Of creates a new Stream with the given slice
func Of[T any](stream []T) Stream[T] {
	return Stream[T]{stream: stream}
}

// Filter returns a new Stream with the elements that satisfy the given predicate
func (s Stream[T]) Filter(f func(T) bool) Stream[T] {
	var result []T
	for _, v := range s.stream {
		if f(v) {
			result = append(result, v)
		}
	}
	return Of(result)
}

// Map returns a new Stream with the elements that are the result of applying the given function
func (s Stream[T]) Map(f func(T) T) Stream[T] {
	var result []T
	for _, v := range s.stream {
		result = append(result, f(v))
	}
	return Of(result)
}

// Reduce returns a single value that is the result of applying the given function to all elements
func (s Stream[T]) Reduce(f func(T, T) T) T {
	result := s.stream[0]
	for _, v := range s.stream[1:] {
		result = f(result, v)
	}
	return result
}

// ForEach applies the given function to all elements
func (s Stream[T]) ForEach(f func(T)) {
	for _, v := range s.stream {
		f(v)
	}
}

// ToSlice returns the slice of the Stream
func (s Stream[T]) ToSlice() []T {
	return s.stream
}
