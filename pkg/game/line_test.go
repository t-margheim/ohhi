package game

import (
	"testing"
)

func TestLine_isFinished(t *testing.T) {
	tests := []struct {
		name string
		l    *Line
		want bool
	}{
		{
			name: "four cells, all empty",
			l:    &Line{Empty, Empty, Empty, Empty},
			want: false,
		},
		{
			name: "four cells, one filled",
			l:    &Line{Empty, Empty, Blue, Empty},
			want: false,
		},
		{
			name: "four cells, two filled",
			l:    &Line{Empty, Empty, Blue, Blue},
			want: false,
		},
		{
			name: "four cells, three filled",
			l:    &Line{Empty, Red, Blue, Blue},
			want: false,
		},
		{
			name: "four cells, all filled invalid",
			l:    &Line{Blue, Red, Blue, Blue},
			want: false,
		},
		{
			name: "four cells, all filled solved",
			l:    &Line{Red, Red, Blue, Blue},
			want: true,
		},
		{
			name: "six cells, all filled solved",
			l:    &Line{Red, Red, Blue, Red, Blue, Blue},
			want: true,
		},
		{
			name: "six cells, all filled invalid",
			l:    &Line{Red, Red, Blue, Blue, Blue, Blue},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.isFinished(); got != tt.want {
				t.Errorf("Line.isFinished() = %v, want %v", got, tt.want)
			}
		})
	}
}
