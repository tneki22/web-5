package main

import (
	"fmt"
	"sync"
	"time"
)

func work() {
	time.Sleep(time.Millisecond * 50)
	fmt.Println("done")
}

func main() {
	var wg sync.WaitGroup // Создаем WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // Увеличиваем счетчик горутин
		go func() {
			defer wg.Done() // Уменьшаем счетчик горутин после выполнения
			work()
		}()
	}

	wg.Wait() // Ожидаем, пока все горутины завершат работу
	fmt.Println("All work done!")
}
