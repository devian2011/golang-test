package matrix

import "errors"

type Matrix struct {
	rows            int
	columns         int
	cells           []Cell
	matrixCellCount int
}

func NewMatrix(rows int, columns int) *Matrix {
	return &Matrix{
		rows:            rows,
		columns:         columns,
		cells:           []Cell{},
		matrixCellCount: rows * columns,
	}
}

func (m *Matrix) GetRows() int {
	return m.rows
}

func (m *Matrix) GetColumns() int {
	return m.columns
}

func (m *Matrix) GetCells() []Cell {
	return m.cells
}

func (m *Matrix) AddCell(X int, Y int, isWhite bool) error {
	if len(m.cells) >= m.matrixCellCount {
		return errors.New("cannot append new cell for matrix, all cells have been defined")
	}
	m.cells = append(m.cells, Cell{
		X:       X,
		Y:       Y,
		IsWhite: isWhite,
	})
	return nil
}

func (m *Matrix) Validate() error {
	if len(m.cells) != m.matrixCellCount {
		return errors.New("matrix size is not the same as matrix definition")
	}
	if !m.cells[len(m.cells)-1].IsWhite {
		return errors.New("last cell of matrix must be white")
	}

	return nil
}

type Cell struct {
	X       int
	Y       int
	IsWhite bool
}
