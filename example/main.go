package main

import (
	"fmt"
	"strings"

	"theskyinflames/stream/pkg/stream"
)

type foo struct {
	ID          int
	accumulated int
	description string
	data        []int
}

func (f foo) String() string {
	return fmt.Sprintf("accumulated: %d, description: %s", f.accumulated, f.description)
}

func main() {
	s := stream.Of([]foo{
		{ID: 1, accumulated: 1, description: "foo", data: []int{1, 2, 3, 4, 5}},
		{ID: 2, accumulated: 2, description: "bar", data: []int{1, 2, 3, 4, 5}},
		{ID: 3, accumulated: 3, description: "baz", data: []int{1, 2, 3, 4, 5}},
		{ID: 4, accumulated: 4, description: "qux", data: []int{1, 2, 3, 4, 5}},
		{ID: 5, accumulated: 5, description: "zoo", data: []int{1, 2, 3, 4, 5}},
		{ID: 6, accumulated: 5, description: "zoo", data: []int{1, 2, 3, 4, 5}},
	})

	result := s.DistinctFunc(func(f, g foo) int {
		return strings.Compare(f.String(), g.String())
	}).Map(func(f foo) foo {
		f.description = strings.ToUpper(f.description)
		return f
	}).Filter(func(f foo) bool {
		return f.ID%2 == 0
	}).Reduce(func(f1, f2 foo) foo {
		f1.accumulated += f2.accumulated
		return f1
	})

	fmt.Printf("Result: %#v\n\n", result)

	// Primary stream is not modified
	s.ForEach(func(f foo) {
		fmt.Printf("ID: %d, accumulated: %d, description: %s, data: %#v\n", f.ID, f.accumulated, f.description, f.data)
	})
}
