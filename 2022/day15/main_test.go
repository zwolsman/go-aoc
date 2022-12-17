package main

import (
	"os"
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
