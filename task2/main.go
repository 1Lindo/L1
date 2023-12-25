/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел массива [2,4,6,8,10] и
выведет их квадраты в stdout.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup     //Чтобы дождаться завершения всех горутин была реализована синхоронизация через WaitGroup
	results := make(chan int) // Канал results служит для получения результатов вычислений.

	for _, num := range numbers {
		wg.Add(1)
		//Запускается горутина для каждого числа из массива и рассчитывает квадрат этого числа
		go func(x int) {
			defer wg.Done()
			results <- x * x // Отправляет результат в канал results
		}(num)
	}

	//Запускается еще одна горутина, которая ждет завершения всех горутин WaitGroup и закрывает канал results.
	go func() {
		wg.Wait()      // Ожидание WaitGroup
		close(results) //Закрывается канал
	}()

	//Проходимся по каналу results и выводим результаты в stdout.
	for res := range results {
		fmt.Println(res)
	}
}
