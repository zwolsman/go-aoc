package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {
	in, _ := os.ReadFile("input_test.txt")
	tests := []struct {
		name string
		y    int
		want int
	}{
		{
			name: "example",
			y:    10,
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(in, tt.y); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	in, _ := os.ReadFile("input_test.txt")
	tests := []struct {
		name string
		size int
		want int
	}{
		{
			name: "example",
			size: 20,
			want: 56000011,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(in, tt.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
