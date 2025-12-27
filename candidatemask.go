package sudoku

import "math/bits"

type CandidateMask uint16

// valid bits are 1..SIZE. 0 bit is unused and over SIZE is unused.
// roughly speaking. 2^SIZE will be 10000000
// -1 yields 0111111111
// -1 again yields 01111111110
const allCandidates = (CandidateMask(1) << (SIZE + 1)) - 2

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
	return bits.OnesCount16(uint16(m))
}

func (m CandidateMask) Solution() Cell {
	// caller must ensure IsSolved() == true
	return Cell(bits.TrailingZeros16(uint16(m)))
}
