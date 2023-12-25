/*
Реализовать пересечение двух неупорядоченных множеств.
*/

/*
Использую функцию intersection, которая принимает два неупорядоченных множества в виде слайсов (срезов) set1 и set2 и
возвращает их пересечение в виде среза.

Для реализации пересечения использую карты (map) для хранения элементов первого множества и пересечения множеств.

Создаю карту set1Map и заполняю ее элементами из первого множества.
Затем создаю карту intersectMap и заполняю ее элементами из второго множества, если они присутствуют в set1Map.
Преобразую карту intersectMap в срез и возвращаем его как результат выполнения функции intersection.

В функции main создаю два неупорядоченных множества set1 и set2 и вызываю функцию intersection,
чтобы найти их пересечение. Результат выводится в stdout.
*/

package main

import "fmt"

func intersection(set1, set2 []int) []int {
	// создаем мапу (карту) для хранения элементов первого множества
	set1Map := make(map[int]bool)
	for _, num := range set1 {
		set1Map[num] = true
	}

	// создаем мапы (карты) для хранения пересечения множеств
	intersectMap := make(map[int]bool)
	for _, num := range set2 {
		if set1Map[num] {
			intersectMap[num] = true
		}
	}

	// преобразуем карту в массив с указанными настройками и возвращаем его
	intersect := make([]int, 0, len(intersectMap))
	for num := range intersectMap {
		intersect = append(intersect, num)
	}
	return intersect
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}
	intersect := intersection(set1, set2)
	fmt.Println(intersect)
}
