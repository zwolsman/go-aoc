package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {
	in, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	tests := []struct {
		name string
		in   []byte
		want any
	}{
		{
			name: "example",
			want: 13,
			in:   in,
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

func Test_part2(t *testing.T) {
	in, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	tests := []struct {
		name string
		in   []byte
		want any
	}{
		{
			name: "example",
			want: 30,
			in:   in,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
