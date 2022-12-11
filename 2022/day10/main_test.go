package main

import (
	"os"
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
			want: 13140,
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