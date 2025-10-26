package day11

import (
	"advent-of-go/utils"
	"fmt"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 11,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	grid := utils.NewInfGrid[color]()
	grid.Set(0, 0, White)
	runBot(input, grid)
	finGrid, e := grid.ToFinGrid()
	if e != nil {
		return "", fmt.Errorf("Error converting infinite grid to finite grid for image processing: %w", e)
	}
	return utils.GridToPngString(finGrid)
}