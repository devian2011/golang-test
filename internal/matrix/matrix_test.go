package matrix

import (
	"testing"
)

func TestMatrix_AddCell(t *testing.T) {
	m := NewMatrix(1, 1)
	emptyErr := m.AddCell(0, 0, true)
	if emptyErr != nil {
		t.Errorf("Wrong behavior for add cell to matrix. Cell must be appended")
	}
	addCellErr := m.AddCell(1, 1, false)
	if addCellErr == nil {
		t.Errorf("Wrong behavior for add cell to matrix. Cell cannot be appended")
	}
}

func TestMatrix_Validate(t *testing.T) {
	tables := []struct {
		name    string
		rows    int
		columns int
		cell    []struct {
			x       int
			y       int
			isWhite bool
		}
		isValid bool
		message string
	}{
		{
			name:    "Check that last cell is white",
			rows:    2,
			columns: 2,
			cell: []struct {
				x       int
				y       int
				isWhite bool
			}{
				{
					x:       0,
					y:       0,
					isWhite: false,
				},
				{
					x:       1,
					y:       0,
					isWhite: false,
				},
				{
					x:       0,
					y:       1,
					isWhite: false,
				},
				{
					x:       1,
					y:       1,
					isWhite: false,
				},
			},
			isValid: false,
			message: "last cell of matrix must be white",
		},
		{
			name:    "Check that everything is OK",
			rows:    2,
			columns: 2,
			cell: []struct {
				x       int
				y       int
				isWhite bool
			}{
				{
					x:       0,
					y:       0,
					isWhite: false,
				},
				{
					x:       1,
					y:       0,
					isWhite: false,
				},
				{
					x:       0,
					y:       1,
					isWhite: false,
				},
				{
					x:       1,
					y:       1,
					isWhite: true,
				},
			},
			isValid: true,
			message: "",
		},
		{
			name:    "Check matrix definition is not the same as count of matrix cells",
			rows:    3,
			columns: 3,
			cell: []struct {
				x       int
				y       int
				isWhite bool
			}{
				{
					x:       0,
					y:       0,
					isWhite: false,
				},
				{
					x:       0,
					y:       1,
					isWhite: false,
				},
			},
			isValid: false,
			message: "matrix size is not the same as matrix definition",
		},
	}

	for _, row := range tables {
		m := NewMatrix(row.rows, row.columns)
		for _, cell := range row.cell {
			_ = m.AddCell(cell.x, cell.y, cell.isWhite)
		}
		validationErr := m.Validate()
		if (validationErr == nil) != row.isValid {
			message := "Empty"
			messageCorrespond := "correspond"
			if validationErr != nil {
				message = validationErr.Error()
				if message != row.message {
					messageCorrespond = "not correspond"
				}
			}
			t.Errorf(
				"Matrix validation method has errors in test case: %s. Message is %s. Result is: %s",
				row.name,
				messageCorrespond,
				message,
			)
		}
	}

}
