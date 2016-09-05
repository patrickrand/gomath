package calculator

type Expression interface {
	Evaluate(v ...float64) (float64, error)
}
