package math_expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpression_ToString_Simple(t *testing.T) {
	var (
		e              *Expression
		expect, actual string
	)

	expect = ""
	actual = e.ToString()
	assert.Equal(t, expect, actual)

	e = &Expression{}
	expect = ""
	actual = e.ToString()
	assert.Equal(t, expect, actual)

	e = &Expression{
		lexemas: make([]lexema, 0),
	}
	expect = ""
	actual = e.ToString()
	assert.Equal(t, expect, actual)

	e = &Expression{
		lexemas: []lexema{operand{name: ""}},
	}
	expect = ""
	actual = e.ToString()
	assert.Equal(t, expect, actual)

	e = &Expression{
		lexemas: []lexema{operand{name: "a"}},
	}
	expect = "a"
	actual = e.ToString()
	assert.Equal(t, expect, actual)

	e = &Expression{
		lexemas: []lexema{
			operand{name: "a"},
			operand{name: "a"}},
	}
	expect = "a a"
	actual = e.ToString()
	assert.Equal(t, expect, actual)

	e = &Expression{
		lexemas: []lexema{
			operand{name: "\\alpha"},
			operand{name: "b"},
			operand{name: "\\beta"},
		},
	}
	expect = "\\alpha b \\beta"
	actual = e.ToString()
	assert.Equal(t, expect, actual)

	e = &Expression{
		lexemas: []lexema{
			operand{name: "\\alpha"},
			operand{name: "+"},
			operand{name: "\\beta"},
		},
	}
	expect = "\\alpha + \\beta"
	actual = e.ToString()
	assert.Equal(t, expect, actual)

	e = &Expression{
		lexemas: []lexema{
			operand{name: "a"},
			operand{name: "+"},
			operand{name: "b"},
			operand{name: "="},
			operand{name: "c"},
		},
	}
	expect = "a + b = c"
	actual = e.ToString()
	assert.Equal(t, expect, actual)
}

func TestExpression_Equals_Simple(t *testing.T) {
	var (
		e1, e2 *Expression
	)

	assert.True(t, e1.Equals(e2))

	e2 = &Expression{}
	assert.True(t, e1.Equals(e2))
	assert.True(t, e2.Equals(e1))

	e1 = &Expression{lexemas: []lexema{
		operand{name: "a"},
	}}
	assert.False(t, e1.Equals(e2))
	assert.False(t, e2.Equals(e1))

	e2 = &Expression{lexemas: []lexema{
		operand{name: "a"},
	}}
	assert.True(t, e1.Equals(e2))
	assert.True(t, e2.Equals(e1))

	e1 = &Expression{
		lexemas: []lexema{
			operand{name: "a"},
			operand{name: "b"},
		},
	}
	e2 = &Expression{
		lexemas: []lexema{
			operand{name: "a"},
			operand{name: "c"},
		},
	}
	assert.False(t, e1.Equals(e2))
	assert.False(t, e2.Equals(e1))

	e1 = &Expression{
		lexemas: []lexema{
			operand{name: "a"},
			operand{name: "b"},
		},
	}
	e2 = &Expression{
		lexemas: []lexema{
			operand{name: "a"},
			operand{name: "b"},
		},
	}
	assert.True(t, e1.Equals(e2))
	assert.True(t, e2.Equals(e1))

	e1 = &Expression{
		lexemas: []lexema{
			operand{name: "a"},
			operand{name: "b"},
			operand{name: "c"},
		},
	}
	e2 = &Expression{
		lexemas: []lexema{
			operand{name: "a"},
			operand{name: "b"},
		},
	}
	assert.False(t, e1.Equals(e2))
	assert.False(t, e2.Equals(e1))

	e1 = &Expression{
		lexemas: []lexema{
			operand{name: "\\alpha"},
			operand{name: "b"},
		},
	}
	e2 = &Expression{
		lexemas: []lexema{
			operand{name: "\\alpha"},
			operand{name: "b"},
		},
	}
	assert.True(t, e1.Equals(e2))
	assert.True(t, e2.Equals(e1))

	e1 = &Expression{
		lexemas: []lexema{
			operand{name: "\\alpha"},
			operand{name: "\\beta"},
		},
	}
	e2 = &Expression{
		lexemas: []lexema{
			operand{name: "\\alpha"},
			operand{name: "\\beta"},
		},
	}
	assert.True(t, e1.Equals(e2))
	assert.True(t, e2.Equals(e1))
}

