package output

import (
	"bufio"
	"fmt"
	"strings"
)

type Writer struct {
	writer *bufio.Writer
	error  *bufio.Writer
}

func NewWriter(writer *bufio.Writer, error *bufio.Writer) *Writer {
	return &Writer{
		writer: writer,
		error:  error,
	}
}

func (w *Writer) Write(matrices [][][]string) {
	for i, m := range matrices {
		w.writer.Write([]byte(fmt.Sprintf("Matrix %d \n", i)))
		for _, row := range m {
			w.writer.Write([]byte(strings.Join(row, " ") + "\n"))
		}
		w.writer.Write([]byte("\n\n"))
		w.writer.Flush()
	}
}

func (w *Writer) WriteErr(err error) {
	b, errorWriteErr := w.error.Write([]byte(err.Error()))
	if errorWriteErr != nil {
		panic(errorWriteErr.Error())
	}
	flushErr := w.error.Flush()
	if flushErr != nil {
		panic(flushErr.Error())
	}
	fmt.Println(b)
}
