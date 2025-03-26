package math_expression_v2

import (
	"container/list"
	"errors"
	"strings"
)

// expression структура для хранения выражений
type Expression struct {
	lexemas  list.List
	operands *[]operand
}

func (e *Expression) Len() int {
	if e == nil {
		return 0
	}
	return e.lexemas.Len()
}

func (e *Expression) ToLaTeX() string {
	var ret strings.Builder

	for i, elem := e.lexemas.Len(), e.lexemas.Front(); i > 0; i, elem = i-1, elem.Next() {
		if lex, ok := elem.Value.(lexema); ok {
			ret.WriteString(lex.toLaTeX())
			ret.WriteByte(' ')
		}

		if idx, ok := elem.Value.(int); ok && idx < len(*e.operands) {
			ret.WriteString((*e.operands)[idx].toLaTeX())
		}
	}

	return ret.String()[:len(ret.String())-1]
}

func (e *Expression) Compare(other *Expression) CompareResult {
	i := min(e.Len(), other.Len())
	elem1, elem2 := e.lexemas.Front(), other.lexemas.Front()
	var cmp CompareResult

	for ; i > 0; i, elem1, elem2 = i-1, elem1.Next(), elem2.Next() {
		if cmp = e.compare2Elem(elem1, elem2); cmp != Equal {
			break
		}
	}
	if i == 0 {
		if e.Len() == other.Len() {
			return Equal
		}
		if other.Len() < e.Len() {
			return Prefix
		}
		return Different
	}
	if i == 1 && cmp == Prefix {
		return Prefix
	}
	return Different
}

func (e *Expression) compare2Elem(elem1, elem2 *list.Element) CompareResult {
	if e == nil {
		return Different
	}
	var lex1, lex2 lexema
	var ok1, ok2 bool
	var cmp CompareResult = Different

	lex1, ok1 = elem1.Value.(lexema)
	lex2, ok2 = elem2.Value.(lexema)

	if !ok1 && !ok2 && elem1.Value.(int) == elem2.Value.(int) {
		cmp = Equal
	} else {
		if !ok1 {
			lex1 = (*e.operands)[elem1.Value.(int)]
		}
		if !ok2 {
			lex2 = (*e.operands)[elem2.Value.(int)]
		}
		if !ok1 && ok2 {
			cmp = Different
		} else {
			cmp = lex1.compare(lex2)
		}
	}
	return cmp
}

// TODO
func (e *Expression) MaskLaTeX(other *Expression) (string, error) {
	if e == nil {
		return "", errors.New("empty lexema")
	}
	if other.Len() > e.Len() {
		return "", &NotPrefixError{
			lexema1: nil,
			lexema2: nil,
		}
	}

	ret := strings.Builder{}
	i := other.Len()
	elem1, elem2 := e.lexemas.Front(), e.lexemas.Front()
	var cmp CompareResult

	for ; i > 0; i, elem1, elem2 = i-1, elem1.Next(), elem2.Next() {
		if cmp = e.compare2Elem(elem1, elem2); cmp != Equal {
			break
		}
		lex, ok := elem2.Value.(lexema)
		if !ok {
			lex = (*e.operands)[elem2.Value.(int)]
		}
		ret.WriteString(lex.toLaTeX())
		ret.WriteString(" ")
	}

	if i == 1 && cmp == Prefix {
		lex1, ok1 := elem1.Value.(lexema)
		if !ok1 {
			lex1 = (*e.operands)[elem1.Value.(int)]
		}
		lex2, ok2 := elem2.Value.(lexema)
		if !ok2 {
			lex2 = (*e.operands)[elem2.Value.(int)]
		}
		var s string
		if ok1 && ok2 {
			s, _ = lex1.maskLaTeX(lex2)
		} else {
			s, _ = lex1.maskLaTeX(operand{name: ""})
		}

		ret.WriteString(s)
		i--
		elem1 = elem1.Next()
		elem2 = elem2.Next()
	}

	if i == 0 {
		if e.Len() == other.Len() {
			return ret.String()[:len(ret.String())-1], nil
		}

		ret.WriteString("\\color{gray} ")
		for i = e.Len() - other.Len(); i > 0; i, elem1 = i-1, elem1.Next() {
			var ok bool
			var lex lexema

			if lex, ok = elem1.Value.(lexema); !ok {
				idx := elem1.Value.(int)
				if idx >= len(*other.operands) {
					lex = (*e.operands)[idx]
				} else {
					lex = (*other.operands)[idx]
				}
			}
			ret.WriteString(lex.toLaTeX())
			ret.WriteString(" ")
		}
		ret.WriteString("\\color{black} ")
	}
	return ret.String()[:len(ret.String())-1], nil
}

func ParseLaTeX(str string) *Expression {
	operands := new([]operand)
	return parseLaTeX(str, operands)
}

func parseLaTeX(str string, operands *[]operand) *Expression {
	ret := &Expression{operands: operands}
	var shift int
	for i := 0; i < len(str); i += shift {
		if str[i] == ' ' || str[i] == '{' || str[i] == '}' {
			shift = 1
			continue
		}
		var lex lexema
		switch {
		case strings.HasPrefix(str[i:], "\\sqrt"):
			lex, shift = parseRoot(str[i+5:])
			shift += 5
		case strings.HasPrefix(str[i:], "\\frac"):
			lex, shift = parseFraction(str[i+5:])
			shift += 5
		case str[i] == '^' || str[i] == '_':
			lex, shift = parseIndex(str[i:])
		default:
			lex, shift = parseOperator(str[i:])
			if lex == nil {
				var op operand
				op, shift = parseOperand(str[i:])
				idx := 0
				for ; idx < len(*ret.operands) && (*ret.operands)[idx].compare(lex) != Equal; idx++ {
				}
				if idx == len(*ret.operands) {
					*ret.operands = append(*ret.operands, op)
				}
				ret.lexemas.PushBack(idx)
			}
		}
		if lex != nil {
			ret.lexemas.PushBack(lex)
		}
	}
	return ret
}
