package sudoku

func Solve(p Puzzle) (Puzzle, bool) {
	candidates := CandidatesFromPuzzle(p)

	state := State{puzzle: p, candidates: candidates}

	result, ok := dfs(state)
	return result.puzzle, ok
}

func dfs(s State) (State, bool) {
	guessIndex := chooseGuessIndex(&s)
	if guessIndex == -1 {
		if s.puzzle.IsSolved() {
			return s, true
		}
		return s, false
	}

	mask := s.candidates[guessIndex]
	for i := Cell(1); i <= SIZE; i++ {
		if mask.Has(i) {
			newState := s
			ok := newState.Apply(guessIndex, i)
			if !ok {
				continue
			}

			res, ok := dfs(newState)
			if ok {
				return res, true
			}
		}
	}

	return s, false
}

func chooseGuessIndex(s *State) int {
	minCount := SIZE + 1
	best := -1

	for i := range s.candidates {
		if s.puzzle[i] != 0 {
			continue
		}

		count := s.candidates[i].Count()
		if count == 0 {
			return -1
		}

		if count == 2 {
			return i
		}

		if count < minCount {
			best = i
			minCount = count
		}
	}

	return best
}
