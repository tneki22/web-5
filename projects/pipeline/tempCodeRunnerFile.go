package main

import "fmt"

func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
	var prev string
	firstValue := true // Флаг для первого значения

	for val := range inputStream {
		if firstValue || val != prev {
			outputStream <- val
			prev = val
			firstValue = false
		}
	}
	close(outputStream) // Закрываем канал
}

func main() {
	// здесь должен быть код для проверки правильности работы функции removeDuplicates(in, out chan string)
	input := make(chan string)
	output := make(chan string)

	go func() {
		input <- "hello"
		input <- "world"
		input <- "world" // Дубликат
		input <- "golang"
		input <- "golang" // Дубликат
		input <- "hello"
		close(input) // Закрываем канал input
	}()

	go removeDuplicates(input, output)

	for i := range output {
		fmt.Println(i)
	}
}
