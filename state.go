package sudoku

type State struct {
	puzzle     Puzzle
	candidates Candidates
}

func NewState(p Puzzle) State {
	candidates := CandidatesFromPuzzle(p)

	return State{
		puzzle:     p,
		candidates: candidates,
	}
}

type assignment struct {
	i int
	d Cell
}

func (s *State) Apply(i int, d Cell) bool {
	work := make([]assignment, 0, 20)
	work = append(work, assignment{i, d})

	for len(work) > 0 {
		// pop
		n := len(work) - 1
		todo := work[n]
		work = work[:n]

		if v := s.puzzle[todo.i]; v != 0 {
			if v != todo.d {
				return false
			}
			continue
		}

		if !s.Assign(todo.i, todo.d) {
			return false
		}

		for _, k := range peers[todo.i] {
			effect := s.candidates.Eliminate(k, todo.d)
			if effect&EffectContradiction != 0 {
				return false
			}
			if effect&EffectSolved != 0 {
				work = append(work, assignment{k, s.candidates[k].Solution()})
			}
		}
	}

	return true
}

func (s *State) Assign(i int, d Cell) bool {
	ok := s.candidates.Assign(i, d)
	if !ok {
		return false
	}

	s.puzzle.Set(i, d)
	return true
}
