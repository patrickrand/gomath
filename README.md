# gomath
`gomath` is a Go CLI that attempts to act as a mathematical calculator.

## Installation and Usage
```bash
go get github.com/patrickrand/gomath

gomath '3 + 4 * sin(12)'
0.8537083279982602   

gomath --postfix '10.25 -8 * 124 +'
42.000000
```

### Features
- supports both infix and postfix notation

### Supported Operations
- `+`, `-`, `*`, `/`, `pow`, `sqrt`
- `lg`, `ln`, `lb`, `abs`
- `cos`, `sin`, `tan`