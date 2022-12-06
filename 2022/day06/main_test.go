package main

import (
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		want any
		data string
	}{
		{
			name: "example",
			want: 5,
			data: "bvwbjplbgvbhsrlpgdmjqwftvncz",
		},
		{
			name: "example",
			want: 6,
			data: "nppdvjthqldpwncqszvftbrmjlhg",
		},
		{
			name: "example",
			want: 10,
			data: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		},
		{
			name: "example",
			want: 11,
			data: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1([]byte(tt.data)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name string
		want any
		data string
	}{
		{
			name: "exaple",
			want: 19,
			data: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		},
		{
			name: "exaple",
			want: 23,
			data: "bvwbjplbgvbhsrlpgdmjqwftvncz",
		},
		{
			name: "exaple",
			want: 23,
			data: "nppdvjthqldpwncqszvftbrmjlhg",
		},
		{
			name: "exaple",
			want: 29,
			data: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		},
		{
			name: "exaple",
			want: 26,
			data: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2([]byte(tt.data)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
