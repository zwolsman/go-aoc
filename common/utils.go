package common

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

var PLACEHOLDER = struct{}{}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func MaxBy[T any, R constraints.Ordered](in []T, selector func(T) R) R {
	var max R
	for _, item := range in {
		max = Max(max, selector(item))
	}
	return max
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}
func MinArr[T constraints.Ordered](arr []T) T {
	if len(arr) == 0 {
		var defaultValue T
		return defaultValue
	}

	min := arr[0]
	for i := 1; i < len(arr); i++ {
		min = Min(min, arr[i])
	}

	return min
}

func MaxArr[T constraints.Ordered](arr []T) T {
	if len(arr) == 0 {
		var defaultValue T
		return defaultValue
	}

	max := arr[0]
	for i := 1; i < len(arr); i++ {
		max = Max(max, arr[i])
	}

	return max
}

func MinBy[T any, R constraints.Ordered](in []T, selector func(T) R) R {
	var min R
	for _, item := range in {
		min = Min(min, selector(item))
	}
	return min
}

func Keys[T map[K]V, K comparable, V any](src T) []K {
	out := make([]K, len(src))
	index := 0

	for k, _ := range src {
		out[index] = k
		index++
	}

	return out
}
func Values[T map[K]V, K comparable, V any](src T) []V {
	out := make([]V, len(src))
	index := 0
	for _, v := range src {
		out[index] = v
		index++
	}
	return out
}

func Copy[T map[K]V, K comparable, V any](src T) T {
	out := make(T)
	for k, v := range src {
		out[k] = v
	}
	return out
}

func Map[T any, O any](in []T, mapFunc func(T) O) []O {
	out := make([]O, len(in))
	for i, item := range in {
		out[i] = mapFunc(item)
	}

	return out
}

func Equals[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, x := range a {
		if b[i] != x {
			return false
		}
	}

	return true
}

func Apply[T any](fn func(T, T) T, items []T, arg T) []T {
	result := make([]T, len(items))
	for i, v := range items {
		result[i] = fn(v, arg)
	}
	return result
}

func Filter[T any](collection []T, predicate func(T) bool) []T {
	var out []T
	for _, c := range collection {
		if predicate(c) {
			out = append(out, c)
		}
	}

	return out
}

type Keyable interface {
	Key() string
}

func Combinations[T any](input []T) [][]T {
	return combinations(input, make(map[string][][]T))
}

func Index[T comparable](haystack []T, needle T) (int, bool) {
	for i, v := range haystack {
		if v == needle {
			return i, true
		}
	}

	return 0, false
}

func combinations[T any](input []T, cache map[string][][]T) [][]T {
	if len(input) == 0 {
		return make([][]T, 0, 0)
	}

	if len(input) == 1 {
		var dst []T
		copy(dst, input)
		return [][]T{
			dst,
		}
	}

	key := ""
	for _, v := range input {
		key += fmt.Sprintf("%v", v)
	}

	if v, ok := cache[key]; ok {
		return v
	}

	var result [][]T

	for i := 1; i <= len(input); i++ {
		current := input[i-1]
		for _, c := range combinations(input[i:], cache) {
			result = append(result, append(c, current))
		}
		result = append(result, []T{current})
	}

	cache[key] = result
	return result

}
