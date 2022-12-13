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
	}{
		{
			name: "example",
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in, _ := os.ReadFile("input_test.txt")
			if got := part1(in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseList(t *testing.T) {
	tests := []struct {
		name  string
		line  string
		want  []any
		want1 int
	}{
		{
			name: "only digits",
			line: "[1,1,3,1,1]",
			want: []any{1, 1, 3, 1, 1},
		},
		{
			name: "1-deep nested list",
			line: "[[1],[2,3,4]]",
			want: []any{[]any{1}, []any{2, 3, 4}},
		},
		{
			name: "empty list",
			line: "[]",
			want: []any{},
		},
		{
			name: "nested empty list",
			line: "[[]]",
			want: []any{[]any{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := parsePacket(tt.line)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("parsePacket() got = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_verify(t *testing.T) {
	type args struct {
		a packet
		b packet
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "right order only ints",
			args: args{
				//[1,1,3,1,1] vs [1,1,5,1,1]
				packet{1, 1, 3, 1, 1},
				packet{1, 1, 5, 1, 1},
			},
			want: CORRECT,
		},
		{
			name: "right order lists, convert int to list",
			args: args{
				//[[1],[2,3,4]] vs [[1],4]
				packet{packet{1}, packet{2, 3, 4}},
				packet{packet{1}, 4},
			},
			want: CORRECT,
		},
		{
			name: "invalid order, int to list",
			args: args{
				//[9] vs [[8,7,6]]
				packet{9},
				packet{packet{8, 7, 6}},
			},
			want: INCORRECT,
		},
		{
			name: "right order, left ran out",
			args: args{
				// [[4,4],4,4] vs [[4,4],4,4,4]
				packet{packet{4, 4}, 4, 4},
				packet{packet{4, 4}, 4, 4, 4},
			},
			want: CORRECT,
		},
		{
			name: "right order, right not enough items",
			args: args{
				//[7,7,7,7] vs [7,7,7]
				packet{7, 7, 7, 7},
				packet{7, 7, 7},
			},
			want: INCORRECT,
		},
		{
			// [] vs [3]
			name: "left is empty, right has item",
			args: args{
				packet{},
				packet{3},
			},
			want: CORRECT,
		},
		{
			name: "left has more empty empty lists than right",
			args: args{
				//[[[]]] vs [[]]
				packet{packet{packet{}}},
				packet{packet{}},
			},
			want: INCORRECT,
		},
		{
			name: "bla",
			args: args{
				//[1,[2,[3,[4,[5,6,7]]]],8,9] vs [1,[2,[3,[4,[5,6,0]]]],8,9]
				packet{1, packet{2, packet{3, packet{4, packet{5, 6, 7}}}}, 8, 9},
				packet{1, packet{2, packet{3, packet{4, packet{5, 6, 0}}}}, 8, 9},
			},
			want: INCORRECT,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verify(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("verify() = %v, want %v", got, tt.want)
			}
		})
	}
}
