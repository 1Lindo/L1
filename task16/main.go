/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

/*
Использую встроенную функцию sort.Slice().
Создаю массив arr из целых чисел и передаю его в sort.Slice().
Вторым аргументом является функция, которая определяет порядок сортировки.
В данном случае сортирую массив по возрастанию.

Функция sort.Slice() изменяет исходный массив arr и возвращает отсортированный массив.
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{5, 3, 8, 2, 1, 9, 4, 7, 6}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	}) // Изменяю исходный массив и возвращаю отсортированный

	fmt.Println(arr) //Возвращаю отсортированный массив
}