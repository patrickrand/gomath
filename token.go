package main

type TokenType int

const (
	NumberToken TokenType = iota
	OperatorToken
	FunctionToken
	CommaToken
	LeftParenthsToken
	RightParenthsToken
	ErrorToken
)

func GetTokenType(token string) TokenType {
	if _, err := ParseFloat(token); err == nil {
		return NumberToken
	}

	if _, ok := Operator(token); ok {
		return OperatorToken
	}

	if _, ok := Function(token); ok {
		return FunctionToken
	}

	if token == "," {
		return CommaToken
	}

	if token == "(" {
		return LeftParenthsToken
	}

	if token == ")" {
		return RightParenthsToken
	}

	return ErrorToken
}
