# gomath
`gomath` is a Go CLI that attempts to act as a mathematical calculator.
Currently, only [postfix (Reverse Polish) notation](https://en.wikipedia.org/wiki/Reverse_Polish_notation) is accepted as input.

## Installation and usage
```bash
go get github.com/patrickrand/gomath

gomath --postfix '10.25 -8 * 124 +'
10.25 -8 * 124 + = 42.000000
```