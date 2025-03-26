package math_expression_v2

import "fmt"

// CompareResult тип для хранения результата сравнения объектов
type CompareResult int

const (
	Different CompareResult = iota // 0: Выражения разные
	Equal                          // 1: Выражения равны
	Prefix                         // 2: Второй объект - префикс первого
)

// NotPrefixError структура для отображения ошибок связанных с префиксом
type NotPrefixError struct {
	lexema1 lexema
	lexema2 lexema
}

func (e *NotPrefixError) Error() string {
	return fmt.Sprintf("'%s' не является префиксом '%s'", e.lexema2.toLaTeX(), e.lexema1.toLaTeX())
}

// lexema является интерфейсом для определения лексемы в математическом выражении
type lexema interface {
	// toString переводит лексему в формат латех
	toLaTeX() string

	// compare сравнивает две лексемы
	compare(lexema) CompareResult

	// maskLaTeX записывает совпадающую часть лексемы черным цветом, остальную серым.
	// Используется только тогда, когда вторая лексема является префиксом первой.
	// В случае, когда вторая лексема не является префиксом первой, возвращается ошибка
	maskLaTeX(lexema) (string, error)
}
