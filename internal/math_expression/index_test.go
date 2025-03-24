package math_expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex_ToString_Simple(t *testing.T) {
	var (
		sHigh, sLow, expect string
	)
	check := func(t *testing.T, sHigh, sLow, expect string) {
		var (
			eHigh, eLow *Expression
			i           index
		)
		eHigh = ParseLaTeX(sHigh)
		eLow = ParseLaTeX(sLow)
		i = index{
			high: eHigh,
			low:  eLow,
		}
		actual := i.toString()
		assert.Equal(t, expect, actual)
	}

	sHigh = "2"
	sLow = ""
	expect = "_{}^2"
	check(t, sHigh, sLow, expect)

	sHigh = ""
	sLow = "2"
	expect = "_2^{}"
	check(t, sHigh, sLow, expect)

	sHigh = "2"
	sLow = "0"
	expect = "_0^2"
	check(t, sHigh, sLow, expect)

	sHigh = "2+2"
	sLow = "0"
	expect = "_0^{2 + 2}"
	check(t, sHigh, sLow, expect)

	sHigh = "2+2"
	sLow = ""
	expect = "_{}^{2 + 2}"
	check(t, sHigh, sLow, expect)

	sHigh = ""
	sLow = "2 + 2"
	expect = "_{2 + 2}^{}"
	check(t, sHigh, sLow, expect)

	sHigh = "0"
	sLow = "2 + 2"
	expect = "_{2 + 2}^0"
	check(t, sHigh, sLow, expect)

	sHigh = "\\alpha"
	sLow = "2 + 2"
	expect = "_{2 + 2}^\\alpha"
	check(t, sHigh, sLow, expect)

	sHigh = "2+2"
	sLow = "\\alpha"
	expect = "_\\alpha^{2 + 2}"
	check(t, sHigh, sLow, expect)

	sHigh = "\\alpha"
	sLow = "\\alpha"
	expect = "_\\alpha^\\alpha"
	check(t, sHigh, sLow, expect)

	sHigh = "\\beta+\\alpha"
	sLow = "abc"
	expect = "_{a b c}^{\\beta + \\alpha}"
	check(t, sHigh, sLow, expect)
}

