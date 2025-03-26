package math_expression_v2

import "strings"

type operand struct {
	name string
}

// toLaTeX переводит операнд в формат LaTeX
func (o operand) toLaTeX() string {
	return o.name
}

// compare сравнивает два операнда
func (o operand) compare(other lexema) CompareResult {
	conv, ok := other.(operand)
	if !ok {
		return Different
	}
	cmp := strings.HasPrefix(o.name, conv.name)
	if !cmp {
		return Different
	}
	if len(conv.name) == len(o.name) {
		return Equal
	}
	return Prefix
}

// maskLaTeX строит маску по операнду other
func (o operand) maskLaTeX(other lexema) (string, error) {
	cmp := o.compare(other)
	if cmp == Different {
		return "", &NotPrefixError{
			lexema1: o,
			lexema2: other,
		}
	}
	if cmp == Equal {
		return o.toLaTeX(), nil
	}
	return "\\color{gray}" + o.name + "\\color{black}", nil
}

// parseOperand достает из строки первую лексему, пропуская все пробельные символы.
// В случае достижения конца строки возвращает nil и длина этой строки
func parseOperand(str string) (operand, int) {
	ret := operand{}
	start := 0
	for ; start < len(str) && str[start] == ' '; start++ {
	}
	if start == len(str) {
		return ret, start
	}
	end := start + 1
	if str[start] == '\\' {
		// TODO определить нужно ли обрабатывать промежутки
		//spaces := []string{
		//	"\\qquad", "\\quad", "\\;", "\\:", "\\,", "\\!",
		//	"\\thickspaceb", "\\medspace", "\\thinspace", "\\negthinspace", "\\negmedspace", "\\negthickspace",
		//}
		//for _, s := range spaces {
		//	if strings.HasPrefix(str[:start], s) {
		//		return nil, start + len(s)
		//	}
		//}
		for ; end < len(str) && (str[end] >= 'a' && str[end] <= 'z' ||
			str[end] >= 'A' && str[end] <= 'Z'); end++ {
		}
	}
	ret.name = str[start:end]
	return ret, end
}
