package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42....) с использованием конкурентных вычислений.

func main() {

	a := [5]int{2, 4, 6, 8, 10}

	resultChan := make(chan int, len(a))

	var wg sync.WaitGroup

	wg.Add(len(a))

	for _, v := range a {
		go func(n int) {
			wg.Done()
			square(v, resultChan)
		}(v)
	}

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

func square(n int, resultChan chan int) {
	resultChan <- n * n
}
