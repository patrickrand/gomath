package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/patrickrand/gomath/calculator"
)

var (
	postfixFlag = flag.Bool("postfix", false, "Evaluate expression using postfix notation.")
)

func main() {
	if len(os.Args) <= 1 {
		exit("no arguments given as input")
	}

	flag.Parse()

	if !*postfixFlag {
		exit("--postfix flag is required, as it is the only notation currently supported")
	}

	calc, err := calculator.New(calculator.POSTFIX)
	if err != nil {
		exit(err)
	}

	expr := os.Args[len(os.Args)-1] // expression string must be the final argument

	result, err := calc.Calculate(expr)
	if err != nil {
		exit(err)
	}

	fmt.Printf("%v = %v\n", os.Args[len(os.Args)-1], result)
}

func exit(v ...interface{}) {
	fmt.Print("error: ")
	fmt.Println(v...)
	os.Exit(1)
}
