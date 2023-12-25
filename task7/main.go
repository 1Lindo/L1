/*
Реализовать конкурентную запись данных в map.
 */


package main

import (
	"flag"
	"log"
	"sync"
)

var (
	data             map[int]int
	concurrencyLimit *int
	countWrite       *int
	mux              = sync.Mutex{}
	wg               = sync.WaitGroup{}
)

func writeToMap(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	mux.Lock()
	data[i] = i
	mux.Unlock()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	data = make(map[int]int) // мапа для записи данных

	concurrencyLimit = flag.Int("t", 5, "threads count")
	countWrite = flag.Int("w", 100, "write count")
	flag.Parse()

	sem := make(chan struct{}, *concurrencyLimit) // создал канал с ограниченной пропускной способностью
	wg.Add(*countWrite) // Задаем полное количество выполнений
	go func() {
		for i := 0; i < *countWrite; i++ {
			sem <- struct{}{} // занимает еденицу в sem (понимаем, что одна гоурутина запущена (чтобы их было лимитированное количество))
			go func(i int) {
				writeToMap(&wg, i) // пишем в мапу
				<-sem
			}(i)
		}
	}()

	wg.Wait()
	log.Println(data)
}