func TestExpression_HasPrefix_Simple(t *testing.T) {
	var (
		e1, e2 *Expression
	)

	assert.True(t, e1.HasPrefix(e2))

	e2 = &Expression{}
	assert.True(t, e1.HasPrefix(e2))
	assert.True(t, e2.HasPrefix(e1))

	e1 = &Expression{lexemas: []lexema{
		operand{name: "a"},
	}}
	assert.True(t, e1.HasPrefix(e2))
	assert.False(t, e2.HasPrefix(e1))

	e2 = &Expression{lexemas: []lexema{
		operand{name: "a"},
	}}
	assert.True(t, e1.HasPrefix(e2))
	assert.True(t, e2.HasPrefix(e1))

	e1 = &Expression{
		lexemas: []lexema{
			operand{name: "a"},
			operand{name: "b"},
		},
	}
	e2 = &Expression{
		lexemas: []lexema{
			operand{name: "a"},
			operand{name: "c"},
		},
	}
	assert.False(t, e1.HasPrefix(e2))
	assert.False(t, e2.HasPrefix(e1))

	e1 = &Expression{
		lexemas: []lexema{
			operand{name: "a"},
			operand{name: "b"},
		},
	}
	e2 = &Expression{
		lexemas: []lexema{
			operand{name: "a"},
			operand{name: "b"},
		},
	}
	assert.True(t, e1.HasPrefix(e2))
	assert.True(t, e2.HasPrefix(e1))

	e1 = &Expression{
		lexemas: []lexema{
			operand{name: "a"},
			operand{name: "b"},
			operand{name: "c"},
		},
	}
	e2 = &Expression{
		lexemas: []lexema{
			operand{name: "a"},
			operand{name: "b"},
		},
	}
	assert.True(t, e1.HasPrefix(e2))
	assert.False(t, e2.HasPrefix(e1))

	e1 = &Expression{
		lexemas: []lexema{
			operand{name: "\\alpha"},
			operand{name: "b"},
		},
	}
	e2 = &Expression{
		lexemas: []lexema{
			operand{name: "\\alpha"},
			operand{name: "b"},
		},
	}
	assert.True(t, e1.HasPrefix(e2))
	assert.True(t, e2.HasPrefix(e1))

	e1 = &Expression{
		lexemas: []lexema{
			operand{name: "\\alpha"},
			operand{name: "\\beta"},
		},
	}
	e2 = &Expression{
		lexemas: []lexema{
			operand{name: "\\alpha"},
			operand{name: "\\beta"},
		},
	}
	assert.True(t, e1.HasPrefix(e2))
	assert.True(t, e2.HasPrefix(e1))

	e1 = &Expression{
		lexemas: []lexema{
			operand{name: "\\alpha"},
			operand{name: "\\beta"},
		},
	}
	e2 = &Expression{
		lexemas: []lexema{
			operand{name: "\\alpha"},
			operand{name: "\\be"},
		},
	}
	assert.True(t, e1.HasPrefix(e2))
	assert.False(t, e2.HasPrefix(e1))
}

