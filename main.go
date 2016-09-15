package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	notationType = InfixNotation
	infixFlag    = flag.Bool("infix", true, "Evaluate expression using infix notation.")
	postfixFlag  = flag.Bool("postfix", false, "Evaluate expression using postfix notation.")
)

func main() {
	if len(os.Args) <= 1 {
		exit("no arguments given as input")
	}

	flag.Parse()

	if *postfixFlag == *infixFlag {
		exit("error: --postfix flag and --infix flag cannot have the same value")
	}

	if *postfixFlag {
		notationType = PostfixNotation
	}

	calc, err := NewCalculator(notationType)
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
