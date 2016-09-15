package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	notationType = InfixNotation
	postfixFlag  = flag.Bool("postfix", false, "Evaluate expression using postfix notation.")
)

func main() {
	if len(os.Args) <= 1 {
		exit("no arguments given as input")
	}

	expr := os.Args[len(os.Args)-1] // expression string must be the final argument
	os.Args = os.Args[:len(os.Args)-1]

	flag.Parse()

	if *postfixFlag {
		notationType = PostfixNotation
	}

	calc, err := NewCalculator(notationType)
	if err != nil {
		exit(err)
	}

	result, err := calc.Calculate(expr)
	if err != nil {
		exit(err)
	}

	fmt.Println(result)
}

func exit(v ...interface{}) {
	fmt.Print("error: ")
	fmt.Println(v...)
	os.Exit(1)
}
