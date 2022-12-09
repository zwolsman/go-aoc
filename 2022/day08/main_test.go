package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {

	tests := []struct {
		name string
		file string
		want int
	}{
		{
			name: "example",
			file: "input_test.txt",
			want: 21,
		},
		{
			name: "example",
			file: "input_test_2.txt",
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in, _ := os.ReadFile(tt.file)
			if got := part1(in); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name string
		file string
		want int
	}{
		{
			name: "example",
			file: "input_test.txt",
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in, _ := os.ReadFile(tt.file)
			if got := part2(in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mask(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "simple",
			arr:  []int{0, 1, 1},
			want: []int{0, 1, 0},
		},
		{
			name: "example",
			arr:  []int{3, 0, 3, 7, 3},
			want: []int{0, 0, 0, 3, 0},
		},
		{
			name: "example",
			arr:  []int{3, 3, 5, 4, 9},
			want: []int{0, 0, 2, 0, 4},
		},
		{
			name: "blocked by same height",
			arr:  []int{3, 5, 3, 5, 3},
			want: []int{0, 1, 0, 2, 0},
		},
		{
			name: "immediate stop by same height",
			arr:  []int{2, 5, 5, 1, 2},
			want: []int{0, 1, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mask(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mask() = %v, want %v", got, tt.want)
			}
		})
	}
}
