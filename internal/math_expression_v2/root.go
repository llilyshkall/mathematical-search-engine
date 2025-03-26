package math_expression_v2

import "strings"

type root struct {
	degree *Expression
	arg    *Expression
}

// toLatex переводит корень в формат LaTeX
func (r root) toLaTeX() string {
	var ret strings.Builder

	ret.WriteString("\\sqrt")
	if r.degree != nil {
		ret.WriteByte('[')
		ret.WriteString(r.degree.ToLaTeX())
		ret.WriteByte(']')
	}
	ret.WriteByte('{')
	ret.WriteString(r.arg.ToLaTeX())
	ret.WriteByte('}')
	return ret.String()
}

// compare сравнивает два корня
func (r root) compare(other lexema) CompareResult {
	op, ok := other.(operand)
	if ok {
		return operand{name: "\\sqrt"}.compare(op)
	}
	conv, ok := other.(root)
	if ok {
		return Different
	}

	cmpDegree := r.degree.Compare(conv.degree)
	if cmpDegree == Different {
		return Different
	}

	cmpArg := r.arg.Compare(conv.arg)
	if cmpArg == Different {
		return Different
	}

	if cmpDegree == Equal && cmpArg == Equal {
		return Equal
	}
	return Prefix
}

// maskLaTeX записывает совпадающую часть лексемы черным цветом, остальную серым.
// Используется только тогда, когда вторая лексема является префиксом первой.
// В случае, когда вторая лексема не является префиксом первой, возвращается ошибка
func (r root) maskLaTeX(other lexema) (string, error) {
	op, ok := other.(operand)
	if ok {
		return operand{name: "\\frac"}.maskLaTeX(op)
	}
	conv, ok := other.(root)
	if !ok {
		return "", &NotPrefixError{lexema1: r, lexema2: other}
	}
	cmpDegree, cmpArg := r.degree.Compare(conv.degree), r.arg.Compare(conv.arg)
	if cmpDegree == Different || cmpArg == Different {
		return "", &NotPrefixError{lexema1: r, lexema2: conv}
	}

	var ret strings.Builder

	ret.WriteString("\\sqrt")
	if r.arg.Len() != 0 {
		ret.WriteString("[")
		s, _ := r.degree.MaskLaTeX(conv.degree)
		ret.WriteString(s)
		ret.WriteString("]")
	}

	ret.WriteString("{")
	s, _ := r.arg.MaskLaTeX(conv.arg)
	ret.WriteString(s)
	ret.WriteString("}")

	return ret.String(), nil
}

// parseFraction достает из строки первый корень, пропуская все пробельные символы.
// В случае достижения конца строки возвращает nil и длина этой строки
func parseRoot(str string) (lexema, int) {
	idx := 0
	ret := root{}
	for ; idx < len(str) && str[idx] == ' '; idx++ {
	}
	if idx == len(str) {
		return ret, idx
	}
	if str[idx] == '[' {
		ret.degree = &Expression{}
		end := idx + 1
		countOpen := 1
		for ; end < len(str) && countOpen > 0; end++ {
			if str[end] == '[' {
				countOpen++
			}
			if str[end] == ']' {
				countOpen--
			}
		}
		if countOpen == 0 {
			ret.degree = ParseLaTeX(str[idx+1 : end-1])
		} else {
			ret.degree = ParseLaTeX(str[idx+1 : end])
		}
		idx = end
	}
	for ; idx < len(str) && (str[idx] == ' ' || str[idx] == '\n'); idx++ {
	}
	if idx == len(str) {
		return ret, idx
	}
	if str[idx] == '{' {
		end := idx + 1
		countOpen := 1
		for ; end < len(str) && countOpen > 0; end++ {
			if str[end] == '{' {
				countOpen++
			}
			if str[end] == '}' {
				countOpen--
			}
		}
		if countOpen == 0 {
			ret.arg = ParseLaTeX(str[idx+1 : end-1])
		} else {
			ret.arg = ParseLaTeX(str[idx+1 : end])
		}
		idx = end
	}
	return ret, idx
}