func TestIndex_Equals_Simple(t *testing.T) {
	var (
		sHigh1, sLow1 string
		sHigh2, sLow2 string
		expect        bool
	)
	check := func(t *testing.T, sHigh1, sLow1, sHigh2, sLow2 string, expect bool) {
		var (
			eHigh1, eLow1 *Expression
			eHigh2, eLow2 *Expression
			i1, i2        index
		)
		eHigh1 = ParseLaTeX(sHigh1)
		eHigh2 = ParseLaTeX(sHigh2)
		eLow1 = ParseLaTeX(sLow1)
		eLow2 = ParseLaTeX(sLow2)
		i1 = index{high: eHigh1, low: eLow1}
		i2 = index{high: eHigh2, low: eLow2}
		if expect {
			assert.True(t, i1.equals(i2))
			assert.True(t, i2.equals(i1))
		} else {
			assert.False(t, i1.equals(i2))
			assert.False(t, i2.equals(i1))
		}
	}

	sHigh1 = ""
	sLow1 = ""
	sHigh2 = ""
	sLow2 = ""
	expect = true
	check(t, sHigh1, sLow1, sHigh2, sLow2, expect)

	sHigh1 = "a"
	sLow1 = ""
	sHigh2 = ""
	sLow2 = ""
	expect = false
	check(t, sHigh1, sLow1, sHigh2, sLow2, expect)

	sHigh1 = "a"
	sLow1 = ""
	sHigh2 = "a  b"
	sLow2 = ""
	expect = false
	check(t, sHigh1, sLow1, sHigh2, sLow2, expect)

	sHigh1 = "ab"
	sLow1 = ""
	sHigh2 = "a  b"
	sLow2 = ""
	expect = true
	check(t, sHigh1, sLow1, sHigh2, sLow2, expect)

	sHigh1 = ""
	sLow1 = "a  b"
	sHigh2 = "a  b"
	sLow2 = ""
	expect = false
	check(t, sHigh1, sLow1, sHigh2, sLow2, expect)

	sHigh1 = "a  b"
	sLow1 = "a  b"
	sHigh2 = "a  b"
	sLow2 = "ab"
	expect = true
	check(t, sHigh1, sLow1, sHigh2, sLow2, expect)

	sHigh1 = "\\alpha  b"
	sLow1 = "a  b"
	sHigh2 = "\\a  b"
	sLow2 = "ab"
	expect = false
	check(t, sHigh1, sLow1, sHigh2, sLow2, expect)

	sHigh1 = "\\alpha + \\beta"
	sLow1 = "a  b"
	sHigh2 = "\\alpha + \\beta"
	sLow2 = "ab"
	expect = true
	check(t, sHigh1, sLow1, sHigh2, sLow2, expect)

	sHigh1 = "\\alpha + \\be"
	sLow1 = "a  b"
	sHigh2 = "\\alpha + \\beta"
	sLow2 = "ab"
	expect = false
	check(t, sHigh1, sLow1, sHigh2, sLow2, expect)

	sHigh1 = "\\alpha + \\be"
	sLow1 = "a  b"
	sHigh2 = "\\alpha + \\beta + \\omega"
	sLow2 = "ab"
	expect = false
	check(t, sHigh1, sLow1, sHigh2, sLow2, expect)

	sHigh1 = ""
	sLow1 = "\\alpha + \\be"
	sHigh2 = ""
	sLow2 = "\\alpha + \\beta + \\omega"
	expect = false
	check(t, sHigh1, sLow1, sHigh2, sLow2, expect)

	sHigh1 = ""
	sLow1 = "\\alpha + \\beta +"
	sHigh2 = ""
	sLow2 = "\\alpha + \\beta + \\omega"
	expect = false
	check(t, sHigh1, sLow1, sHigh2, sLow2, expect)

	sHigh1 = ""
	sLow1 = "\\alpha + \\beta + \\omega"
	sHigh2 = ""
	sLow2 = "\\alpha + \\beta + \\omega"
	expect = true
	check(t, sHigh1, sLow1, sHigh2, sLow2, expect)

	sHigh1 = "a+b"
	sLow1 = "\\alpha + \\beta + \\omega"
	sHigh2 = "a+b"
	sLow2 = "\\alpha + \\beta + \\omega"
	expect = true
	check(t, sHigh1, sLow1, sHigh2, sLow2, expect)
}

