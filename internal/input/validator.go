package input

import (
	"errors"
	"fmt"
)

type validator struct {
	config *ParserConfig
}

func (v *validator) isMatricesCountValid(count int) error {
	if v.config.MaximumMatrixCount < count && v.config.MinimumMatrixCount < count {
		return errors.New(
			fmt.Sprintf(
				"Too much matrix count. Maximum: %d, Sent %d",
				v.config.MaximumMatrixCount,
				count))
	}

	return nil
}

func (v *validator) isMatrixSizeValid(rows int, columns int) error {
	if (v.config.MaximumRows < rows && v.config.MinimumRows < rows) ||
		(v.config.MaximumColumns < columns && v.config.MinimumColumns < columns) {
		return errors.New(
			fmt.Sprintf(
				"Wrong matrix format. Sent rows: %d columns: %d . Allowed rows: %d columns: %d.",
				rows, columns,
				v.config.MaximumRows, v.config.MaximumColumns,
			))
	}

	return nil
}
