package main

import (
	"flag"
	"fmt"
	"os"
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

	postfix, err := NewCalculator(POSTFIX)
	if err != nil {
		exit(err)
	}

	expr := os.Args[len(os.Args)-1] // expression string must be the final argument

	result, err := postfix.Calculate(expr)
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
