package math_expression

import "strings"

type Expression struct {
	lexemas []lexema
}

func (e *Expression) ToString() string {
	if e == nil || len(e.lexemas) == 0 {
		return ""
	}
	var ret strings.Builder

	for _, lex := range e.lexemas {
		ret.WriteString(lex.toString())
		ret.WriteByte(' ')
	}

	return ret.String()[:len(ret.String())-1]
}

func (e *Expression) len() int {
	if e == nil {
		return 0
	}
	return len(e.lexemas)
}

func (e *Expression) Equals(other *Expression) bool {
	if e.len() != other.len() {
		return false
	}
	if e.len() == 0 || other.len() == 0 {
		return e.len() == other.len()
	}
	for i := range e.lexemas {
		if !e.lexemas[i].equals(other.lexemas[i]) {
			return false
		}
	}
	return true
}

func ParseLaTeX(str string) *Expression {
	ret := &Expression{make([]lexema, 0)}
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
			lex, shift = parseOperand(str[i:])
		}
		ret.lexemas = append(ret.lexemas, lex)
	}
	return ret
}

func (e *Expression) HasPrefix(prefix *Expression) bool {
	if prefix == nil {
		return true
	}
	i := 0
	for ; i < prefix.len() && i < e.len(); i++ {
		if !e.lexemas[i].equals(prefix.lexemas[i]) {
			break
		}
	}
	if i == prefix.len() {
		return true
	}
	if i == e.len() {
		return false
	}
	if i == prefix.len()-1 {
		return e.lexemas[i].hasPrefix(prefix.lexemas[i])
	}
	return false
}

func (e *Expression) Mask(prefix *Expression) string {
	if e == nil || len(e.lexemas) == 0 {
		return ""
	}
	ret := strings.Builder{}
	i := 0
	for ; i < prefix.len() && i < e.len(); i++ {
		if e.lexemas[i].equals(prefix.lexemas[i]) {
			ret.WriteString(e.lexemas[i].toString())
			ret.WriteString(" ")
		} else {
			break
		}
	}
	if i == e.len() {
		return ret.String()
	}
	if i == prefix.len()-1 && e.lexemas[i].hasPrefix(prefix.lexemas[i]) {
		ret.WriteString(e.lexemas[i].mask(prefix.lexemas[i]))
		i++
	}
	ret.WriteString("\\color{gray}")
	for ; i < e.len(); i++ {
		ret.WriteString(e.lexemas[i].toString())
		ret.WriteString(" ")
	}
	ret.WriteString("\\color{black}")
	return ret.String()
}
