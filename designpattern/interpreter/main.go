package main

import "fmt"

func main() {
	e1 := &TerminalExpression{"a"}
	e2 := &TerminalExpression{"b"}

	and := AndExpression{
		A: e1,
		B: e2,
	}
	fmt.Println(and.Interpret("a"))

	or := OrExpression{
		A: e1,
		B: e2,
	}
	fmt.Println(or.Interpret("a"))
}
