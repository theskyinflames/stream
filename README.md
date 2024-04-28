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

## Combine several stream operations

```go
pacage main

import (
    "fmt"
    "theskyinflames/stream/pkg/stream"
)

func main() {
    s := stream.Of([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

    result := s.Filter(func(v int) bool {
        return v%2 == 0
    }).Map(func(v int) int {
        return v * 2
    }).Reduce(func(acc, v int) int {
        return acc + v
    })

    fmt.Printf("Result: %d\n", result)

    s.ForEach(func(v int) {
        println(v)
    })

    fmt.Printf("Result: %d\n", s.Reduce(func(acc, v int) int {
        return acc + v
    }))
}

```
