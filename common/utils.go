package common

import "fmt"

var PLACEHOLDER = struct{}{}

func Max[T int](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min[T int](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
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
