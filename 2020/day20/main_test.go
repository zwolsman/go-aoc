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

func Test_edges(t *testing.T) {
	tests := []struct {
		name string
		body []string
		want []map[location]string
	}{
		{
			name: "simple",
			body: []string{
				"123",
				"456",
				"789",
			},
			want: []map[location]string{
				{
					TOP:    "123",
					BOTTOM: "789",
					LEFT:   "147",
					RIGHT:  "369",
				},
				{
					TOP:    "741",
					BOTTOM: "963",
					LEFT:   "789",
					RIGHT:  "123",
				},
				{
					TOP:    "987",
					BOTTOM: "321",
					LEFT:   "963",
					RIGHT:  "741",
				},
				{
					TOP:    "369",
					BOTTOM: "147",
					LEFT:   "321",
					RIGHT:  "987",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := edges(tt.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("edges() = %v, want %v", got, tt.want)
			}
		})
	}
}
