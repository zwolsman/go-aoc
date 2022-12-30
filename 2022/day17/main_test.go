package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
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
		{
			name:  "example 1000000000000 rounds",
			rocks: 1000000000000,
			want:  1514285714288,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(in, tt.rocks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
