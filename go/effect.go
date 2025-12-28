package sudoku

type Effect int8

const (
	// EffectNone means the value of the candidate mask did not change.
	EffectNone Effect = 0
	// EffectChanged means the value of the mask changed but there are still more options.
	EffectChanged Effect = 1 << iota
	// EffectSolved means the value of the candidate mask changed and there is exactly one option left.
	EffectSolved
	// EffectContradiction means the value of the candidate mask changed and it has no valid options,
	// so the puzzle state is invalid.
	EffectContradiction
)
