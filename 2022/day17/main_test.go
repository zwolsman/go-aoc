package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {
	in, _ := os.ReadFile("input_test.txt")
	tests := []struct {
		name  string
		rocks int
		want  int
	}{
		{
			name:  "example 2022 rounds",
			rocks: 2022,
			want:  3068,
		},
		{
			name:  "example 10 rounds",
			rocks: 10,
			want:  17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(in, tt.rocks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
