package main

import (
	"fmt"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

const N = 5

func sender(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
}

func receiver(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	ch := make(chan int)                       // Канал для передачи данных
	timeoutChan := time.After(N * time.Second) // Таймер с кейсом времени N секунд

	// Запускаем горутину для отправки значений в канал
	go sender(ch)

	// Запускаем горутину для чтения из канала
	go receiver(ch)

	<-timeoutChan

	fmt.Println("Конец программы")

}

//--------------------------------------------------------------

func Check2() {

	// Задаем время ожидания в секундах
	timeout := 5

	// Создаем канал
	dataChan := make(chan string)

	// Запускаем горутину для записи в канал
	go func() {
		for i := 0; i < 10; i++ {
			dataChan <- fmt.Sprintf("Значение %d", i)
			time.Sleep(1 * time.Second)
		}
		close(dataChan) // Закрываем канал после отправки всех значений
	}()

	// Запускаем горутину для чтения из канала
	go func() {
		for data := range dataChan {
			fmt.Println("Получено:", data)
		}
	}()

	// Ожидаем истечения времени ожидания
	time.Sleep(time.Duration(timeout) * time.Second)

	fmt.Println("Время ожидания истекло. Завершение программы.")
}