func TestExpression_ParseLaTeX_Simple(t *testing.T) {
	var (
		actualExpr, expectExpr *Expression
		expectStr, actualStr   string
		latex                  string
	)

	latex = "a"
	expectExpr = &Expression{lexemas: []lexema{
		operand{name: "a"},
	}}
	expectStr = "a"
	actualExpr = ParseLaTeX(latex)
	actualStr = actualExpr.ToString()
	assert.True(t, expectExpr.Equals(actualExpr))
	assert.True(t, actualExpr.Equals(expectExpr))
	assert.Equal(t, expectStr, actualStr)

	latex = "ab"
	expectExpr = &Expression{lexemas: []lexema{
		operand{name: "a"},
		operand{name: "b"},
	}}
	expectStr = "a b"
	actualExpr = ParseLaTeX(latex)
	actualStr = actualExpr.ToString()
	assert.True(t, expectExpr.Equals(actualExpr))
	assert.True(t, actualExpr.Equals(expectExpr))
	assert.Equal(t, expectStr, actualStr)

	latex = "a   b"
	expectExpr = &Expression{lexemas: []lexema{
		operand{name: "a"},
		operand{name: "b"},
	}}
	expectStr = "a b"
	actualExpr = ParseLaTeX(latex)
	actualStr = actualExpr.ToString()
	assert.True(t, expectExpr.Equals(actualExpr))
	assert.True(t, actualExpr.Equals(expectExpr))
	assert.Equal(t, expectStr, actualStr)

	latex = "a+b"
	expectExpr = &Expression{lexemas: []lexema{
		operand{name: "a"},
		operand{name: "+"},
		operand{name: "b"},
	}}
	expectStr = "a + b"
	actualExpr = ParseLaTeX(latex)
	actualStr = actualExpr.ToString()
	assert.True(t, expectExpr.Equals(actualExpr))
	assert.True(t, actualExpr.Equals(expectExpr))
	assert.Equal(t, expectStr, actualStr)

	latex = "\\alpha+\\beta"
	expectExpr = &Expression{lexemas: []lexema{
		operand{name: "\\alpha"},
		operand{name: "+"},
		operand{name: "\\beta"},
	}}
	expectStr = "\\alpha + \\beta"
	actualExpr = ParseLaTeX(latex)
	actualStr = actualExpr.ToString()
	assert.True(t, expectExpr.Equals(actualExpr))
	assert.True(t, actualExpr.Equals(expectExpr))
	assert.Equal(t, expectStr, actualStr)

	latex = "\\alpha \\cdot \\beta"
	expectExpr = &Expression{lexemas: []lexema{
		operand{name: "\\alpha"},
		operand{name: "\\cdot"},
		operand{name: "\\beta"},
	}}
	expectStr = "\\alpha \\cdot \\beta"
	actualExpr = ParseLaTeX(latex)
	actualStr = actualExpr.ToString()
	assert.True(t, expectExpr.Equals(actualExpr))
	assert.True(t, actualExpr.Equals(expectExpr))
	assert.Equal(t, expectStr, actualStr)
}

