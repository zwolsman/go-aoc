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
		want any
	}{
		{
			name: "example",
			file: "input_test.txt",
			want: 15,
		},
		{
			name: "answer",
			file: "input.txt",
			want: 12740,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in, _ := os.ReadFile(tt.file)
			if got := part1(in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name string
		file string
		want any
	}{
		{
			name: "example",
			file: "input_test.txt",
			want: 12,
		},
		{
			name: "answer",
			file: "input.txt",
			want: 11980,
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
