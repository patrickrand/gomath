package calculator

import "fmt"

type Calculator interface {
	Calculate() (float64, error)
}

type notation string

const (
	PREFIX  notation = "prefix"
	INFIX   notation = "infix"
	POSTFIX notation = "postfix"
)

func New(expr string, n notation) (Calculator, error) {
	switch n {
	case POSTFIX:
		return parsePostfix(expr)
	default:
		return nil, fmt.Errorf("whoops! %q notation is not implemented", string(n))
	}
}