func TestExpression_ParseLaTeX(t *testing.T) {
	var (
		latex       string
		expectExpr  *Expression
		expectLatex string
		allExpr     *Expression
	)
	check := func(t *testing.T, latex string, expectExpr *Expression, expectLatex string, allExpr *Expression) {
		var (
			actualExpr  *Expression
			actualLatex string
		)
		actualExpr = ParseLaTeX(latex)
		actualLatex = actualExpr.ToString()
		assert.True(t, actualExpr.Equals(expectExpr))
		assert.True(t, expectExpr.Equals(actualExpr))
		assert.Equal(t, expectLatex, actualLatex)
		assert.True(t, allExpr.HasPrefix(actualExpr))
	}

	latex = ""
	expectExpr = &Expression{make([]lexema, 0)}
	expectLatex = ""
	allExpr = &Expression{make([]lexema, 0)}
	check(t, latex, expectExpr, expectLatex, allExpr)

	allExpr = &Expression{lexemas: []lexema{
		operand{name: "|"},
		operand{name: "x"},
		operand{name: "+"},
		operand{name: "1"},
		operand{name: "|"},
		operand{name: "<"},
		operand{name: "0"},
		operand{name: ","},
		operand{name: "0"},
		operand{name: "1"},
	}}

	latex = "|"
	expectExpr = &Expression{lexemas: []lexema{
		operand{name: "|"},
	}}
	expectLatex = "|"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "|x"
	expectExpr.lexemas = append(expectExpr.lexemas, operand{name: "x"})
	expectLatex = "| x"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "|x+"
	expectExpr.lexemas = append(expectExpr.lexemas, operand{name: "+"})
	expectLatex = "| x +"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "|x+  1"
	expectExpr.lexemas = append(expectExpr.lexemas, operand{name: "1"})
	expectLatex = "| x + 1"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "|x+  1|"
	expectExpr.lexemas = append(expectExpr.lexemas, operand{name: "|"})
	expectLatex = "| x + 1 |"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "|x+  1|<"
	expectExpr.lexemas = append(expectExpr.lexemas, operand{name: "<"})
	expectLatex = "| x + 1 | <"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "|x+  1|<0"
	expectExpr.lexemas = append(expectExpr.lexemas, operand{name: "0"})
	expectLatex = "| x + 1 | < 0"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "|x+  1|<0,"
	expectExpr.lexemas = append(expectExpr.lexemas, operand{name: ","})
	expectLatex = "| x + 1 | < 0 ,"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "|x+  1|<0,0"
	expectExpr.lexemas = append(expectExpr.lexemas, operand{name: "0"})
	expectLatex = "| x + 1 | < 0 , 0"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "|x+  1|<0,01"
	expectExpr.lexemas = append(expectExpr.lexemas, operand{name: "1"})
	expectLatex = "| x + 1 | < 0 , 0 1"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "|x+  1|<0,01     "
	expectLatex = "| x + 1 | < 0 , 0 1"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "|x+  1|<0,01     "
	expectLatex = "| x + 1 | < 0 , 0 1"
	check(t, latex, expectExpr, expectLatex, allExpr)

	allExpr = &Expression{lexemas: []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{lexemas: []lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			},
			},
			denumerator: &Expression{lexemas: []lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{high: &Expression{lexemas: []lexema{operand{name: "2"}}}},
		operand{name: "+"},
		operand{name: "("},
		fraction{
			numerator: &Expression{lexemas: []lexema{
				operand{name: "x"},
				operand{name: "-"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			},
			},
			denumerator: &Expression{lexemas: []lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{high: &Expression{lexemas: []lexema{operand{name: "2"}}}},
		operand{name: "="},
		operand{name: "x"},
		index{high: &Expression{lexemas: []lexema{operand{name: "2"}}}},
	}}
	expectExpr.lexemas = make([]lexema, 0)

	latex = "("
	expectExpr.lexemas = append(expectExpr.lexemas, operand{name: "("})
	expectLatex = "("
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\"
	expectExpr.lexemas = append(expectExpr.lexemas, operand{name: "\\"})
	expectLatex = "( \\"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\f"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		operand{name: "\\f"},
	}
	expectLatex = "( \\f"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\fr"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		operand{name: "\\fr"},
	}
	expectLatex = "( \\fr"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\fra"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		operand{name: "\\fra"},
	}
	expectLatex = "( \\fra"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator:   &Expression{make([]lexema, 0)},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator:   &Expression{make([]lexema, 0)},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
			}},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
			}},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x +}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
			}},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x + |}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
			}},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x + | x}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x + | x |}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x + | x |}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x + | x |}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2}"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
	}
	expectLatex = "( \\frac{x + | x |}{2} )"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{make([]lexema, 0)},
			low:  &Expression{nil},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 +"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+("
	expectExpr.lexemas = []lexema{

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},
		operand{name: "("},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ("
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		operand{name: "\\"},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\f"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		operand{name: "\\f"},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\f"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\fr"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		operand{name: "\\fr"},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\fr"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\fra"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		operand{name: "\\fra"},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\fra"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator:   &Expression{make([]lexema, 0)},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator:   &Expression{make([]lexema, 0)},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
			}},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x-"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "-"},
			}},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x -}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x-|"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "-"},
				operand{name: "|"},
			}},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x - |}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x-|x"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "-"},
				operand{name: "|"},
				operand{name: "x"},
			}},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x - | x}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x-|x  |"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "-"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x - | x |}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x-|x  |}"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "-"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x - | x |}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x-|x  |}{"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "-"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x - | x |}{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x-|x  |}{2"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "-"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x - | x |}{2}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x-|x  |}{2}"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "-"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x - | x |}{2}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x-|x  |}{2})"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "-"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x - | x |}{2} )"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x-|x  |}{2})^"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "-"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{make([]lexema, 0)},
			low:  &Expression{nil},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x - | x |}{2} ) ^{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x-|x  |}{2})^2"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "-"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x - | x |}{2} ) ^2"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x-|x  |}{2})^2="
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "-"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "="},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x - | x |}{2} ) ^2 ="
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x-|x  |}{2})^2=x"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "-"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "="},
		operand{name: "x"},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x - | x |}{2} ) ^2 = x"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x-|x  |}{2})^2=x^"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "-"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "="},
		operand{name: "x"},
		index{
			high: &Expression{lexemas: []lexema{}},
			low:  &Expression{nil},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x - | x |}{2} ) ^2 = x ^{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "(\\frac{x+|x  |}{2})^2+(\\frac{x-|x  |}{2})^2=x^{2}"
	expectExpr.lexemas = []lexema{
		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "+"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "+"},

		operand{name: "("},
		fraction{
			numerator: &Expression{[]lexema{
				operand{name: "x"},
				operand{name: "-"},
				operand{name: "|"},
				operand{name: "x"},
				operand{name: "|"},
			}},
			denumerator: &Expression{[]lexema{
				operand{name: "2"},
			}},
		},
		operand{name: ")"},
		index{
			high: &Expression{lexemas: []lexema{operand{"2"}}},
			low:  &Expression{nil},
		},
		operand{name: "="},
		operand{name: "x"},
		index{
			high: &Expression{lexemas: []lexema{operand{name: "2"}}},
			low:  &Expression{nil},
		},
	}
	expectLatex = "( \\frac{x + | x |}{2} ) ^2 + ( \\frac{x - | x |}{2} ) ^2 = x ^2"
	check(t, latex, expectExpr, expectLatex, allExpr)

	allExpr = &Expression{lexemas: []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{lexemas: []lexema{
				operand{name: "\\sin"},
				operand{name: "("},
				root{
					degree: nil,
					arg: &Expression{lexemas: []lexema{
						operand{name: "x"},
					}},
				},
				operand{name: ")"},
			}},
		},
	}}

	latex = "y"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
	}
	expectLatex = "y"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y="
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
	}
	expectLatex = "y ="
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		operand{name: "\\"},
	}
	expectLatex = "y = \\"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\s"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		operand{name: "\\s"},
	}
	expectLatex = "y = \\s"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sq"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		operand{name: "\\sq"},
	}
	expectLatex = "y = \\sq"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqr"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		operand{name: "\\sqr"},
	}
	expectLatex = "y = \\sqr"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg:    &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "y = \\sqrt{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg:    &Expression{make([]lexema, 0)},
		},
	}
	expectLatex = "y = \\sqrt{}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{\\"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{[]lexema{
				operand{name: "\\"},
			}},
		},
	}
	expectLatex = "y = \\sqrt{\\}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{\\s"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{[]lexema{
				operand{name: "\\s"},
			}},
		},
	}
	expectLatex = "y = \\sqrt{\\s}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{\\si"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{[]lexema{
				operand{name: "\\si"},
			}},
		},
	}
	expectLatex = "y = \\sqrt{\\si}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{\\sin"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{[]lexema{
				operand{name: "\\sin"},
			}},
		},
	}
	expectLatex = "y = \\sqrt{\\sin}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{\\sin("
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{[]lexema{
				operand{name: "\\sin"},
				operand{name: "("},
			}},
		},
	}
	expectLatex = "y = \\sqrt{\\sin (}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{\\sin(\\"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{[]lexema{
				operand{name: "\\sin"},
				operand{name: "("},
				operand{name: "\\"},
			}},
		},
	}
	expectLatex = "y = \\sqrt{\\sin ( \\}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{\\sin(\\s"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{[]lexema{
				operand{name: "\\sin"},
				operand{name: "("},
				operand{name: "\\s"},
			}},
		},
	}
	expectLatex = "y = \\sqrt{\\sin ( \\s}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{\\sin(\\sq"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{[]lexema{
				operand{name: "\\sin"},
				operand{name: "("},
				operand{name: "\\sq"},
			}},
		},
	}
	expectLatex = "y = \\sqrt{\\sin ( \\sq}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{\\sin(\\sqr"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{[]lexema{
				operand{name: "\\sin"},
				operand{name: "("},
				operand{name: "\\sqr"},
			}},
		},
	}
	expectLatex = "y = \\sqrt{\\sin ( \\sqr}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{\\sin(\\sqrt"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{[]lexema{
				operand{name: "\\sin"},
				operand{name: "("},
				root{
					degree: nil,
					arg:    &Expression{[]lexema{}},
				},
			}},
		},
	}
	expectLatex = "y = \\sqrt{\\sin ( \\sqrt{}}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{\\sin(\\sqrt{"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{[]lexema{
				operand{name: "\\sin"},
				operand{name: "("},
				root{
					degree: nil,
					arg:    &Expression{[]lexema{}},
				},
			}},
		},
	}
	expectLatex = "y = \\sqrt{\\sin ( \\sqrt{}}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{\\sin(\\sqrt{x"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{[]lexema{
				operand{name: "\\sin"},
				operand{name: "("},
				root{
					degree: nil,
					arg: &Expression{[]lexema{
						operand{name: "x"},
					}},
				},
			}},
		},
	}
	expectLatex = "y = \\sqrt{\\sin ( \\sqrt{x}}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{\\sin(\\sqrt{x}"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{[]lexema{
				operand{name: "\\sin"},
				operand{name: "("},
				root{
					degree: nil,
					arg: &Expression{[]lexema{
						operand{name: "x"},
					}},
				},
			}},
		},
	}
	expectLatex = "y = \\sqrt{\\sin ( \\sqrt{x}}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{\\sin(\\sqrt{x}  )"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{[]lexema{
				operand{name: "\\sin"},
				operand{name: "("},
				root{

					degree: nil,
					arg: &Expression{[]lexema{
						operand{name: "x"},
					}},
				},
				operand{name: ")"},
			}},
		},
	}
	expectLatex = "y = \\sqrt{\\sin ( \\sqrt{x} )}"
	check(t, latex, expectExpr, expectLatex, allExpr)

	latex = "y= \\sqrt{\\sin(\\sqrt{x}  )}"
	expectExpr.lexemas = []lexema{
		operand{name: "y"},
		operand{name: "="},
		root{
			degree: nil,
			arg: &Expression{[]lexema{
				operand{name: "\\sin"},
				operand{name: "("},
				root{

					degree: nil,
					arg: &Expression{[]lexema{
						operand{name: "x"},
					}},
				},
				operand{name: ")"},
			}},
		},
	}
	expectLatex = "y = \\sqrt{\\sin ( \\sqrt{x} )}"
	check(t, latex, expectExpr, expectLatex, allExpr)
}

func TestExpression_HasPrefix(t *testing.T) {
	var (
		latex, latexPrefix string
		expr, exprPrefix   *Expression
	)

	latex = "y = \\sin x"
	latexPrefix = "y = \\sin xa"
	expr, exprPrefix = ParseLaTeX(latex), ParseLaTeX(latexPrefix)
	assert.False(t, expr.HasPrefix(exprPrefix))
}
