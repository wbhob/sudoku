package sudoku

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wbhob/sudoku/testdata"
)

func Test_NewPuzzle(t *testing.T) {
	puzzle1 := NewPuzzle(testdata.Puzzle1)

	expect := Puzzle{
		0, 7, 5, 0, 9, 0, 3, 0, 8,
		0, 0, 0, 0, 0, 0, 1, 0, 0,
		0, 0, 9, 0, 0, 4, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 5, 0,
		0, 3, 0, 0, 4, 0, 9, 0, 7,
		2, 0, 0, 6, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 8, 0, 0, 1, 0,
		3, 0, 0, 0, 0, 2, 8, 0, 5,
		0, 5, 0, 0, 0, 0, 0, 4, 0,
	}

	if diff := cmp.Diff(puzzle1, expect); diff != "" {
		t.Errorf("NewPuzzle() mismatch (-want +got)\n%s", diff)
	}
}

func Test_PuzzleString(t *testing.T) {
	puzzle1 := Puzzle{
		0, 7, 5, 0, 9, 0, 3, 0, 8,
		0, 0, 0, 0, 0, 0, 1, 0, 0,
		0, 0, 9, 0, 0, 4, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 5, 0,
		0, 3, 0, 0, 4, 0, 9, 0, 7,
		2, 0, 0, 6, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 8, 0, 0, 1, 0,
		3, 0, 0, 0, 0, 2, 8, 0, 5,
		0, 5, 0, 0, 0, 0, 0, 4, 0,
	}

	expect := testdata.Puzzle1

	if diff := cmp.Diff(puzzle1.String(), expect); diff != "" {
		t.Errorf("Puzzle.String() mismatch (-want +got):\n%s", diff)
	}
}

func Test_Row(t *testing.T) {
	type test struct {
		name      string
		got, want [SIZE]Cell
	}

	puzzle1 := NewPuzzle(testdata.Puzzle1)

	tests := []test{
		{
			name: "Row(0)",
			got:  puzzle1.Row(0),
			want: [SIZE]Cell{0, 7, 5, 0, 9, 0, 3, 0, 8},
		},
		{
			name: "Row(8)",
			got:  puzzle1.Row(8),
			want: [SIZE]Cell{0, 5, 0, 0, 0, 0, 0, 4, 0},
		},
	}

	for _, tt := range tests {
		if diff := cmp.Diff(tt.got, tt.want); diff != "" {
			t.Errorf("Row() mismatch (-want +got)\n%s", diff)
		}
	}
}

func Test_Col(t *testing.T) {
	type test struct {
		name      string
		got, want [SIZE]Cell
	}

	puzzle1 := NewPuzzle(testdata.Puzzle1)

	tests := []test{
		{
			name: "Col(0)",
			got:  puzzle1.Col(0),
			want: [SIZE]Cell{0, 0, 0, 0, 0, 2, 0, 3, 0},
		},
		{
			name: "Col(8)",
			got:  puzzle1.Col(8),
			want: [SIZE]Cell{8, 0, 0, 0, 7, 0, 0, 5, 0},
		},
	}

	for _, tt := range tests {
		if diff := cmp.Diff(tt.got, tt.want); diff != "" {
			t.Errorf("%s mismatch (-want +got)\n%s", tt.name, diff)
		}
	}
}

func Test_Box(t *testing.T) {
	type test struct {
		name      string
		got, want [SIZE]Cell
	}

	puzzle1 := NewPuzzle(testdata.Puzzle1)

	tests := []test{
		{
			name: "Box(0)",
			got:  puzzle1.Box(0),
			want: [SIZE]Cell{
				0, 7, 5,
				0, 0, 0,
				0, 0, 9,
			},
		},
		{
			name: "Box(8)",
			got:  puzzle1.Box(8),
			want: [SIZE]Cell{
				0, 1, 0,
				8, 0, 5,
				0, 4, 0,
			},
		},
	}

	for _, tt := range tests {
		if diff := cmp.Diff(tt.want, tt.got); diff != "" {
			t.Errorf("%s mismatch (-want +got)\n%s", tt.name, diff)
		}
	}
}

func Test_RowOf(t *testing.T) {
	type test struct {
		name      string
		got, want int
	}

	tests := []test{
		{
			name: "RowOf(0)",
			got:  rowOf(0),
			want: 0,
		},
		{
			name: "RowOf(8)",
			got:  rowOf(8),
			want: 0,
		},
		{
			name: "RowOf(80)",
			got:  rowOf(80),
			want: 8,
		},
	}

	for _, tt := range tests {
		if diff := cmp.Diff(tt.want, tt.got); diff != "" {
			t.Errorf("%s mismatch (-want +got)\n%s", tt.name, diff)
		}
	}
}

func Test_ColOf(t *testing.T) {
	type test struct {
		name      string
		got, want int
	}

	tests := []test{
		{
			name: "ColOf(0)",
			got:  colOf(0),
			want: 0,
		},
		{
			name: "ColOf(8)",
			got:  colOf(8),
			want: 8,
		},
		{
			name: "ColOf(80)",
			got:  colOf(80),
			want: 8,
		},
	}

	for _, tt := range tests {
		if diff := cmp.Diff(tt.want, tt.got); diff != "" {
			t.Errorf("%s mismatch (-want +got)\n%s", tt.name, diff)
		}
	}
}

func Test_BoxOf(t *testing.T) {
	type test struct {
		name      string
		got, want int
	}

	tests := []test{
		{
			name: "BoxOf(0)",
			got:  boxOf(0),
			want: 0,
		},
		{
			name: "BoxOf(8)",
			got:  boxOf(8),
			want: 2,
		},
		{
			name: "BoxOf(80)",
			got:  boxOf(80),
			want: 8,
		},
	}

	for _, tt := range tests {
		if diff := cmp.Diff(tt.want, tt.got); diff != "" {
			t.Errorf("%s mismatch (-want +got)\n%s", tt.name, diff)
		}
	}
}

func Test_Valid(t *testing.T) {
	type test struct {
		name   string
		puzzle Puzzle
		want   bool
	}

	tests := []test{
		{
			name:   "puzzle1.Valid()",
			puzzle: NewPuzzle(testdata.Puzzle1),
			want:   true,
		},
		{
			name:   "solved puzzle.Valid()",
			puzzle: NewPuzzle(testdata.SolvedPuzzle),
			want:   true,
		},
		{
			name: "invalid row puzzle",
			puzzle: Puzzle{
				1, 0, 0, 1,
			},
			want: false,
		},
		{
			name: "invalid col puzzle",
			puzzle: Puzzle{
				1, 0, 0, 0, 0, 0, 0, 0, 0,
				1,
			},
			want: false,
		},
		{
			name: "invalid box puzzle",
			puzzle: Puzzle{
				1, 1,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		if diff := cmp.Diff(tt.want, tt.puzzle.Valid()); diff != "" {
			t.Errorf("%s mismatch (-want +got)\n%s", tt.name, diff)
		}
	}
}

func BenchmarkValid(b *testing.B) {
	puzzle1 := NewPuzzle(testdata.Puzzle1)

	for b.Loop() {
		puzzle1.Valid()
	}
}

func BenchmarkValidSolved(b *testing.B) {
	solvedPuzzle := NewPuzzle(testdata.SolvedPuzzle)

	for b.Loop() {
		solvedPuzzle.Valid()
	}
}
