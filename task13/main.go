/*
Поменять местами два числа без создания временной переменной.
*/

/*
Использую побитовый оператор XOR ^.
*/
package main

import "fmt"

func main() {
	a, b := 5, 10

	fmt.Println("Before swapping:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("After swapping:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
