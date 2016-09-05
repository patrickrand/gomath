package calculator

import "errors"

type Calculator interface {
	Calculate(expr string) (float64, error)
}

type notation string

const (
	PREFIX  notation = "prefix"
	INFIX   notation = "infix"
	POSTFIX notation = "postfix"
)

var ErrNotationNotImplemented = errors.New("notation not implemented")

func New(n notation) (Calculator, error) {
	switch n {
	case POSTFIX:
		return &postfix{}, nil
	default:
		return nil, ErrNotationNotImplemented
	}
}
