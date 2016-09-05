package main

import (
	"fmt"
	"os"

	"github.com/patrickrand/gomath/calculator"
)

func main() {
	if len(os.Args) <= 1 {
		exit("no expression given as input")
	}

	c, err := calculator.New(os.Args[1], calculator.POSTFIX)
	if err != nil {
		exit(err)
	}

	result, err := c.Calculate()
	if err != nil {
		exit(err)
	}

	fmt.Printf("%v = %f\n", os.Args[1:], result)
}

func exit(v ...interface{}) {
	fmt.Print("error: ")
	fmt.Println(v...)
	os.Exit(1)
}
