# Stream

This project is a **generic** stream implementation. It provides a way to process data in a declarative way, leveraging functions as a mechanism for abstraction.

## Stream Methods

### `filter(predicate)`

The `filter` method is used to create a new stream that contains only the elements that satisfy the given predicate function.

```go
// Filter the stream to only include even numbers
filteredStream := stream.Of([]int{1, 2, 3, 4, 5}).Filter(func(num int) bool {
    return num%2 == 0
})

// Print the filtered stream
for _, num := range filteredStream {
    fmt.Println(num)
}
```

### `map(transform)`

The `map` method is used to create a new stream that contains the results of applying a transformation function to each element of the original stream.

```go
// Map the stream to uppercase strings
mappedStream := stream.Of([]string{"apple", "banana", "cherry"}).Map(strings, func(str string) string {
    return strings.ToUpper(str)
})
// Print the mapped stream
for _, str := range mappedStream {
    fmt.Println(str)
}
```

### `reduce(accumulator, initialValue)`

The `reduce` method is used to reduce the elements of a stream to a single value. It takes an accumulator function and an optional initial value. The accumulator function is applied to each element of the stream, along with the current accumulated value, and returns the updated accumulated value.

```go
// Reduce the stream to the sum of all numbers
sum := stream.Of([]int{1, 2, 3, 4, 5}).Reduce(func(acc, num int) int {
    return acc + num
}, 0)
// Print the sum
fmt.Println(sum)
```

### `forEach(action)`

The `forEach` method is used to perform an action on each element of the stream. It takes an action function as a parameter, which is applied to each element of the stream.

```go
// Print each element of the stream
stream.Of([]string{"apple", "banana", "cherry"}).ForEach(func(str string) {
    fmt.Println(str)
})
```

### Count

The `Count` method returns the length of the underlying slice of the stream

```go
stream.Of([]int{1,2,3}).Count()
```

### DistinctFunc

This method returns a new stream which is a copy of the primary one without duplicates

```go
stream.Of([]int{1,1,2,3}).DistinctFunc(func(v, w int) int {
    if v == w {
        return 0
    }
    if v < w {
        return -1
    }
    return 1
})
```

## Chaining several stream operations

```go
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
```

Run it:

```sh
❯ go run example/main.go
Result: main.foo{ID:2, accumulated:6, description:"BAR", data:[]int{1, 2, 3, 4, 5}}

ID: 1, accumulated: 1, description: foo, data: []int{1, 2, 3, 4, 5}
ID: 2, accumulated: 2, description: bar, data: []int{1, 2, 3, 4, 5}
ID: 3, accumulated: 3, description: baz, data: []int{1, 2, 3, 4, 5}
ID: 4, accumulated: 4, description: qux, data: []int{1, 2, 3, 4, 5}
ID: 5, accumulated: 5, description: zoo, data: []int{1, 2, 3, 4, 5}
ID: 6, accumulated: 5, description: zoo, data: []int{1, 2, 3, 4, 5}
```
