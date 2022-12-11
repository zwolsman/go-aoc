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
		{
			name: "answer",
			file: "input.txt",
			want: 13440,
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

func Test_part2(t *testing.T) {
	tests := []struct {
		name string
		file string
		want string
	}{
		{
			name: "example",
			file: "input_test.txt",
			want: `
##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in, _ := os.ReadFile(tt.file)
			if got := part2(in); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
