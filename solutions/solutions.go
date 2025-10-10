package solutions

import (
	"advent-of-go/utils"
	"slices"
)

func Solutions() []utils.Solution {
	return slices.Concat[[]utils.Solution]()
}
