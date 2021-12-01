package main

import (
	"bmap/internal"
	"bmap/internal/input"
	"bufio"
	"flag"
	"os"
)

var (
	envType = flag.String("env", "debug", "Environment variable (debug, production and so on)")
)

func main() {
	flag.Parse()
	app := internal.NewApplication(
		&internal.ApplicationConfig{
			Env: *envType,
			Parser: &input.ParserConfig{
				MaximumMatrixCount: 1000,
				MaximumRows:        182,
				MaximumColumns:     182,
			},
		},
		bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout),
		bufio.NewWriter(os.Stderr))

	app.Run()
}
