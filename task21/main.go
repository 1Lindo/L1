/*
Реализовать паттерн «адаптер» на любом примере.
*/

/*
Паттерн "Адаптер" позволяет использовать интерфейс объекта в другом контексте, не изменяя его исходного кода.
Для реализации паттерна необходимо создать адаптер, который будет обеспечивать связь между двумя интерфейсами.

Создам структуру Square с методом Area() для вычисления площади квадрата и
интерфейс Shape с методом Area() для вычисления площади любой фигуры.

Мне нужно использовать объект Square в контексте, где требуется объект типа Shape.
Для этого создаю адаптер SquareAdapter, который руализует интерфейс Shape и
использует метод Area() объекта Square:

В данном случае созданная структуру SquareAdapter содержит указатель на объект Square, а метод Area()
структуры SquareAdapter вызывает метод Area() объекта Square.

Использую объект Square в контексте, где требуется объект типа Shape через адаптер SquareAdapter:

Для этого создал объект Square с длиной стороны 5 и создали адаптер SquareAdapter, который содержит этот объект.
Вызвал метод Area() как для объекта Square, так и для объекта типа Shape, используя адаптер.

*/

package main

import "fmt"

type Square struct {
	Side float64
}

func (s *Square) Area() float64 {
	return s.Side * s.Side
}

type Shape interface {
	Area() float64
}

type SquareAdapter struct {
	*Square
}

func (sa *SquareAdapter) Area() float64 {
	return sa.Square.Area()
}

func main() {
	square := &Square{Side: 5}
	adapter := &SquareAdapter{square}

	fmt.Println("Square area:", square.Area())
	fmt.Println("Shape area:", adapter.Area())
}
