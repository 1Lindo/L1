/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

/*
Определяю структуру Counter, которая содержит поле value для хранения значения счетчика и мьютекс mutex для
защиты доступа к этому полю в конкурентной среде (чтобы в одну "дверь" не стучались срезу две рутины).

Определяю два метода для работы со счетчиком:
метод Increment(), который инкрементирует значение счетчика;
метод Value(), который возвращает текущее значение счетчика.

В функции main() создаю экземпляр структуры Counter и запускаю 1000 горутин,
каждая из которых вызывает метод Increment() для инкрементирования значения счетчика.
Использую sync.WaitGroup для ожидания завершения всех горутин.

После завершения всех горутин вывожу итоговое значение счетчика на экран с помощью метода Value().

*/

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mutex sync.Mutex
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
} // метод инкрементирует value int

func (c *Counter) Value() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
} //метод возвращает значение value int

func main() {
	var wg sync.WaitGroup
	counter := Counter{} // определяем экземпляр структуры

	for i := 0; i < 1000; i++ {
		wg.Add(1) // добавляет в wg
		go func() {
			counter.Increment() // value ++
			wg.Done()           // убирает из wg
		}()
	}

	wg.Wait() // блокирует пока wg != 0

	fmt.Println("Final value of counter:", counter.Value())
}
