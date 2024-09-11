package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2....) с использованием конкурентных вычислений.

func main() {

	//создаём массив на 5 элементов
	a := [5]int{2, 4, 6, 8, 10}

	resultChan := make(chan int, len(a))

	var wg sync.WaitGroup

	//применяем waitgroup для синхронизации
	wg.Add(len(a))

	//Запускаем горутины
	for _, v := range a {
		go func(n int) {
			wg.Done()
			square(v, resultChan)
		}(v)
	}

	//Ждём выполнения всех горутин
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	sum := 0
	for i := 0; i < len(a); i++ {
		// Суммируем значения из канала
		sum += <-resultChan
	}
	fmt.Printf("Итоговое значение суммы: %d", sum)
}

// функция, возводящая элемент массива в квадрат
func square(n int, resultChan chan int) {
	resultChan <- n * n
}
