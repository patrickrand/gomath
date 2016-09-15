package main

import "errors"

type Calculator interface {
	Calculate(expr string) (float64, error)
}

func NewCalculator(n notation) (Calculator, error) {
	switch n {
	case InfixNotation:
		return &infix{}, nil
	case PostfixNotation:
		return &postfix{}, nil
	default:
		return nil, ErrNotationNotImplemented
	}
}

type notation string

const (
	PrefixNotation  notation = "prefix"
	InfixNotation   notation = "infix"
	PostfixNotation notation = "postfix"
)

var ErrNotationNotImplemented = errors.New("notation not implemented")
