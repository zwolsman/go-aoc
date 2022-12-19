package common

import (
	"reflect"
	"testing"
)

type testString string

func (s testString) Key() string { return string(s) }
func TestCombinations(t *testing.T) {
	tests := []struct {
		name  string
		input []testString
		want  [][]testString
	}{
		{
			name:  "abc",
			input: []testString{"a", "b", "c"},
			want: [][]testString{
				{"c", "b", "a"},
				{"b", "a"},
				{"c", "a"},
				{"a"},
				{"c", "b"},
				{"b"},
				{"c"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Combinations(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Combinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
