package common

import "testing"

func TestVector_Span(t *testing.T) {
	type args struct {
		radius int
		y      int
	}
	tests := []struct {
		name string
		v    Vector
		args args
		want int
	}{
		{
			name: "y = -r",
			v:    Vector{5, 5},
			args: args{
				radius: 5,
				y:      0,
			},
			want: 1,
		},
		{
			name: "y = r-1",
			v:    Vector{5, 5},
			args: args{
				radius: 5,
				y:      1,
			},
			want: 3,
		},
		{
			name: "y = r+1",
			v:    Vector{5, 5},
			args: args{
				radius: 5,
				y:      9,
			},
			want: 3,
		},
		{
			name: "y = +r",
			v:    Vector{5, 5},
			args: args{
				radius: 5,
				y:      10,
			},
			want: 1,
		},
		{
			name: "y == r",
			v:    Vector{5, 5},
			args: args{
				radius: 5,
				y:      5,
			},
			want: 11,
		},
		{
			name: "y = r-2",
			v:    Vector{5, 5},
			args: args{
				radius: 5,
				y:      3,
			},
			want: 7,
		},
		{
			name: "out of bounds",
			v:    Vector{5, 5},
			args: args{
				radius: 5,
				y:      11,
			},
			want: 0,
		},
		{
			name: "bla",
			v:    Vector{20, 1},
			args: args{
				radius: 7,
				y:      0,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := tt.v.Span(tt.args.radius, tt.args.y); got != tt.want {
				t.Errorf("Span() = %v, want %v", got, tt.want)
			}
		})
	}
}
