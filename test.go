package main

import (
	"fmt"
	"time"
)
// Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!
func task2(c chan string, s string){
    for i := 0; i < 5; i++{c <- s}
}
func worker(ch chan int) {
	time.Sleep(time.Second)
	ch <- 42 // Отправляем значение в канал
}

func main() {
	ch := make(chan int)

	go worker(ch) // Запускаем горутину

	value := <-ch // Ожидаем значение из канала
	fmt.Println("Получено значение:", value)
}
