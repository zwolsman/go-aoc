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
			want: 20899048083289,
			file: "input_test.txt",
		},
		{
			name: "answer",
			want: 104831106565027,
			file: "input.txt",
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
		want int
		file string
	}{
		{
			name: "example",
			want: 273,
			file: "input_test.txt",
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

func Test_strip(t *testing.T) {
	tests := []struct {
		name string
		body []string
		want []string
	}{
		{
			name: "example",
			body: []string{
				"#...##.#..",
				"..#.#..#.#",
				".###....#.",
				"###.##.##.",
				".###.#####",
				".##.#....#",
				"#...######",
				".....#..##",
				"#.####...#",
				"#.##...##.",
			},
			want: []string{
				".#.#..#.",
				"###....#",
				"##.##.##",
				"###.####",
				"##.#....",
				"...#####",
				"....#..#",
				".####...",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strip(tt.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("strip() = %v, want %v", got, tt.want)
			}
		})
	}
}
