package sudoku

type Candidates [CELLS]CandidateMask

func CandidatesFromPuzzle(p Puzzle) Candidates {
	var rowUseds, colUseds, boxUseds [SIZE]CandidateMask

	for i := range SIZE {
		row := p.Row(i)
		col := p.Col(i)
		box := p.Box(i)
		for j := range SIZE {
			rowUseds[i] |= 1 << row[j] &^ 1
			colUseds[i] |= 1 << col[j] &^ 1
			boxUseds[i] |= 1 << box[j] &^ 1
		}
	}

	var candidates Candidates
	for cell := range CELLS {
		if p[cell] != 0 {
			candidates[cell] = 1 << p[cell]
			continue
		}

		r := rowUseds[rowOf(cell)]
		c := colUseds[colOf(cell)]
		b := boxUseds[boxOf(cell)]

		used := r | c | b

		candidates[cell] = allCandidates &^ used
	}

	return candidates
}

// Assign assigns the cell at index i to digit d. Returns true if the set was valid,
// false if the cell can't be assigned to d.
func (c *Candidates) Assign(i int, d Cell) bool {
	if d == 0 || d > SIZE || c[i] == 0 {
		return false
	}

	mask := bit(d)

	// if already set, no op
	if c[i] == mask {
		return true
	}

	if c[i]&mask == 0 {
		return false
	}

	c[i] = mask
	return true
}

func (c *Candidates) Eliminate(index int, d Cell) Effect {
	old := c[index]
	c[index].Remove(d)

	if c[index] == old {
		return EffectNone
	}

	if c[index].IsSolved() {
		return EffectChanged | EffectSolved
	}

	if c[index].Count() == 0 {
		return EffectChanged | EffectContradiction
	}

	return EffectChanged
}
