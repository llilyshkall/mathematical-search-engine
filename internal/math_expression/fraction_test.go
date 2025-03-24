package math_expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFraction_ToString_Simple(t *testing.T) {
	var (
		sNumerator, sDenumerator, expect string
	)
	check := func(t *testing.T, sNumerator, sDenumerator, expect string) {
		var (
			eNumerator, eDenumerator *Expression
			f                        fraction
		)
		eNumerator = ParseLaTeX(sNumerator)
		eDenumerator = ParseLaTeX(sDenumerator)
		f = fraction{
			numerator:   eNumerator,
			denumerator: eDenumerator,
		}
		actual := f.toString()
		assert.Equal(t, expect, actual)
	}

	sNumerator = ""
	sDenumerator = ""
	expect = "\\frac{}{}"
	check(t, sNumerator, sDenumerator, expect)

	sNumerator = "2"
	sDenumerator = ""
	expect = "\\frac{2}{}"
	check(t, sNumerator, sDenumerator, expect)

	sNumerator = ""
	sDenumerator = "2"
	expect = "\\frac{}{2}"
	check(t, sNumerator, sDenumerator, expect)

	sNumerator = "2+2"
	sDenumerator = "2"
	expect = "\\frac{2 + 2}{2}"
	check(t, sNumerator, sDenumerator, expect)

	sNumerator = "\\alpha"
	sDenumerator = "2"
	expect = "\\frac{\\alpha}{2}"
	check(t, sNumerator, sDenumerator, expect)
}

func TestFraction_Equals_Simple(t *testing.T) {
	var (
		sNumerator1, sDenumerator1 string
		sNumerator2, sDenumerator2 string
		expect                     bool
	)
	check := func(t *testing.T, sDenumerator1, sNumerator1, sDenumerator2, sNumerator2 string, expect bool) {
		var (
			eDenumerator1, eNumerator1 *Expression
			eDenumerator2, eNumerator2 *Expression
			f1, f2                     fraction
		)
		eDenumerator1 = ParseLaTeX(sDenumerator1)
		eDenumerator2 = ParseLaTeX(sDenumerator2)
		eNumerator1 = ParseLaTeX(sNumerator1)
		eNumerator2 = ParseLaTeX(sNumerator2)
		f1 = fraction{denumerator: eDenumerator1, numerator: eNumerator1}
		f2 = fraction{denumerator: eDenumerator2, numerator: eNumerator2}
		if expect {
			assert.True(t, f1.equals(f2))
			assert.True(t, f2.equals(f1))
		} else {
			assert.False(t, f1.equals(f2))
			assert.False(t, f2.equals(f1))
		}
	}

	sNumerator1 = ""
	sDenumerator1 = ""
	sNumerator2 = ""
	sDenumerator2 = ""
	expect = true
	check(t, sNumerator1, sDenumerator1, sNumerator2, sDenumerator2, expect)

	sNumerator1 = "a"
	sDenumerator1 = ""
	sNumerator2 = ""
	sDenumerator2 = ""
	expect = false
	check(t, sNumerator1, sDenumerator1, sNumerator2, sDenumerator2, expect)

	sNumerator1 = "a"
	sDenumerator1 = ""
	sNumerator2 = "a  b"
	sDenumerator2 = ""
	expect = false
	check(t, sNumerator1, sDenumerator1, sNumerator2, sDenumerator2, expect)

	sNumerator1 = "ab"
	sDenumerator1 = ""
	sNumerator2 = "a  b"
	sDenumerator2 = ""
	expect = true
	check(t, sNumerator1, sDenumerator1, sNumerator2, sDenumerator2, expect)

	sNumerator1 = ""
	sDenumerator1 = "a  b"
	sNumerator2 = "a  b"
	sDenumerator2 = ""
	expect = false
	check(t, sNumerator1, sDenumerator1, sNumerator2, sDenumerator2, expect)

	sNumerator1 = "a  b"
	sDenumerator1 = "a  b"
	sNumerator2 = "a  b"
	sDenumerator2 = "ab"
	expect = true
	check(t, sNumerator1, sDenumerator1, sNumerator2, sDenumerator2, expect)

	sNumerator1 = "\\alpha  b"
	sDenumerator1 = "a  b"
	sNumerator2 = "\\a  b"
	sDenumerator2 = "ab"
	expect = false
	check(t, sNumerator1, sDenumerator1, sNumerator2, sDenumerator2, expect)

	sNumerator1 = "\\alpha + \\beta"
	sDenumerator1 = "a  b"
	sNumerator2 = "\\alpha + \\beta"
	sDenumerator2 = "ab"
	expect = true
	check(t, sNumerator1, sDenumerator1, sNumerator2, sDenumerator2, expect)

	sNumerator1 = "\\alpha + \\be"
	sDenumerator1 = "a  b"
	sNumerator2 = "\\alpha + \\beta"
	sDenumerator2 = "ab"
	expect = false
	check(t, sNumerator1, sDenumerator1, sNumerator2, sDenumerator2, expect)

	sNumerator1 = "\\alpha + \\be"
	sDenumerator1 = "a  b"
	sNumerator2 = "\\alpha + \\beta + \\omega"
	sDenumerator2 = "ab"
	expect = false
	check(t, sNumerator1, sDenumerator1, sNumerator2, sDenumerator2, expect)

	sNumerator1 = ""
	sDenumerator1 = "\\alpha + \\be"
	sNumerator2 = ""
	sDenumerator2 = "\\alpha + \\beta + \\omega"
	expect = false
	check(t, sNumerator1, sDenumerator1, sNumerator2, sDenumerator2, expect)

	sNumerator1 = ""
	sDenumerator1 = "\\alpha + \\beta +"
	sNumerator2 = ""
	sDenumerator2 = "\\alpha + \\beta + \\omega"
	expect = false
	check(t, sNumerator1, sDenumerator1, sNumerator2, sDenumerator2, expect)

	sNumerator1 = ""
	sDenumerator1 = "\\alpha + \\beta + \\omega"
	sNumerator2 = ""
	sDenumerator2 = "\\alpha + \\beta + \\omega"
	expect = true
	check(t, sNumerator1, sDenumerator1, sNumerator2, sDenumerator2, expect)

	sNumerator1 = "a+b"
	sDenumerator1 = "\\alpha + \\beta + \\omega"
	sNumerator2 = "a+b"
	sDenumerator2 = "\\alpha + \\beta + \\omega"
	expect = true
	check(t, sNumerator1, sDenumerator1, sNumerator2, sDenumerator2, expect)
}

