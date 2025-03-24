package math_expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoot_ToString_Simple(t *testing.T) {
	var (
		sDegree, sArg, expect string
	)
	check := func(t *testing.T, sDegree, sArg, expect string) {
		var (
			eDegree, eArg *Expression
			r             root
		)
		eDegree = ParseLaTeX(sDegree)
		eArg = ParseLaTeX(sArg)
		r = root{
			degree: eDegree,
			arg:    eArg,
		}
		actual := r.toString()
		assert.Equal(t, expect, actual)
	}

	sDegree = ""
	sArg = ""
	expect = "\\sqrt[]{}"
	check(t, sDegree, sArg, expect)

	sDegree = ""
	sArg = "2"
	expect = "\\sqrt[]{2}"
	check(t, sDegree, sArg, expect)

	sDegree = ""
	sArg = "2+2"
	expect = "\\sqrt[]{2 + 2}"
	check(t, sDegree, sArg, expect)

	sDegree = "5"
	sArg = "2+2"
	expect = "\\sqrt[5]{2 + 2}"
	check(t, sDegree, sArg, expect)

	sDegree = "5+5"
	sArg = "2+2"
	expect = "\\sqrt[5 + 5]{2 + 2}"
	check(t, sDegree, sArg, expect)
}

func TestRoot_Equals_Simple(t *testing.T) {
	var (
		sArg1, sDegree1 string
		sArg2, sDegree2 string
		expect          bool
	)
	check := func(t *testing.T, sDegree1, sArg1, sDegree2, sArg2 string, expect bool) {
		var (
			eDegree1, eArg1 *Expression
			eDegree2, eArg2 *Expression
			r1, r2          root
		)
		eDegree1 = ParseLaTeX(sDegree1)
		eDegree2 = ParseLaTeX(sDegree2)
		eArg1 = ParseLaTeX(sArg1)
		eArg2 = ParseLaTeX(sArg2)
		r1 = root{degree: eDegree1, arg: eArg1}
		r2 = root{degree: eDegree2, arg: eArg2}
		if expect {
			assert.True(t, r1.equals(r2))
			assert.True(t, r2.equals(r1))
		} else {
			assert.False(t, r1.equals(r2))
			assert.False(t, r2.equals(r1))
		}
	}

	sArg1 = ""
	sDegree1 = ""
	sArg2 = ""
	sDegree2 = ""
	expect = true
	check(t, sArg1, sDegree1, sArg2, sDegree2, expect)

	sArg1 = "a"
	sDegree1 = ""
	sArg2 = ""
	sDegree2 = ""
	expect = false
	check(t, sArg1, sDegree1, sArg2, sDegree2, expect)

	sArg1 = "a"
	sDegree1 = ""
	sArg2 = "a  b"
	sDegree2 = ""
	expect = false
	check(t, sArg1, sDegree1, sArg2, sDegree2, expect)

	sArg1 = "ab"
	sDegree1 = ""
	sArg2 = "a  b"
	sDegree2 = ""
	expect = true
	check(t, sArg1, sDegree1, sArg2, sDegree2, expect)

	sArg1 = ""
	sDegree1 = "a  b"
	sArg2 = "a  b"
	sDegree2 = ""
	expect = false
	check(t, sArg1, sDegree1, sArg2, sDegree2, expect)

	sArg1 = "a  b"
	sDegree1 = "a  b"
	sArg2 = "a  b"
	sDegree2 = "ab"
	expect = true
	check(t, sArg1, sDegree1, sArg2, sDegree2, expect)

	sArg1 = "\\alpha  b"
	sDegree1 = "a  b"
	sArg2 = "\\a  b"
	sDegree2 = "ab"
	expect = false
	check(t, sArg1, sDegree1, sArg2, sDegree2, expect)

	sArg1 = "\\alpha + \\beta"
	sDegree1 = "a  b"
	sArg2 = "\\alpha + \\beta"
	sDegree2 = "ab"
	expect = true
	check(t, sArg1, sDegree1, sArg2, sDegree2, expect)

	sArg1 = "\\alpha + \\be"
	sDegree1 = "a  b"
	sArg2 = "\\alpha + \\beta"
	sDegree2 = "ab"
	expect = false
	check(t, sArg1, sDegree1, sArg2, sDegree2, expect)

	sArg1 = "\\alpha + \\be"
	sDegree1 = "a  b"
	sArg2 = "\\alpha + \\beta + \\omega"
	sDegree2 = "ab"
	expect = false
	check(t, sArg1, sDegree1, sArg2, sDegree2, expect)

	sArg1 = ""
	sDegree1 = "\\alpha + \\be"
	sArg2 = ""
	sDegree2 = "\\alpha + \\beta + \\omega"
	expect = false
	check(t, sArg1, sDegree1, sArg2, sDegree2, expect)

	sArg1 = ""
	sDegree1 = "\\alpha + \\beta +"
	sArg2 = ""
	sDegree2 = "\\alpha + \\beta + \\omega"
	expect = false
	check(t, sArg1, sDegree1, sArg2, sDegree2, expect)

	sArg1 = ""
	sDegree1 = "\\alpha + \\beta + \\omega"
	sArg2 = ""
	sDegree2 = "\\alpha + \\beta + \\omega"
	expect = true
	check(t, sArg1, sDegree1, sArg2, sDegree2, expect)

	sArg1 = "a+b"
	sDegree1 = "\\alpha + \\beta + \\omega"
	sArg2 = "a+b"
	sDegree2 = "\\alpha + \\beta + \\omega"
	expect = true
	check(t, sArg1, sDegree1, sArg2, sDegree2, expect)
}

