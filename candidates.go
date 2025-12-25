package sudoku

type Candidates [CELLS]CandidateMask

type CandidateMask uint

func bit(d Cell) CandidateMask {
	return 1 << d
}

func (m CandidateMask) Has(d Cell) bool {
	return d != 0 && m&bit(d) != 0
}

func (m *CandidateMask) Remove(d Cell) {
	*m &^= bit(d)
}

func (m CandidateMask) IsSolved() bool {
	// exactly one bit set, and not zero
	return m != 0 && (m&(m-1)) == 0
}

func (m CandidateMask) Count() int {
	var count int
	for i := Cell(1); i <= SIZE; i++ {
		if m&bit(i) != 0 {
			count++
		}
	}
	return count
}
