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
			want: 142,
			data: "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet",
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
	type args struct {
		in []byte
	}
	tests := []struct {
		name string
		want any
		data string
	}{
		{
			name: "example",
			want: 281,
			data: "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen",
		}, {
			name: "example",
			want: 18,
			data: "oneight",
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
