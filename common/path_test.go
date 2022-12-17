package common

import "testing"

func TestPath_Intersects(t *testing.T) {
	tests := []struct {
		name  string
		path  Path
		other Vector
		want  bool
	}{
		{
			name: "point is in line",
			path: Path{
				[][2]Vector{
					{
						{0, 0},
						{10, 0},
					},
				},
			},
			other: Vector{5, 0},
			want:  true,
		},
		{
			name: "point not in line",
			path: Path{
				[][2]Vector{
					{
						{0, 0},
						{10, 0},
					},
				},
			},
			other: Vector{20, 5},
			want:  false,
		},
		{
			name: "point is origin",
			path: Path{
				[][2]Vector{
					{
						{0, 0},
						{10, 0},
					},
				},
			},
			other: Vector{0, 0},
			want:  true,
		},
		{
			name: "example",
			path: Path{
				[][2]Vector{
					{
						{503, 4},
						{502, 4},
					},
					{
						{502, 4},
						{502, 9},
					},
					{
						{502, 9},
						{494, 9},
					},
				},
			},
			other: Vector{500, 1},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.path.Intersects(tt.other); got != tt.want {
				t.Errorf("Intersects() = %v, want %v", got, tt.want)
			}
		})
	}
}
