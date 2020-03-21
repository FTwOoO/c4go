package types

import (
	"github.com/FTwOoO/c4go/program"
)

// ResolveTypeForBinaryOperator determines the result Go type when performing a
// binary expression.
func ResolveTypeForBinaryOperator(p *program.Program, operator, leftType, rightType string) string {
	if operator == "==" ||
		operator == "!=" ||
		operator == ">" ||
		operator == ">=" ||
		operator == "<" ||
		operator == "<=" {
		return "bool"
	}

	return leftType
}
