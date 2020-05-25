package game

import "testing"

func TestBoard_isSolved(t *testing.T) {
	type fields struct {
		Rows    []Line
		Columns []Line
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "4x4 unsolved",
			fields: fields{
				Rows: []Line{
					{Empty, Red, Blue, Red},
					{Red, Red, Blue, Blue},
					{Red, Blue, Red, Blue},
					{Blue, Blue, Red, Red},
				},
				Columns: []Line{
					{Empty, Red, Red, Blue},
					{Red, Red, Blue, Blue},
					{Blue, Blue, Red, Red},
					{Red, Blue, Blue, Red},
				},
			},
			want: false,
		},
		{
			name: "4x4 solved",
			fields: fields{
				Rows: []Line{
					{Blue, Red, Blue, Red},
					{Red, Red, Blue, Blue},
					{Red, Blue, Red, Blue},
					{Blue, Blue, Red, Red},
				},
				Columns: []Line{
					{Blue, Red, Red, Blue},
					{Red, Red, Blue, Blue},
					{Blue, Blue, Red, Red},
					{Red, Blue, Blue, Red},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Rows:    tt.fields.Rows,
				Columns: tt.fields.Columns,
			}
			if got := b.isSolved(); got != tt.want {
				t.Errorf("Board.isSolved() = %v, want %v", got, tt.want)
			}
		})
	}
}
