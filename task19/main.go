/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/

/*
Определяю функцию reverseString, которая принимает строку и возвращает ее перевернутую версию.

Для этого:
Использую функцию utf8.RuneCountInString для определения количества символов в строке, а затем создаем срез runes из рун строки.
Прохожусь по каждой руне в строке s и добавляю ее в срез runes в обратном порядке.
Преобразую срез рун обратно в строку.
Возвращаю строку.

В функции main использую reverseString с исходной строкой и вывожу результат в stdout.
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseString(s string) string {
	runes := make([]rune, utf8.RuneCountInString(s))
	for i, r := range s {
		runes[len(runes)-1-i] = r
	}
	return string(runes)
}

func main() {
	s := "главрыба — абырвалг"
	reversed := reverseString(s)
	fmt.Println(reversed)
}
