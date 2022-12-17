package common

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
