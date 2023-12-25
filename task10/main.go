/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

*/

/*
Создаю слайс temps с температурными колебаниями и пустую мапу groups,
в которой будут храниться группы температурных значений.

Итерируюсь по каждому значению в слайсе temps и вычисляю группу для каждого значения (округляем до ближайшего десятка),
затем добавляю это значение в соответствующую группу в мапе groups.

Итерируюсь по всем группам в мапе groups и вывожу их с помощью fmt.Printf.

Результатом работы будет вывод на экран значений как в примере.

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} //Слайс
	groups := make(map[int][]float64)                                     //Мапа

	for _, temp := range temps {
		group := int(math.Floor(temp/10.0)) * 10
		groups[group] = append(groups[group], temp)
	} //Итерирование по объектам в салайсе и округление до 10

	for group, values := range groups {
		fmt.Printf("%d: %v\n", group, values)
	} //Вывод набора значений из мапы groups
}
