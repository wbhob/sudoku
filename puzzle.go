package sudoku

import (
	"strconv"
	"unicode"
)

const (
	BOX   = 3
	SIZE  = BOX * BOX
	CELLS = SIZE * SIZE
)

type Cell = uint8

// Puzzle is a 9x9 grid of integers that represents a sudoku puzzle.
// The integers are in the range 1 to 9, and 0 represents an empty cell.
type Puzzle [CELLS]Cell

// NewPuzzle creates a new puzzle with the given string.
// The string is a string of 81 characters, each character is a digit in the range 1 to 9 or a dot.
// The dots represent empty cells.
func NewPuzzle(s string) Puzzle {
	puzzle := Puzzle{}
	i := 0
	for _, c := range s {
		if unicode.IsSpace(c) {
			continue
		}

		if c == '.' {
			puzzle[i] = 0
		} else {
			puzzle[i] = Cell(c - '0')
		}
		i++
	}
	return puzzle
}

// String returns the string representation of the puzzle.
// The string is a string of 81 characters, each character is a digit in the range 1 to 9 or a dot.
// The dots represent empty cells.
func (p Puzzle) String() string {
	s := ""
	for i, v := range p {
		if v == 0 {
			s += "."
		} else {
			s += strconv.Itoa(int(v))
		}
		if (i+1)%SIZE == 0 {
			s += "\n"
		}
	}
	return s
}

func (p *Puzzle) Set(row, col int, val Cell) {
	index := row*SIZE + col
	p[index] = Cell(val)
}

func (p Puzzle) Row(row int) (result [SIZE]Cell) {
	for i := range SIZE {
		result[i] = p[rowIndex(row, i)]
	}
	return
}

func (p Puzzle) Col(col int) (result [SIZE]Cell) {
	for i := range SIZE {
		result[i] = p[colIndex(col, i)]
	}
	return
}

// Box gets the box. Boxes are numbered:
//
//	0 1 2
//	3 4 5
//	6 7 8
func (p Puzzle) Box(box int) (result [SIZE]Cell) {
	row := box / BOX * BOX
	col := box % BOX * BOX

	for i := range BOX {
		r := row + i
		roffset := r * SIZE
		for j := range BOX {
			c := col + j
			index := roffset + c
			result[i*BOX+j] = p[index]
		}
	}

	return
}

func (p Puzzle) RowOf(index int) int {
	return index / SIZE
}

func (p Puzzle) ColOf(index int) int {
	return index % SIZE
}

func (p Puzzle) BoxOf(index int) int {
	row := p.RowOf(index) / BOX
	col := p.ColOf(index) / BOX

	return row*BOX + col
}

func (p Puzzle) Valid() bool {
	for i := range SIZE {
		if !nodup(p.Row(i)) || !nodup(p.Col(i)) || !nodup(p.Box(i)) {
			return false
		}
	}
	return true
}

func rowIndex(r, i int) int { return r*SIZE + i }
func colIndex(c, i int) int { return i*SIZE + c }

// nodup checks that there are no duplicates in the array, ignoring 0.
func nodup(row [SIZE]Cell) bool {
	var seen uint
	for i := range row {
		v := row[i]
		if v == 0 {
			continue
		}

		if seen&(1<<v) != 0 {
			return false
		}

		seen |= 1 << v
	}

	return true
}
