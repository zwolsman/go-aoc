package main

import (
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {
	var (
		m0 = monkey{
			items: []int{79, 98},
			operation: func(i item) item {
				return i * 19
			},
			test: 23,
		}
		m1 = monkey{
			operation: func(i item) item {
				return i + 6
			},
			items: []int{54, 65, 75, 74},
			test:  19,
		}
		m2 = monkey{
			items: []int{79, 60, 97},
			operation: func(i item) item {
				return i * i
			},
			test: 13,
		}

		m3 = monkey{
			items: []int{74},
			operation: func(i item) item {
				return i + 3
			},
			test: 17,
		}
	)

	m0.chain = [2]*monkey{
		&m2,
		&m3,
	}
	m1.chain = [2]*monkey{
		&m2,
		&m0,
	}
	m2.chain = [2]*monkey{
		&m1,
		&m3,
	}
	m3.chain = [2]*monkey{
		&m0,
		&m1,
	}

	tests := []struct {
		name string
		in   []*monkey
		want int
	}{
		{
			name: "example",
			in: []*monkey{
				&m0, &m1, &m2, &m3,
			},
			want: 10605,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
