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

func (c *Candidates) Assign(index int, d Cell) {
	row := rowOf(index)
	col := colOf(index)
	box := boxOf(index)

	// index in row = col
	// index in col = row
	bi := (row%BOX)*BOX + (col % BOX) // 0..8

	for i := range SIZE {
		rindex := rowIndex(row, i)
		cindex := colIndex(col, i)
		bindex := boxIndex(box, i)

		if i != col && !c[rindex].IsSolved() {
			c.Eliminate(rindex, d)
		}

		if i != row && !c[cindex].IsSolved() {
			c.Eliminate(cindex, d)
		}

		if i != bi && !c[bindex].IsSolved() {
			c.Eliminate(bindex, d)
		}
	}
}

func (c *Candidates) Eliminate(index int, d Cell) {
	c[index].Remove(d)
}
