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
			want: 64,
			file: "input_test.txt",
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
