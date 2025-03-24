package math_expression

type lexema interface {
	toString() string
	equals(lexema) bool
	hasPrefix(lexema) bool
	mask(lexema) string
}
