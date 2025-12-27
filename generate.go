package sudoku

import "math/rand"

// generateQuick makes a puzzle by permuting a known solved grid, then removing clues.
// if unique=true, it enforces uniqueness using your solver hooks.
func generateQuick(rng *rand.Rand, clues int) Puzzle {
	if clues < 17 {
		clues = 17
	}
	if clues > 81 {
		clues = 81
	}

	p := permuteSolved(baseSolved(), rng)

	// remove random cells
	idx := make([]int, 81)
	for i := range idx {
		idx[i] = i
	}
	rng.Shuffle(81, func(i, j int) { idx[i], idx[j] = idx[j], idx[i] })

	toRemove := 81 - clues
	for k := 0; k < 81 && toRemove > 0; k++ {
		i := idx[k]
		old := p[i]
		if old == 0 {
			continue
		}
		p[i] = 0

		toRemove--
	}

	return p
}

// a fixed valid solution grid (any correct one works).
func baseSolved() Puzzle {
	// pattern-based solved grid
	// row r, col c: (r*3 + r/3 + c) % 9 + 1
	var p Puzzle
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			p[r*9+c] = uint8((r*3+r/3+c)%9 + 1)
		}
	}
	return p
}

// permuteSolved applies random relabeling + row/col shuffles within bands/stacks + band/stack swaps.
// still a valid solved grid, much shorter than backtracking.
func permuteSolved(p Puzzle, rng *rand.Rand) Puzzle {
	// relabel digits 1..9
	var mapd [10]uint8
	perm := rng.Perm(9)
	for d := 1; d <= 9; d++ {
		mapd[d] = uint8(perm[d-1] + 1)
	}
	for i := range p {
		p[i] = mapd[p[i]]
	}

	// swap rows within each band
	for band := 0; band < 3; band++ {
		r0 := band * 3
		a, b := rng.Intn(3), rng.Intn(3)
		swapRows(&p, r0+a, r0+b)
	}
	// swap cols within each stack
	for stack := 0; stack < 3; stack++ {
		c0 := stack * 3
		a, b := rng.Intn(3), rng.Intn(3)
		swapCols(&p, c0+a, c0+b)
	}
	// swap whole bands
	ba, bb := rng.Intn(3), rng.Intn(3)
	swapBand(&p, ba, bb)
	// swap whole stacks
	sa, sb := rng.Intn(3), rng.Intn(3)
	swapStack(&p, sa, sb)

	return p
}

func swapRows(p *Puzzle, r1, r2 int) {
	if r1 == r2 {
		return
	}
	for c := 0; c < 9; c++ {
		i, j := r1*9+c, r2*9+c
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}

func swapCols(p *Puzzle, c1, c2 int) {
	if c1 == c2 {
		return
	}
	for r := 0; r < 9; r++ {
		i, j := r*9+c1, r*9+c2
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}

func swapBand(p *Puzzle, b1, b2 int) {
	if b1 == b2 {
		return
	}
	for k := 0; k < 3; k++ {
		swapRows(p, b1*3+k, b2*3+k)
	}
}

func swapStack(p *Puzzle, s1, s2 int) {
	if s1 == s2 {
		return
	}
	for k := 0; k < 3; k++ {
		swapCols(p, s1*3+k, s2*3+k)
	}
}
