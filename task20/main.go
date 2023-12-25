/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

/*
Определяю функцию reverseWords, которая принимает строку и возвращает ее с перевернутыми словами.

Для этого:
Использую функцию strings.Split для разделения строки на слова.
Прохожу по срезу слов и меняю местами первое и последнее слово, второе и предпоследнее и т.д.
Объединяю срез слов обратно в строку с помощью функции strings.Join.
Возвращаю строку.

В функции main мы вызываю функцию reverseWords с исходной строкой и вывожу результат в stdout.
*/

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ") //разделил на слова => получился срез слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i] //меняю местами
	}
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun — sun dog snow"
	reversed := reverseWords(s)
	fmt.Println(reversed)
}
