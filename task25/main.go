/*
Реализовать собственную функцию sleep.
*/

/*
Использую функцию time.After() для создания канала, который будет закрыт через указанное время.

Sleep принимает количество секунд, на которое нужно заснуть.
Внутри функции создается канал, который будет закрыт через указанное количество секунд с помощью time.After().

В функции main вызываю функцию sleep со значением 7 (программа заснет на время существоввания канала).
*/

package main

import (
	"fmt"
	"time"
)

func sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Start")
	sleep(3)
	fmt.Println("End")
}
