/*
Разработать программу,
которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

/*
В данном случае использую переменную типа interface{} и оператор switch для определения типа переменной.
Создаю переменную v типа interface{} и присваиваю ей значение "hello".
Использую оператор switch с ключевым словом type, чтобы определить тип переменной v.
Если тип переменной соответствует одному из кейсов switch, то будет выполнен соответствующий блок кода.

В данном случае переменная v содержит строку, поэтому программа выведет на экран сообщение "v is a string".
Если v имела значение типа int, то программа вывела бы сообщение "v is an int", и т.д.
*/

package main

import (
	"fmt"
)

func main() {
	var v interface{} = "hello"

	switch v.(type) {
	case int:
		fmt.Println("v is an int")
	case string:
		fmt.Println("v is a string")
	case bool:
		fmt.Println("v is a bool")
	case chan int:
		fmt.Println("v is a channel of ints")
	default:
		fmt.Println("v is of an unknown type")
	}
}
