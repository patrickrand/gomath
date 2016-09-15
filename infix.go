package main

import (
	"errors"
	"strings"
	"text/scanner"
)

type infix struct{}

var (
	ErrInvalidInfixToken = errors.New("invalid infix token")
)

// Calculate implements the Calculator interface for expressions given in infix notation.
func (i *infix) Calculate(expr string) (float64, error) {
	return 0, nil
}

func (i *infix) Tokenize(expr string) []string {
	scan := new(scanner.Scanner).Init(strings.NewReader(expr))
	scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats

	var tokens []string
	for token := scan.Scan(); token != scanner.EOF; token = scan.Scan() {
		tokens = append(tokens, scan.TokenText())
	}

	return tokens
}
