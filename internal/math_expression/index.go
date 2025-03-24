package math_expression

import "strings"

type index struct {
	high, low *Expression
}

func (i index) toString() string {
	var ret strings.Builder

	if i.low != nil {
		ret.WriteString("_")
		if i.low.len() != 1 {
			ret.WriteString("{")
		}
		ret.WriteString(i.low.ToString())
		if i.low.len() != 1 {
			ret.WriteString("}")
		}
	}

	if i.high != nil {
		ret.WriteString("^")
		if i.high.len() != 1 {
			ret.WriteString("{")
		}
		ret.WriteString(i.high.ToString())
		if i.high.len() != 1 {
			ret.WriteString("}")
		}
	}
	return ret.String()
}

func (i index) equals(other lexema) bool {
	conv, ok := other.(index)
	if !ok {
		return false
	}
	return i.low.Equals(conv.low) && i.high.Equals(conv.high)
}

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
				e.lexemas = append(e.lexemas, lex)
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
				ret.high.lexemas = append(ret.high.lexemas, e.lexemas...)
			} else {
				ret.low.lexemas = append(ret.low.lexemas, e.lexemas...)
			}

		} else {
			break
		}
	}
	return ret, idx
}

func (i index) hasPrefix(other lexema) bool {
	conv, ok := other.(index)
	if !ok {
		return false
	}
	return i.low.HasPrefix(conv.low) && i.high.HasPrefix(conv.high)
}

func (i index) mask(other lexema) string {
	conv, ok := other.(index)

	var ret strings.Builder

	if i.low != nil {
		ret.WriteString("_")
		if i.low.len() != 1 {
			ret.WriteString("{")
		}
		if ok {
			ret.WriteString(i.low.Mask(conv.low))
		} else {
			ret.WriteString(i.low.Mask(&Expression{make([]lexema, 0)}))
		}
		if i.low.len() != 1 {
			ret.WriteString("}")
		}
	}

	if i.high != nil {
		ret.WriteString("^")
		if i.high.len() != 1 {
			ret.WriteString("{")
		}
		if ok {
			ret.WriteString(i.high.Mask(conv.high))
		} else {
			ret.WriteString(i.high.Mask(&Expression{make([]lexema, 0)}))
		}
		if i.high.len() != 1 {
			ret.WriteString("}")
		}
	}

	return ret.String()
}
