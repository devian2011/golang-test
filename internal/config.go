package internal

import "bmap/internal/input"

type ApplicationConfig struct {
	Env    string
	Parser *input.ParserConfig
}
