package math_expression

import "strings"

type root struct {
	degree *Expression
	arg    *Expression
}

func (r root) toString() string {
	var ret strings.Builder

	ret.WriteString("\\sqrt")
	if r.degree != nil {
		ret.WriteByte('[')
		ret.WriteString(r.degree.ToString())
		ret.WriteByte(']')
	}
	ret.WriteByte('{')
	ret.WriteString(r.arg.ToString())
	ret.WriteByte('}')
	return ret.String()
}

func (r root) equals(other lexema) bool {
	conv, ok := other.(root)
	if !ok {
		return false
	}
	return r.degree.Equals(conv.degree) && r.arg.Equals(conv.arg)
}

func parseRoot(str string) (lexema, int) {
	idx := 0
	ret := root{arg: &Expression{make([]lexema, 0)}}
	for ; idx < len(str) && (str[idx] == ' ' || str[idx] == '\n'); idx++ {
	}
	if idx == len(str) {
		return ret, idx
	}
	if str[idx] == '[' {
		ret.degree = &Expression{make([]lexema, 0)}
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

func (r root) hasPrefix(other lexema) bool {
	o, ok := other.(operand)
	opSqrt := operand{name: "\\sqrt"}
	if ok && opSqrt.hasPrefix(o) {
		return true
	}
	conv, ok := other.(root)
	if !ok {
		return false
	}
	return r.degree.HasPrefix(conv.degree) && r.arg.HasPrefix(conv.arg)
}

func (r root) mask(other lexema) string {
	conv, ok := other.(root)
	var ret strings.Builder

	ret.WriteString("\\sqrt")
	if r.degree != nil {
		ret.WriteByte('[')
		if ok {
			ret.WriteString(r.degree.Mask(conv.degree))
		} else {
			ret.WriteString(r.degree.Mask(&Expression{make([]lexema, 0)}))
		}
		ret.WriteByte(']')
	}
	ret.WriteByte('{')
	if ok {
		ret.WriteString(r.arg.Mask(conv.arg))
	} else {
		ret.WriteString(r.arg.Mask(&Expression{make([]lexema, 0)}))
	}
	ret.WriteByte('}')
	return ret.String()
}
