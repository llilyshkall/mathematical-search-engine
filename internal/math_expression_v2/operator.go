package math_expression_v2

import (
	"sort"
	"strings"
)

type operator struct {
	name string
}

// toLaTeX переводит операнд в формат LaTeX
func (o operator) toLaTeX() string {
	return o.name
}

// compare сравнивает два операнда
func (o operator) compare(other lexema) CompareResult {
	op := operand{name: o.name}
	return op.compare(other)
	//op, ok := other.(operand)
	//if ok {
	//	return operand{name: o.name}.compare(op)
	//}
	//conv, ok := other.(operator)
	//if !ok {
	//	return Different
	//}
	//cmp := strings.HasPrefix(o.name, conv.name)
	//if !cmp {
	//	return Different
	//}
	//if len(conv.name) == len(o.name) {
	//	return Equal
	//}
	//return Prefix
}

// maskLaTeX строит маску по операнду other
func (o operator) maskLaTeX(other lexema) (string, error) {
	op := operand{name: o.name}
	return op.maskLaTeX(other)
	//cmp := o.compare(other)
	//if cmp == Different {
	//	return "", &NotPrefixError{
	//		lexema1: o,
	//		lexema2: other,
	//	}
	//}
	//if cmp == Equal {
	//	return o.toLaTeX(), nil
	//}
	//return "\\color{gray}" + o.name + "\\color{black}", nil
}

// parseOperator достает из строки первую лексему, пропуская все пробельные символы и проверяет:
// содержится ли она в списке операторов.
// В случае достижения конца строки возвращает nil и длина этой строки
func parseOperator(str string) (lexema, int) {
	start := 0
	for ; start < len(str) && str[start] == ' '; start++ {
	}
	if start == len(str) {
		return nil, start
	}
	i := sort.Search(len(operators), func(i int) bool {
		return strings.Compare(operators[i], str[start:min(len(str), start+len(operators[i]))]) >= 0
	})
	if i < len(operators) && strings.Compare(operators[i], str[start:min(len(str), start+len(operators[i]))]) == 0 {
		return operator{name: operators[i]}, start + len(operators[i])
	}
	return nil, 0
}
