package testdata

import _ "embed"

//go:embed puzzle1.txt
var Puzzle1 string

//go:embed puzzle1solution.txt
var Puzzle1Solution string

//go:embed solved.txt
var SolvedPuzzle string

//go:embed almostsolved.txt
var AlmostSolvedPuzzle string
