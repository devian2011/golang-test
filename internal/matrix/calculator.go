package matrix

import (
	"math"
	"strconv"
)

type Calculator struct {
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) GetMatrixWithNearestDistancesBtwBlackAndWhiteCells(matrix Matrix) [][]string {
	result := make([][]string, matrix.GetRows())
	for i := range result {
		result[i] = make([]string, matrix.GetColumns())
	}
	whiteCells := c.getWhiteCells(matrix)
	for _, mCell := range matrix.cells {
		result[mCell.Y][mCell.X] = strconv.Itoa(c.getDistanceForNearestWhiteCell(mCell, whiteCells))
	}

	return result
}

func (c *Calculator) getWhiteCells(matrix Matrix) []Cell {
	var whiteCells []Cell
	for _, cell := range matrix.cells {
		if cell.IsWhite {
			whiteCells = append(whiteCells, cell)
		}
	}
	return whiteCells
}

func (c *Calculator) getDistanceForNearestWhiteCell(cell Cell, whiteCells []Cell) int {
	distance := 0
	// I don't know why, but this behavior has been written in example
	if cell.IsWhite {
		return distance
	}
	for _, wCell := range whiteCells {
		calculatedDistance := c.calculateDistanceBtwCells(cell, wCell)
		if distance == 0 || distance > calculatedDistance {
			distance = calculatedDistance
		}
	}
	return distance
}

func (c *Calculator) calculateDistanceBtwCells(cell1 Cell, cell2 Cell) int {
	return int(math.Abs(float64(cell1.X)-float64(cell2.X)) + math.Abs(float64(cell1.Y)-float64(cell2.Y)))
}
