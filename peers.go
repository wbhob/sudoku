package sudoku

import (
	"fmt"
	"slices"
)

// 8 per row, 8 per col, 8 per box, minus the 2 already in row and 2 in the col
const PEERS = ((SIZE - 1) * 3) - ((BOX - 1) * 2)

var peers = mustPeers()

func mustPeers() [CELLS][PEERS]int {
	var p [CELLS][PEERS]int

	for i := range CELLS {
		row := rowOf(i)
		col := colOf(i)
		box := boxOf(i)

		s := make([]int, 0, 24)
		for j := range SIZE {
			rindex := rowIndex(row, j)
			if rindex != i {
				s = append(s, rindex)
			}

			cindex := colIndex(col, j)
			if cindex != i {
				s = append(s, cindex)
			}

			bindex := boxIndex(box, j)
			if bindex != i {
				s = append(s, bindex)
			}

		}

		slices.Sort(s)
		s = slices.Compact(s)

		if len(s) != 20 {
			panic(fmt.Sprintf("Expected 20 peers, got %d", len(s)))
		}

		copy(p[i][:], s)
	}

	return p
}

func peersContains(a [PEERS]int, v int) bool {
	for i := range PEERS {
		if a[i] == v {
			return true
		}
	}

	return false
}
