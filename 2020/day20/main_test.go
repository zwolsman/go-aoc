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
		want int
	}{
		{
			name: "example",
			want: 20899048083289,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
