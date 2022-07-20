package behavior_pattern

import (
	"fmt"
	"testing"
)

func TestInterpreterPattern(t *testing.T) {
	marriedExpr := NewTerminalExpression("married")
	hasChildExpr := NewTerminalExpression("hasChild")

	fmt.Println(NewAndExpression(marriedExpr, hasChildExpr).Interpreter("married hasChild"))
	fmt.Println(NewAndExpression(marriedExpr, hasChildExpr).Interpreter("married"))
	fmt.Println(NewAndExpression(marriedExpr, hasChildExpr).Interpreter("hasChild"))
	fmt.Println(NewOrExpression(marriedExpr, hasChildExpr).Interpreter("hasChild"))
	fmt.Println(NewOrExpression(marriedExpr, hasChildExpr).Interpreter("married"))
}
