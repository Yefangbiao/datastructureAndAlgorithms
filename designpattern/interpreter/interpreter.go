package main

import "strings"

type Expression interface {
	Interpret(context string) bool
}

type TerminalExpression struct {
	Word string
}

// 终结符
func (te *TerminalExpression) Interpret(context string) bool {
	if strings.Contains(context, te.Word) {
		return true
	}
	return false
}

// 或
type OrExpression struct {
	A Expression
	B Expression
}

func (oe *OrExpression) Interpret(context string) bool {
	return oe.A.Interpret(context) || oe.B.Interpret(context)
}

// 与
type AndExpression struct {
	A Expression
	B Expression
}

func (ae *AndExpression) Interpret(context string) bool {
	return ae.A.Interpret(context) && ae.B.Interpret(context)
}