func TestFraction_ParseFraction_Simple(t *testing.T) {
	var (
		latex       string
		expectF     lexema
		expectShift int
		expectLatex string
	)
	check := func(t *testing.T, latex string, expectF lexema, expectShifh int, expectLatex string) {
		var (
			actualF     lexema
			actualShift int
			actualLatex string
		)
		actualF, actualShift = parseFraction(latex)
		actualLatex = actualF.toString()
		assert.True(t, actualF.equals(expectF))
		assert.Equal(t, expectShifh, actualShift)
		assert.Equal(t, expectLatex, actualLatex)
	}

	latex = ""
	expectF = fraction{}
	expectShift = 0
	expectLatex = "\\frac{}{}"
	check(t, latex, expectF, expectShift, expectLatex)

	latex = "1"
	expectF = fraction{}
	expectShift = 0
	expectLatex = "\\frac{}{}"
	check(t, latex, expectF, expectShift, expectLatex)

	latex = "{1}{2}"
	expectF = fraction{
		numerator:   &Expression{lexemas: []lexema{operand{name: "1"}}},
		denumerator: &Expression{lexemas: []lexema{operand{name: "2"}}},
	}
	expectShift = 6
	expectLatex = "\\frac{1}{2}"
	check(t, latex, expectF, expectShift, expectLatex)

	latex = "{1}{2a}"
	expectF = fraction{
		numerator: &Expression{lexemas: []lexema{
			operand{name: "1"},
		}},
		denumerator: &Expression{lexemas: []lexema{
			operand{name: "2"},
			operand{name: "a"},
		}},
	}
	expectShift = 7
	expectLatex = "\\frac{1}{2 a}"
	check(t, latex, expectF, expectShift, expectLatex)

	latex = "{1}{2a"
	expectF = fraction{
		numerator: &Expression{lexemas: []lexema{
			operand{name: "1"},
		}},
		denumerator: &Expression{lexemas: []lexema{
			operand{name: "2"},
			operand{name: "a"},
		}},
	}
	expectShift = 6
	expectLatex = "\\frac{1}{2 a}"
	check(t, latex, expectF, expectShift, expectLatex)

	latex = "{1{2a}"
	expectF = fraction{
		numerator: &Expression{lexemas: []lexema{
			operand{name: "1"},
			operand{name: "2"},
			operand{name: "a"},
		}},
		denumerator: &Expression{lexemas: []lexema{}},
	}
	expectShift = 6
	expectLatex = "\\frac{1 2 a}{}"
	check(t, latex, expectF, expectShift, expectLatex)
}
