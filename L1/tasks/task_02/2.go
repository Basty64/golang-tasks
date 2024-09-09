package task_02

import (
	"fmt"
	"sort"
	"sync"
)

func Example2() {
	a := [5]int{2, 4, 6, 8, 10}

	for i := range a {
		go func() {
			a[i] = a[i] * a[i]
		}()
	}
	fmt.Println(a)
}

func square(n int, resultChan chan int) {
	resultChan <- n * n
}

func Check() {
	numbers := [5]int{2, 4, 6, 8, 10}

	fmt.Printf("Массив на входе: %d\n", numbers)
	resultChan := make(chan int, len(numbers)) // Канал для результатов
	var wg sync.WaitGroup

	wg.Add(len(numbers)) // Увеличиваем счетчик WaitGroup на количество элементов

	for _, n := range numbers {
		go func(num int) {
			defer wg.Done() // Уменьшаем счетчик WaitGroup после завершения
			square(num, resultChan)
		}(n)
	}

	go func() {
		wg.Wait()         // Ждем завершения всех goroutines
		close(resultChan) // Закрываем канал после завершения
	}()

	squaredNumbers := make([]int, len(numbers))
	for i := range squaredNumbers {
		squaredNumbers[i] = <-resultChan // Читаем результат из канала

	}
	sort.Ints(squaredNumbers)

	squaredNumbers = append(squaredNumbers)
	fmt.Println("Квадраты элементов массива:", squaredNumbers) // Вывод квадратов

}
