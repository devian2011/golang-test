package internal

import (
	"bmap/internal/input"
	"bmap/internal/matrix"
	"bmap/internal/output"
	"bufio"
)

type Application struct {
	config     *ApplicationConfig
	parser     *input.Parser
	writer     *output.Writer
	calculator *matrix.Calculator
}

func NewApplication(config *ApplicationConfig, reader *bufio.Reader, writer *bufio.Writer, error *bufio.Writer) *Application {
	return &Application{
		config:     config,
		parser:     input.NewParser(reader, config.Parser),
		writer:     output.NewWriter(writer, error),
		calculator: matrix.NewCalculator(),
	}
}

func (app *Application) Run() {
	matrices, err := app.parser.Parse()
	if err != nil {
		app.writer.WriteErr(err)
		return
	}
	var result [][][]string
	for _, m := range matrices {
		result = append(result, app.calculator.GetMatrixWithNearestDistancesBtwBlackAndWhiteCells(*m))
	}

	app.writer.Write(result)
}
