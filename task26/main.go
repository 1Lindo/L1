/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.
*/

/*
Определяю функцию UniqueChars, которая принимает строку и возвращает булево значение (если символы уникальные - true, если нет - false).

Привожу строку к нижнему регистру с помощью strings.ToLower(для регистронезависимости).

Затем мы создаем мапу chars, которая будет содержать уникальные символы из строки.
Итерируемся по символам строки.
Если символ уже присутствует в карте, то мы возвращаем false. Если символа нет в карте, то мы добавляем его в карту, а в конце выводим true.

*/

package main

import (
	"fmt"
	"strings"
)

func UniqueChars(str string) bool {
	str = strings.ToLower(str)
	chars := make(map[rune]bool)
	for _, char := range str {
		if chars[char] {
			return false
		}
		chars[char] = true
	}
	return true
}

func main() {
	str1 := "Hello, world!"
	str2 := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(UniqueChars(str1)) // false
	fmt.Println(UniqueChars(str2)) // true
}