func TestIndex_ParseIndex_Simple(t *testing.T) {
	var (
		latex       string
		expectI     index
		expectLatex string
		expectShift int
	)
	check := func(t *testing.T, latex string, expectI index, expectLatex string, expectShift int) {
		var (
			actualI     lexema
			actualLatex string
			actualShift int
		)
		actualI, actualShift = parseIndex(latex)
		actualLatex = actualI.toString()
		assert.True(t, actualI.equals(expectI))
		assert.True(t, expectI.equals(actualI))
		assert.Equal(t, expectLatex, actualLatex)
		assert.Equal(t, expectShift, actualShift)
	}

	latex = ""
	expectShift = 0
	expectLatex = ""
	expectI = index{
		high: nil,
		low:  nil,
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = ""
	expectShift = 0
	expectLatex = ""
	expectI = index{
		high: &Expression{lexemas: make([]lexema, 0)},
		low:  nil,
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = ""
	expectShift = 0
	expectLatex = ""
	expectI = index{
		high: nil,
		low:  &Expression{lexemas: make([]lexema, 0)},
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = ""
	expectShift = 0
	expectLatex = ""
	expectI = index{
		high: &Expression{lexemas: make([]lexema, 0)},
		low:  &Expression{lexemas: make([]lexema, 0)},
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = "a"
	expectShift = 0
	expectLatex = ""
	expectI = index{
		high: &Expression{lexemas: make([]lexema, 0)},
		low:  &Expression{lexemas: make([]lexema, 0)},
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = "_a"
	expectShift = 2
	expectLatex = "_a"
	expectI = index{
		high: &Expression{lexemas: nil},
		low: &Expression{lexemas: []lexema{
			operand{name: "a"},
		}},
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = "^a"
	expectShift = 2
	expectLatex = "^a"
	expectI = index{
		high: &Expression{lexemas: []lexema{
			operand{name: "a"},
		}},
		low: &Expression{lexemas: []lexema{}},
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = "_0^a"
	expectShift = 4
	expectLatex = "_0^a"
	expectI = index{
		high: &Expression{lexemas: []lexema{
			operand{name: "a"},
		}},
		low: &Expression{lexemas: []lexema{
			operand{name: "0"},
		}},
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = "_0x^a"
	expectShift = 2
	expectLatex = "_0"
	expectI = index{
		high: &Expression{lexemas: []lexema{}},
		low: &Expression{lexemas: []lexema{
			operand{name: "0"},
		}},
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = "_0^ax"
	expectShift = 4
	expectLatex = "_0^a"
	expectI = index{
		high: &Expression{lexemas: []lexema{
			operand{name: "a"},
		}},
		low: &Expression{lexemas: []lexema{
			operand{name: "0"},
		}},
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = "_0^ax"
	expectShift = 4
	expectLatex = "_0^a"
	expectI = index{
		high: &Expression{lexemas: []lexema{
			operand{name: "a"},
		}},
		low: &Expression{lexemas: []lexema{
			operand{name: "0"},
		}},
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = "_{0}^ax"
	expectShift = 6
	expectLatex = "_0^a"
	expectI = index{
		high: &Expression{lexemas: []lexema{
			operand{name: "a"},
		}},
		low: &Expression{lexemas: []lexema{
			operand{name: "0"},
		}},
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = "_0^{ax}"
	expectShift = 7
	expectLatex = "_0^{a x}"
	expectI = index{
		high: &Expression{lexemas: []lexema{
			operand{name: "a"},
			operand{name: "x"},
		}},
		low: &Expression{lexemas: []lexema{
			operand{name: "0"},
		}},
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = "_{0}^{ax}"
	expectShift = 9
	expectLatex = "_0^{a x}"
	expectI = index{
		high: &Expression{lexemas: []lexema{
			operand{name: "a"},
			operand{name: "x"},
		}},
		low: &Expression{lexemas: []lexema{
			operand{name: "0"},
		}},
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = "_{a + b} ^{ax"
	expectShift = 13
	expectLatex = "_{a + b}^{a x}"
	expectI = index{
		high: &Expression{lexemas: []lexema{
			operand{name: "a"},
			operand{name: "x"},
		}},
		low: &Expression{lexemas: []lexema{
			operand{name: "a"},
			operand{name: "+"},
			operand{name: "b"},
		}},
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = "_{a + b ax"
	expectShift = 10
	expectLatex = "_{a + b a x}"
	expectI = index{
		high: &Expression{lexemas: []lexema{}},
		low: &Expression{lexemas: []lexema{
			operand{name: "a"},
			operand{name: "+"},
			operand{name: "b"},
			operand{name: "a"},
			operand{name: "x"},
		}},
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = "_{a + b ax "
	expectShift = 11
	expectLatex = "_{a + b a x}"
	expectI = index{
		high: &Expression{lexemas: []lexema{}},
		low: &Expression{lexemas: []lexema{
			operand{name: "a"},
			operand{name: "+"},
			operand{name: "b"},
			operand{name: "a"},
			operand{name: "x"},
		}},
	}
	check(t, latex, expectI, expectLatex, expectShift)

	latex = "_{a}^b_a^{B}"
	expectShift = 12
	expectLatex = "_{a a}^{b B}"
	expectI = index{
		high: &Expression{lexemas: []lexema{
			operand{name: "b"},
			operand{name: "B"},
		}},
		low: &Expression{lexemas: []lexema{
			operand{name: "a"},
			operand{name: "a"},
		}},
	}
	check(t, latex, expectI, expectLatex, expectShift)
}
