package sudoku

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wbhob/sudoku/testdata"
)

func TestCandidatesFromPuzzle(t *testing.T) {
	type test struct {
		name      string
		want, got Candidates
	}

	tests := []test{
		{
			name: "solved puzzle has no candidates",
			got:  CandidatesFromPuzzle(NewPuzzle(testdata.SolvedPuzzle)),
			want: Candidates{
				1 << 6, 1 << 5, 1 << 9, 1 << 3, 1 << 1, 1 << 4, 1 << 2, 1 << 8, 1 << 7,
				1 << 1, 1 << 8, 1 << 7, 1 << 6, 1 << 5, 1 << 2, 1 << 4, 1 << 3, 1 << 9,
				1 << 2, 1 << 3, 1 << 4, 1 << 8, 1 << 9, 1 << 7, 1 << 5, 1 << 1, 1 << 6,
				1 << 4, 1 << 2, 1 << 6, 1 << 1, 1 << 3, 1 << 5, 1 << 9, 1 << 7, 1 << 8,
				1 << 8, 1 << 7, 1 << 1, 1 << 9, 1 << 4, 1 << 6, 1 << 3, 1 << 5, 1 << 2,
				1 << 5, 1 << 9, 1 << 3, 1 << 2, 1 << 7, 1 << 8, 1 << 6, 1 << 4, 1 << 1,
				1 << 3, 1 << 1, 1 << 2, 1 << 5, 1 << 8, 1 << 9, 1 << 7, 1 << 6, 1 << 4,
				1 << 7, 1 << 6, 1 << 5, 1 << 4, 1 << 2, 1 << 1, 1 << 8, 1 << 9, 1 << 3,
				1 << 9, 1 << 4, 1 << 8, 1 << 7, 1 << 6, 1 << 3, 1 << 1, 1 << 2, 1 << 5,
			},
		},
		{
			name: "almost solved puzzle has candidates",
			got:  CandidatesFromPuzzle(NewPuzzle(testdata.AlmostSolvedPuzzle)),
			want: Candidates{
				1 << 6, 1 << 5, 1 << 9, 1 << 3, 1 << 1, 1 << 4, (1<<2 | 1<<6), 1 << 8, 1 << 7,
				1 << 1, 1 << 8, 1 << 7, 1 << 6, 1 << 5, 1 << 2, 1 << 4, 1 << 3, 1 << 9,
				1 << 2, 1 << 3, 1 << 4, 1 << 8, 1 << 9, 1 << 7, 1 << 5, 1 << 1, 1 << 6,
				1 << 4, 1 << 2, 1 << 6, 1 << 1, 1 << 3, 1 << 5, 1 << 9, 1 << 7, 1 << 8,
				1 << 8, 1 << 7, 1 << 1, 1 << 9, 1 << 4, 1 << 6, 1 << 3, 1 << 5, 1 << 2,
				1 << 5, 1 << 9, 1 << 3, 1 << 2, 1 << 7, 1 << 8, 1 << 6, 1 << 4, 1 << 1,
				1 << 3, 1 << 1, 1 << 2, 1 << 5, 1 << 8, 1 << 9, 1 << 7, 1 << 6, 1 << 4,
				1 << 7, 1 << 6, 1 << 5, 1 << 4, 1 << 2, 1 << 1, 1 << 8, 1 << 9, 1 << 3,
				1 << 9, 1 << 4, 1 << 8, 1 << 7, 1 << 6, 1 << 3, 1 << 1, 1 << 2, 1 << 5,
			},
		},
	}

	for _, tt := range tests {
		if diff := cmp.Diff(tt.want, tt.got); diff != "" {
			t.Errorf("%s mismatch (-want +got)\n%s", tt.name, diff)
		}
	}
}

func BenchmarkCandidatesFromPuzzle(b *testing.B) {
	puzzle := NewPuzzle(testdata.Puzzle1)
	for b.Loop() {
		_ = CandidatesFromPuzzle(puzzle)
	}
}

func BenchmarkCandidatesFromPuzzleAlmostSolved(b *testing.B) {
	puzzle := NewPuzzle(testdata.AlmostSolvedPuzzle)
	for b.Loop() {
		_ = CandidatesFromPuzzle(puzzle)
	}
}

func BenchmarkCandidatesFromPuzzleSolved(b *testing.B) {
	puzzle := NewPuzzle(testdata.SolvedPuzzle)
	for b.Loop() {
		_ = CandidatesFromPuzzle(puzzle)
	}
}
