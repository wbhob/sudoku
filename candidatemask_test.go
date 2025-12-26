package sudoku

import "testing"

func TestHas(t *testing.T) {
	type test struct {
		name string
		mask CandidateMask
		cell Cell
		want bool
	}

	tests := []test{
		{
			name: "zero has nothing",
			mask: 0,
			cell: 1,
			want: false,
		},
		{
			name: "zero bit ignored",
			mask: 1,
			cell: 1,
			want: false,
		},
		{
			name: "10 has 1",
			mask: 1 << 1,
			cell: 1,
			want: true,
		},
		{
			name: "100 does not have 1",
			mask: 1 << 2,
			cell: 1,
			want: false,
		},
		{
			name: "1<<9 has 9",
			mask: 1 << 9,
			cell: 9,
			want: true,
		},
		{
			name: "1<<10 has none",
			mask: 1 << 10,
			cell: 9,
			want: false,
		},
	}

	for _, tt := range tests {
		if got := tt.mask.Has(tt.cell); got != tt.want {
			t.Errorf("%s mismatch: got %t, want %t", tt.name, got, tt.want)
		}
	}
}
