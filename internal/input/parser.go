package input

import (
	"bmap/internal/matrix"
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const LineBreak = '\n'

type Parser struct {
	reader    *bufio.Reader
	config    *ParserConfig
	validator *validator
}

func NewParser(reader *bufio.Reader, config *ParserConfig) *Parser {
	return &Parser{
		reader:    reader,
		config:    config,
		validator: &validator{config: config},
	}
}

func (p *Parser) Parse() ([]*matrix.Matrix, error) {
	matricesCount, err := p.getMatricesCount()
	if err != nil {
		return nil, err
	}
	if err := p.validator.isMatricesCountValid(matricesCount); err != nil {
		return nil, err
	}
	var matrices []*matrix.Matrix
	for counter := 0; counter < matricesCount; counter++ {
		buildMatrix, err := p.buildMatrix()
		if err != nil {
			return nil, errors.New(
				fmt.Sprintf(
					"Failed to build matrix. Index: %d. Error: %s",
					counter,
					err.Error()))
		}
		matrices = append(matrices, buildMatrix)
	}

	return matrices, nil
}

func (p *Parser) buildMatrix() (*matrix.Matrix, error) {
	rows, columns, err := p.getMatrixSize()
	if err != nil {
		return nil, err
	}
	if err := p.validator.isMatrixSizeValid(rows, columns); err != nil {
		for c := 0; c < rows; c++ {
			_, _ = p.readString()
		}
		return nil, errors.New(fmt.Sprintf("Wrong matrix size. Error: %s", err.Error()))
	}
	buildMatrix := matrix.NewMatrix(rows, columns)
	err = p.getMatrixRows(buildMatrix)
	if err != nil {
		return nil, err
	}

	return buildMatrix, nil
}

func (p *Parser) readString() (string, error) {
	str, err := p.reader.ReadString(LineBreak)
	if err != nil {
		return "", errors.New("cannot read string")
	}
	return str[:len(str)-1], err
}

func (p *Parser) getMatricesCount() (int, error) {
	inputString, err := p.readString()
	testCaseCount, err := strconv.Atoi(inputString)
	if err != nil {
		return 0, errors.New("cannot get test case count: " + err.Error())
	}

	return testCaseCount, nil
}

func (p *Parser) getMatrixSize() (int, int, error) {
	readAndMString, err := p.reader.ReadString(LineBreak)
	arr := strings.Fields(readAndMString[:len(readAndMString)-1])
	if len(arr) != 2 {
		return 0, 0, errors.New("wrong matrix size. It must contain two values")
	}
	rows, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, 0, errors.New("N variable is not number. Details: " + err.Error())
	}
	columns, err := strconv.Atoi(arr[1])
	if err != nil {
		return 0, 0, errors.New("M variable is not number. Details: " + err.Error())
	}

	return rows, columns, nil
}

func (p *Parser) getMatrixRows(matrix *matrix.Matrix) error {
	var resultErr error
	for counter := 0; counter < matrix.GetRows(); counter++ {
		matrixRow, err := p.readString()
		matched, err := regexp.MatchString(fmt.Sprintf("^(0|1){%d,%d}$", matrix.GetColumns(), matrix.GetColumns()), matrixRow)
		if err != nil || !matched {
			resultErr = errors.New(
				fmt.Sprintf(
					"wrong matrix row. Matrix must contain only 1 and 0 with length: %d",
					matrix.GetColumns()))
		}
		for pos, char := range matrixRow {
			resultErr = matrix.AddCell(pos, counter, char == '1')
		}
	}

	return resultErr
}