func TestRoot_ParseRoot_Simple(t *testing.T) {
	var (
		latex       string
		expectR     lexema
		expectLatex string
		expectShift int
	)
	check := func(t *testing.T, latex string, expectR lexema, expectLatex string, expectShift int) {
		var (
			actualR     lexema
			actualLatex string
			actualShift int
		)
		actualR, actualShift = parseRoot(latex)
		actualLatex = actualR.toString()
		assert.True(t, actualR.equals(expectR))
		assert.True(t, expectR.equals(actualR))
		assert.Equal(t, expectLatex, actualLatex)
		assert.Equal(t, expectShift, actualShift)
	}

	latex = ""
	expectShift = 0
	expectR = root{}
	expectLatex = "\\sqrt{}"
	check(t, latex, expectR, expectLatex, expectShift)

	latex = "2"
	expectShift = 0
	expectR = root{}
	expectLatex = "\\sqrt{}"
	check(t, latex, expectR, expectLatex, expectShift)

	latex = "{2}"
	expectShift = 3
	expectR = root{
		arg: &Expression{lexemas: []lexema{
			operand{name: "2"},
		}}}
	expectLatex = "\\sqrt{2}"
	check(t, latex, expectR, expectLatex, expectShift)

	latex = "{2"
	expectShift = 2
	expectR = root{
		arg: &Expression{lexemas: []lexema{
			operand{name: "2"},
		}}}
	expectLatex = "\\sqrt{2}"
	check(t, latex, expectR, expectLatex, expectShift)

	latex = "[3]{2}"
	expectShift = 6
	expectR = root{
		degree: &Expression{lexemas: []lexema{
			operand{name: "3"},
		}},
		arg: &Expression{lexemas: []lexema{
			operand{name: "2"},
		}}}
	expectLatex = "\\sqrt[3]{2}"
	check(t, latex, expectR, expectLatex, expectShift)

	latex = "[3{2}"
	expectShift = 5
	expectR = root{
		degree: &Expression{lexemas: []lexema{
			operand{name: "3"},
			operand{name: "2"},
		}},
		arg: &Expression{lexemas: []lexema{}}}
	expectLatex = "\\sqrt[3 2]{}"
	check(t, latex, expectR, expectLatex, expectShift)

	latex = "[a + b]{c^2+b_2}"
	expectShift = 16
	expectR = root{
		degree: &Expression{lexemas: []lexema{
			operand{name: "a"},
			operand{name: "+"},
			operand{name: "b"},
		}},
		arg: &Expression{lexemas: []lexema{
			operand{name: "c"},
			index{high: &Expression{lexemas: []lexema{operand{name: "2"}}}},
			operand{name: "+"},
			operand{name: "b"},
			index{low: &Expression{lexemas: []lexema{operand{name: "2"}}}},
		}}}
	expectLatex = "\\sqrt[a + b]{c ^2 + b _2}"
	check(t, latex, expectR, expectLatex, expectShift)

	latex = "[a + b]{c^2+b_2"
	expectShift = 15
	expectR = root{
		degree: &Expression{lexemas: []lexema{
			operand{name: "a"},
			operand{name: "+"},
			operand{name: "b"},
		}},
		arg: &Expression{lexemas: []lexema{
			operand{name: "c"},
			index{high: &Expression{lexemas: []lexema{operand{name: "2"}}}},
			operand{name: "+"},
			operand{name: "b"},
			index{low: &Expression{lexemas: []lexema{operand{name: "2"}}}},
		}}}
	expectLatex = "\\sqrt[a + b]{c ^2 + b _2}"
	check(t, latex, expectR, expectLatex, expectShift)

	latex = "[a + b]{c^2+b_2} a b c"
	expectShift = 16
	expectR = root{
		degree: &Expression{lexemas: []lexema{
			operand{name: "a"},
			operand{name: "+"},
			operand{name: "b"},
		}},
		arg: &Expression{lexemas: []lexema{
			operand{name: "c"},
			index{high: &Expression{lexemas: []lexema{operand{name: "2"}}}},
			operand{name: "+"},
			operand{name: "b"},
			index{low: &Expression{lexemas: []lexema{operand{name: "2"}}}},
		}}}
	expectLatex = "\\sqrt[a + b]{c ^2 + b _2}"
	check(t, latex, expectR, expectLatex, expectShift)
}
