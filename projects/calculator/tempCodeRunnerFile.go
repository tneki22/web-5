package main

import "fmt"

// Реализация функции calculator
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int) // Канал для возвращаемого результата

	go func() {
		defer close(output) // Закрываем канал

		select {
		case val := <-firstChan:
			output <- val * val // Если получили из первого каналаквадрат
		case val := <-secondChan:
			output <- val * 3 // Если получили из второго канала умножение на 3
		case <-stopChan:
			// Если получили сигнал из stopChan, просто выходим
			return
		}
	}()

	return output
}

func main() {
	// Каналы для тестирования
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	// Запуск функции calculator
	resultChan := calculator(firstChan, secondChan, stopChan)

	// Пример использования
	go func() {

		// Отправляем значение в первый канал для получения квадрата
		// firstChan <- 4

		// Отправляем значение во второй канал для умножения на 3
		secondChan <- 5

		// Можно также отправить сигнал завершения
		// close(stopChan)
	}()

	for result := range resultChan {
		fmt.Println(result)
	}
}
