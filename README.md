# gomath
`gomath` is a Go CLI that attempts to act as a mathematical calculator.
Currently, only [postfix (Reverse Polish) notation](https://en.wikipedia.org/wiki/Reverse_Polish_notation) is accepted as input.

## Installation and Usage
```bash
go get github.com/patrickrand/gomath

gomath --postfix '10.25 -8 * 124 +'
10.25 -8 * 124 + = 42.000000
```

### Supported Operations
- `+`, `-`, `*`, `/`
- `pow`, `sqrt`, `lg`, `ln`, `lb`, `abs`
- `cos`, `sin`, `tan`