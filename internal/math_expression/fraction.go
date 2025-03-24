package math_expression

import "strings"

type fraction struct {
	numerator   *Expression
	denumerator *Expression
}

func (f fraction) toString() string {
	var ret strings.Builder

	ret.WriteString("\\frac")
	ret.WriteByte('{')
	ret.WriteString(f.numerator.ToString())
	ret.WriteByte('}')
	ret.WriteByte('{')
	ret.WriteString(f.denumerator.ToString())
	ret.WriteByte('}')
	return ret.String()
}

func (f fraction) equals(other lexema) bool {

	conv, ok := other.(fraction)
	if !ok {
		return false
	}
	return f.numerator.Equals(conv.numerator) && f.denumerator.Equals(conv.denumerator)
}

func parseFraction(str string) (lexema, int) {
	idx := 0
	ret := fraction{numerator: &Expression{make([]lexema, 0)},
		denumerator: &Expression{make([]lexema, 0)}}
	for i := 0; i < 2; i++ {
		for ; idx < len(str) && (str[idx] == ' ' || str[idx] == '\n'); idx++ {

		}
		if idx == len(str) || str[idx] != '{' {
			break
		}
		idx++
		start := idx
		countOpen := 1
		for ; idx < len(str) && countOpen > 0; idx++ {
			if str[idx] == '{' {
				countOpen++
			}
			if str[idx] == '}' {
				countOpen--
			}
		}
		var e *Expression
		if countOpen > 0 {
			e = ParseLaTeX(str[start:idx])
		} else {
			e = ParseLaTeX(str[start : idx-1])
		}
		if i == 0 {
			ret.numerator = e
		} else {
			ret.denumerator = e
		}
	}
	return ret, idx
}

func (f fraction) hasPrefix(other lexema) bool {
	o, ok := other.(operand)
	opFrac := operand{name: "\\frac"}
	if ok && opFrac.hasPrefix(o) {
		return true
	}
	conv, ok := other.(fraction)
	if !ok {
		return false
	}
	return f.numerator.HasPrefix(conv.numerator) && f.denumerator.len() == 0 ||
		f.numerator.Equals(conv.numerator) && f.numerator.HasPrefix(conv.denumerator)
}

func (f fraction) mask(other lexema) string {
	conv, ok := other.(fraction)

	var ret strings.Builder

	ret.WriteString("\\frac")

	ret.WriteByte('{')
	if ok {
		ret.WriteString(f.numerator.Mask(conv.numerator))
	} else {
		ret.WriteString(f.numerator.Mask(&Expression{make([]lexema, 0)}))
	}
	ret.WriteByte('}')

	ret.WriteByte('{')
	if ok {
		ret.WriteString(f.denumerator.Mask(conv.denumerator))
	} else {
		ret.WriteString(f.denumerator.Mask(&Expression{make([]lexema, 0)}))
	}
	ret.WriteByte('}')
	return ret.String()
}
