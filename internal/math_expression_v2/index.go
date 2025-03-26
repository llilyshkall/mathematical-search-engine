package math_expression_v2

import "strings"

type index struct {
	high, low *Expression
}

// toLatex переводит индекс в формат LaTeX
func (i index) toLaTeX() string {
	var ret strings.Builder

	if i.low != nil {
		ret.WriteString("_")
		if i.low.Len() != 1 {
			ret.WriteString("{")
		}
		ret.WriteString(i.low.ToLaTeX())
		if i.low.Len() != 1 {
			ret.WriteString("}")
		}
	}

	if i.high != nil {
		ret.WriteString("^")
		if i.high.Len() != 1 {
			ret.WriteString("{")
		}
		ret.WriteString(i.high.ToLaTeX())
		if i.high.Len() != 1 {
			ret.WriteString("}")
		}
	}
	return ret.String()
}

// compare сравнивает два операнда
func (i index) compare(other lexema) CompareResult {
	conv, ok := other.(index)
	if ok {
		return Different
	}
	cmpHigh := i.high.Compare(conv.high)
	if cmpHigh == Different {
		return Different
	}
	cmpLow := i.low.Compare(conv.low)
	if cmpLow == Different {
		return Different
	}
	if cmpLow == Equal && cmpHigh == Equal {
		return Equal
	}
	return Prefix
}

// maskLaTeX записывает совпадающую часть лексемы черным цветом, остальную серым.
// Используется только тогда, когда вторая лексема является префиксом первой.
// В случае, когда вторая лексема не является префиксом первой, возвращается ошибка
func (i index) maskLaTeX(other lexema) (string, error) {
	conv, ok := other.(index)
	if !ok {
		return "", &NotPrefixError{lexema1: i, lexema2: other}
	}
	cmpLow, cmpHigh := i.low.Compare(conv.low), i.high.Compare(conv.high)
	if cmpLow == Different || cmpHigh == Different {
		return "", &NotPrefixError{lexema1: i, lexema2: conv}
	}

	var ret strings.Builder

	if i.low != nil {
		ret.WriteString("_")
		if i.low.Len() != 1 || cmpLow == Different {
			ret.WriteString("{")
		}
		s, _ := i.low.MaskLaTeX(conv.low)
		ret.WriteString(s)
		if i.low.Len() != 1 || cmpLow == Different {
			ret.WriteString("}")
		}
	}

	if i.high != nil {
		ret.WriteString("_")
		if i.low.Len() != 1 || cmpLow == Different {
			ret.WriteString("{")
		}
		s, _ := i.low.MaskLaTeX(conv.low)
		ret.WriteString(s)
		if i.low.Len() != 1 || cmpLow == Different {
			ret.WriteString("}")
		}
	}

	return ret.String(), nil
}

// parseOperand достает из строки первую лексему, пропуская все пробельные символы.
// В случае достижения конца строки возвращает nil и длина этой строки
func parseIndex(str string) (lexema, int) {
	idx := 0
	ret := index{}
	for ; idx < len(str); idx++ {
		if str[idx] == ' ' {
			continue
		}

		if str[idx] == '^' || str[idx] == '_' {
			op := str[idx]
			if op == '^' && ret.high == nil {
				ret.high = &Expression{}
			}
			if op == '_' && ret.low == nil {
				ret.low = &Expression{}
			}
			idx++
			if idx >= len(str) {
				break
			}
			e := &Expression{}
			end := idx
			if str[idx] != '{' {
				lex, shift := parseOperand(str[idx:])
				e.lexemas.PushBack(lex)
				end = idx + shift
			} else {
				countOpen := 1
				end++
				for ; end < len(str) && countOpen > 0; end++ {
					if str[end] == '{' {
						countOpen++
					}
					if str[end] == '}' {
						countOpen--
					}
				}
				if countOpen > 0 {
					e = ParseLaTeX(str[idx+1 : end])
				} else {
					e = ParseLaTeX(str[idx+1 : end-1])
				}
				idx = end - 1
			}
			if op == '^' {
				ret.high.lexemas.PushFrontList(&e.lexemas)
			} else {
				ret.low.lexemas.PushFrontList(&e.lexemas)
			}
		} else {
			break
		}
	}
	return ret, idx
}
