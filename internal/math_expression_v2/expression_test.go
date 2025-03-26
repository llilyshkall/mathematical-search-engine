package math_expression_v2

import "testing"

func Test(t *testing.T) {
	var (
		latex, prefix string
		e, other      *Expression
	)

	latex = "\\int e^x \\, dx"
	prefix = "\\i"
	e, other = ParseLaTeX(latex), ParseLaTeX(prefix)
	e.Compare(other)
}
