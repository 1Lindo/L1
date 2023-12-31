/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

/*
Использую три горутины для реализации конвейера чисел.
Первая горутина передает числа из массива в канал ch1,
вторая горутина читает числа из ch1, умножает их на 2 и передает в канал ch2,
третья горутина читает числа из ch2 и выводит их в stdout.

Использую функцию close для закрытия каналов после передачи данных, чтобы горутины, которые читают из каналов,
могли корректно завершить свою работу.
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// каналы для передачи данных между горутинами
	ch1 := make(chan int)
	ch2 := make(chan int)

	// первая горутина передает числа из массива в канал ch1
	go func() {
		for _, num := range nums {
			ch1 <- num
		}
		close(ch1)
	}()

	// вторая горутина читает числа из канала ch1, умножает их на 2 и передает в канал ch2
	go func() {
		for num := range ch1 {
			ch2 <- num * 2
		}
		close(ch2)
	}()

	// третья горутина читает числа из канала ch2 и выводит их в stdout
	for num := range ch2 {
		fmt.Println(num)
	}
}
