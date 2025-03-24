package math_expression

import "strings"

type operand struct {
	name string
}

func (o operand) toString() string {
	return o.name
}

func (o operand) equals(other lexema) bool {
	conv, ok := other.(operand)
	if !ok {
		return false
	}
	return o.name == conv.name
}

func parseOperand(str string) (lexema, int) {
	ret := operand{}
	start := 0
	for ; start < len(str) && str[start] == ' '; start++ {
	}
	if start == len(str) {
		return ret, start
	}
	end := start + 1
	if str[start] == '\\' {
		for ; end < len(str) && (str[end] >= 'a' && str[end] <= 'z' ||
			str[end] >= 'A' && str[end] <= 'Z'); end++ {
		}
	}
	ret.name = str[start:end]
	return ret, end
}

func (o operand) hasPrefix(other lexema) bool {
	conv, ok := other.(operand)
	if !ok {
		return false
	}
	return strings.HasPrefix(o.name, conv.name)
}

func (o operand) mask(other lexema) string {
	return "\\color{gray}" + o.name + "\\color{black}"
}
