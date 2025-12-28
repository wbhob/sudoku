package sudoku

import (
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wbhob/sudoku/testdata"
)

func TestSolve(t *testing.T) {
	type test struct {
		name   string
		puzzle Puzzle
		want   Puzzle
	}

	tests := []test{
		{
			name:   "Puzzle 1",
			puzzle: NewPuzzle(testdata.Puzzle1),
			want:   NewPuzzle(testdata.Puzzle1Solution),
		},
		{
			name:   "Puzzle 1",
			puzzle: NewPuzzle(testdata.AlmostSolvedPuzzle),
			want:   NewPuzzle(testdata.SolvedPuzzle),
		},
	}

	for _, tt := range tests {
		soln, ok := Solve(tt.puzzle)
		if !ok {
			t.Errorf("Solve() failed, want %t, got %t", true, ok)
		}
		if diff := cmp.Diff(tt.want, soln); diff != "" {
			t.Errorf("%s mismatch (-want +got)\n%s", tt.name, diff)
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	puzzle1 := NewPuzzle(testdata.Puzzle1)

	for b.Loop() {
		_, _ = Solve(puzzle1)
	}
}

func BenchmarkSolveMany(b *testing.B) {
	const count = 2048
	const clues = 23

	rng := rand.New(rand.NewSource(rand.Int63()))

	puzzles := make([]Puzzle, count)
	for i := range count {
		puzzles[i] = generateQuick(rng, clues)
	}

	b.ResetTimer()

	for b.Loop() {
		puzzle := puzzles[rng.Int()%2048]
		_, _ = Solve(puzzle)
	}
}
