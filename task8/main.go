/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0
*/

/*
Использую операторы побитового сдвига и побитового ИЛИ (|) и И (&) для установки или сброса определенного бита.

Устанавливю i-й бит в 1, используя оператор сдвига влево (<<) для создания маски,
которая имеет единицу только в i-м бите,а затем применяю оператор ИЛИ, чтобы установить этот бит в нашем числе.

Сбрасываю i-й бит в 0, используя оператор отрицания (^) для создания маски,
которая имеет ноль только в i-м бите, а затем применяем оператор И, чтобы сбросить этот бит в нашем числе.
*/

package main

import "fmt"

func main() {
	var num int64 = 8 // исходное число
	var i uint = 2    // номер бита, который нужно установить (начиная с 0)

	// установить i-й бит в 1
	num |= 1 << i
	fmt.Printf("Установлен %d-й бит в 1: %d\n", i, num)

	// установить i-й бит в 0
	num &= ^(1 << i)
	fmt.Printf("Установлен %d-й бит в 0: %d\n", i, num)
}
