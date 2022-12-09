package common

import (
	"reflect"
	"testing"
)

func TestVector_Dir(t *testing.T) {
	tests := []struct {
		name string
		a, b Vector
		want Vector
	}{
		{
			name: "right",
			a:    Vector{0, 0},
			b:    Vector{2, 0},
			want: Vector{1, 0},
		},
		{
			name: "left",
			a:    Vector{0, 0},
			b:    Vector{-2, 0},
			want: Vector{-1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Dir(tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dir() = %v, want %v", got, tt.want)
			}
		})
	}
}
