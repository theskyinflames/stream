package stream_test

import (
	"testing"

	"theskyinflames/stream/pkg/stream"

	"github.com/stretchr/testify/assert"
)

var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestStream_Filter(t *testing.T) {
	result := stream.Of(slice).Filter(func(v int) bool {
		return v%2 == 0
	})

	assert.ElementsMatch(t, []int{2, 4, 6, 8, 10}, result.ToSlice())
}

func TestStream_Map(t *testing.T) {
	result := stream.Of(slice).Map(func(v int) int {
		return v * 2
	})

	assert.ElementsMatch(t, []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, result.ToSlice())
}

func TestStream_Reduce(t *testing.T) {
	result := stream.Of(slice).Reduce(func(acc, v int) int {
		return acc + v
	})

	assert.Equal(t, 55, result)
}

func TestStream_ForEach(t *testing.T) {
	var result []int
	stream.Of(slice).ForEach(func(v int) {
		result = append(result, v)
	})

	assert.ElementsMatch(t, slice, result)
}

func TestStream_ToSlice(t *testing.T) {
	result := stream.Of(slice).ToSlice()

	assert.ElementsMatch(t, slice, result)
}

func TestStream_Count(t *testing.T) {
	result := stream.Of(slice).Count()

	assert.Equal(t, 10, result)
}

func TestStream_DistinctFunc(t *testing.T) {
	result := stream.Of(slice).DistinctFunc(func(v, w int) int {
		if v == w {
			return 0
		}
		if v < w {
			return -1
		}
		return 1
	})

	assert.ElementsMatch(t, slice, result.ToSlice())
}
