package math_expression_v2

import "strings"

type fraction struct {
	numerator   *Expression
	denumerator *Expression
}

// toLaTeX переводит дробь в формат LaTeX
func (f fraction) toLaTeX() string {
	var ret strings.Builder

	ret.WriteString("\\frac")
	ret.WriteByte('{')
	ret.WriteString(f.numerator.ToLaTeX())
	ret.WriteByte('}')

	ret.WriteByte('{')
	ret.WriteString(f.denumerator.ToLaTeX())
	ret.WriteByte('}')

	return ret.String()
}

// compare сравнивает две дроби
func (f fraction) compare(other lexema) CompareResult {
	op, ok := other.(operand)
	if ok {
		return operand{name: "\\frac"}.compare(op)
	}
	conv, ok := other.(fraction)
	if ok {
		return Different
	}

	cmpNumerator := f.numerator.Compare(conv.numerator)
	if cmpNumerator == Different {
		return Different
	}

	cmpDenumerator := f.denumerator.Compare(conv.denumerator)
	if cmpDenumerator == Different {
		return Different
	}

	if cmpDenumerator == Equal && cmpNumerator == Equal {
		return Equal
	}
	return Prefix
}

// maskLaTeX записывает совпадающую часть лексемы черным цветом, остальную серым.
// Используется только тогда, когда вторая лексема является префиксом первой.
// В случае, когда вторая лексема не является префиксом первой, возвращается ошибка
func (f fraction) maskLaTeX(other lexema) (string, error) {
	op, ok := other.(operand)
	if ok {
		return operand{name: "\\frac"}.maskLaTeX(op)
	}
	conv, ok := other.(fraction)
	if !ok {
		return "", &NotPrefixError{lexema1: f, lexema2: other}
	}
	cmpNumerator, cmpDenumerator := f.numerator.Compare(conv.numerator), f.denumerator.Compare(conv.denumerator)
	if cmpNumerator == Different || cmpDenumerator == Different {
		return "", &NotPrefixError{lexema1: f, lexema2: conv}
	}

	var ret strings.Builder

	ret.WriteString("\\frac")
	ret.WriteString("{")
	s, _ := f.numerator.MaskLaTeX(conv.numerator)
	ret.WriteString(s)
	ret.WriteString("}")

	ret.WriteString("{")
	s, _ = f.denumerator.MaskLaTeX(conv.denumerator)
	ret.WriteString(s)
	ret.WriteString("}")

	return ret.String(), nil
}

// parseFraction достает из строки первую дробь, пропуская все пробельные символы.
// В случае достижения конца строки возвращает nil и длина этой строки
func parseFraction(str string) (lexema, int) {
	idx := 0
	ret := fraction{}
	for i := 0; i < 2; i++ {
		for ; idx < len(str) && str[idx] == ' '; idx++ {
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
