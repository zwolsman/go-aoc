package main

import (
	"os"
	"testing"
)

func Test_toDec(t *testing.T) {
	tests := []struct {
		name  string
		SNAFU string
		want  int
	}{
		{
			SNAFU: "1=-0-2",
			want:  1747,
		},
		{
			SNAFU: "12111",
			want:  906,
		},
		{
			SNAFU: "2=0=",
			want:  198,
		},
		{
			SNAFU: "21",
			want:  11,
		},
		{
			SNAFU: "2=01",
			want:  201,
		},
		{
			SNAFU: "111",
			want:  31,
		},
		{
			SNAFU: "20012",
			want:  1257,
		},
		{
			SNAFU: "112",
			want:  32,
		},
		{
			SNAFU: "1=-1=",
			want:  353,
		},
		{
			SNAFU: "1-12",
			want:  107,
		},
		{
			SNAFU: "12",
			want:  7,
		},
		{
			SNAFU: "1=",
			want:  3,
		},
		{
			SNAFU: "122",
			want:  37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toDec(tt.SNAFU); got != tt.want {
				t.Errorf("toDec() = %v, dec %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		want string
		file string
	}{
		{
			name: "example",
			want: "2=-1=0",
			file: "input_test.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in, _ := os.ReadFile(tt.file)
			if got := part1(in); got != tt.want {
				t.Errorf("part1() = %v, dec %v", got, tt.want)
			}
		})
	}
}

func Test_toSNAFU(t *testing.T) {
	tests := []struct {
		name string
		dec  int
		want string
	}{
		{
			want: "1=-0-2",
			dec:  1747,
		},
		{
			want: "12111",
			dec:  906,
		},
		{
			want: "2=0=",
			dec:  198,
		},
		{
			want: "21",
			dec:  11,
		},
		{
			want: "2=01",
			dec:  201,
		},
		{
			want: "111",
			dec:  31,
		},
		{
			want: "20012",
			dec:  1257,
		},
		{
			want: "112",
			dec:  32,
		},
		{
			want: "1=-1=",
			dec:  353,
		},
		{
			want: "1-12",
			dec:  107,
		},
		{
			want: "12",
			dec:  7,
		},
		{
			want: "1=",
			dec:  3,
		},
		{
			want: "122",
			dec:  37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toSNAFU(tt.dec); got != tt.want {
				t.Errorf("toSNAFU() = %v, SNAFU %v", got, tt.want)
			}
		})
	}
}
