package math_expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperand_Equals(t *testing.T) {
	var (
		o1, o2 operand
		i      index
	)

	assert.True(t, o1.equals(o2))

	o1 = operand{name: "a"}
	assert.False(t, o1.equals(o2))

	o2 = operand{name: "a"}
	assert.True(t, o1.equals(o2))

	o1 = operand{name: "b"}
	o2 = operand{name: "c"}
	assert.False(t, o1.equals(o2))

	o1 = operand{name: "\\a"}
	o2 = operand{name: "a"}
	assert.False(t, o1.equals(o2))

	o1 = operand{name: "\\a"}
	o2 = operand{name: "\\a"}
	assert.True(t, o1.equals(o2))

	o1 = operand{name: "\\alpha"}
	o2 = operand{name: "\\a"}
	assert.False(t, o1.equals(o2))

	o1 = operand{name: "\\alpha"}
	o2 = operand{name: "\\beta"}
	assert.False(t, o1.equals(o2))

	o1 = operand{name: "\\alpha"}
	o2 = operand{name: "\\alpha"}
	assert.True(t, o1.equals(o2))

	o1 = operand{name: "\\alpha"}
	assert.False(t, o1.equals(nil))

	o1 = operand{name: "\\alpha"}
	assert.False(t, o1.equals(i))
}

func TestOperand_HasPrefix(t *testing.T) {
	var (
		o1, o2 operand
		i      index
	)

	assert.True(t, o1.hasPrefix(o2))

	o1 = operand{name: "a"}
	assert.True(t, o1.hasPrefix(o2))

	o2 = operand{name: "a"}
	assert.True(t, o1.hasPrefix(o2))

	o1 = operand{name: "b"}
	o2 = operand{name: "c"}
	assert.False(t, o1.hasPrefix(o2))

	o1 = operand{name: "\\a"}
	o2 = operand{name: "a"}
	assert.False(t, o1.hasPrefix(o2))

	o1 = operand{name: "\\a"}
	o2 = operand{name: "\\a"}
	assert.True(t, o1.hasPrefix(o2))

	o1 = operand{name: "\\alpha"}
	o2 = operand{name: "\\a"}
	assert.True(t, o1.hasPrefix(o2))

	o1 = operand{name: "\\alpha"}
	o2 = operand{name: "\\beta"}
	assert.False(t, o1.hasPrefix(o2))

	o1 = operand{name: "\\alpha"}
	o2 = operand{name: "\\alpha"}
	assert.True(t, o1.hasPrefix(o2))

	o1 = operand{name: "\\alpha"}
	assert.False(t, o1.hasPrefix(nil))

	o1 = operand{name: "\\alpha"}
	assert.False(t, o1.hasPrefix(i))
}

func TestOperand_ToString(t *testing.T) {
	var (
		o              operand
		actual, expect string
	)

	expect = ""
	actual = o.toString()
	assert.Equal(t, expect, actual)

	o = operand{name: ""}
	expect = ""
	actual = o.toString()
	assert.Equal(t, expect, actual)

	o = operand{name: "\\alpha"}
	expect = "\\alpha"
	actual = o.toString()
	assert.Equal(t, expect, actual)

	o = operand{name: "a"}
	expect = "a"
	actual = o.toString()
	assert.Equal(t, expect, actual)

	o = operand{name: "+"}
	expect = "+"
	actual = o.toString()
	assert.Equal(t, expect, actual)
}

func TestOperand_ParseOperand(t *testing.T) {
	var (
		expectOp, actualOp       lexema
		expectShift, actualShift int
	)

	expectOp, expectShift = operand{}, 0
	actualOp, actualShift = parseOperand("")
	assert.True(t, expectOp.equals(actualOp))
	assert.Equal(t, expectShift, actualShift)

	expectOp, expectShift = operand{name: "a"}, 1
	actualOp, actualShift = parseOperand("a")
	assert.True(t, expectOp.equals(actualOp))
	assert.Equal(t, expectShift, actualShift)

	expectOp, expectShift = operand{name: "a"}, 1
	actualOp, actualShift = parseOperand("ab")
	assert.True(t, expectOp.equals(actualOp))
	assert.Equal(t, expectShift, actualShift)

	expectOp, expectShift = operand{name: "a"}, 1
	actualOp, actualShift = parseOperand("a b")
	assert.True(t, expectOp.equals(actualOp))
	assert.Equal(t, expectShift, actualShift)

	expectOp, expectShift = operand{name: "\\alpha"}, 6
	actualOp, actualShift = parseOperand("\\alpha")
	assert.True(t, expectOp.equals(actualOp))
	assert.Equal(t, expectShift, actualShift)

	expectOp, expectShift = operand{name: "\\alpha"}, 6
	actualOp, actualShift = parseOperand("\\alpha a")
	assert.True(t, expectOp.equals(actualOp))
	assert.Equal(t, expectShift, actualShift)

	expectOp, expectShift = operand{name: "\\beta"}, 5
	actualOp, actualShift = parseOperand("\\beta \\alpha")
	assert.True(t, expectOp.equals(actualOp))
	assert.Equal(t, expectShift, actualShift)

	expectOp, expectShift = operand{name: "\\beta"}, 6
	actualOp, actualShift = parseOperand(" \\beta \\alpha")
	assert.True(t, expectOp.equals(actualOp))
	assert.Equal(t, expectShift, actualShift)

	expectOp, expectShift = operand{name: "\\"}, 1
	actualOp, actualShift = parseOperand("\\ a")
	assert.True(t, expectOp.equals(actualOp))
	assert.Equal(t, expectShift, actualShift)
}
