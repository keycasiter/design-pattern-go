package behavior_pattern

import "strings"

type Expression interface {
	Interpreter(data string) bool
}

type TerminalExpression struct {
	data string
}

func NewTerminalExpression(data string) *TerminalExpression {
	return &TerminalExpression{data: data}
}

func (t *TerminalExpression) Interpreter(data string) bool {
	if strings.Contains(data, t.data) {
		return true
	}
	return false
}

type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func NewAndExpression(expr1 Expression, expr2 Expression) *AndExpression {
	return &AndExpression{expr1: expr1, expr2: expr2}
}

func (t *AndExpression) Interpreter(context string) bool {
	if t.expr1.Interpreter(context) && t.expr2.Interpreter(context) {
		return true
	}
	return false
}

type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func NewOrExpression(expr1 Expression, expr2 Expression) *OrExpression {
	return &OrExpression{expr1: expr1, expr2: expr2}
}

func (t *OrExpression) Interpreter(context string) bool {
	if t.expr1.Interpreter(context) || t.expr2.Interpreter(context) {
		return true
	}
	return false
}
