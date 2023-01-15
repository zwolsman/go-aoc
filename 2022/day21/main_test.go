package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		want int
		file string
	}{
		{
			name: "example",
			want: 152,
			file: "input_test.txt",
		},
		{
			name: "answer",
			want: 194058098264286,
			file: "input.txt",
		},
	}
	for _, tt := range tests {
		in, _ := os.ReadFile(tt.file)
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name string
		want int
		file string
	}{
		{
			name: "example",
			want: 301,
			file: "input_test.txt",
		},
		{
			name: "answer",
			want: 3592056845086,
			file: "input.txt",
		},
	}
	for _, tt := range tests {
		in, _ := os.ReadFile(tt.file)
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
