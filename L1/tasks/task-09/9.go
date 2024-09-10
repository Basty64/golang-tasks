package main

import "fmt"

//Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

func main() {
	// Создаем каналы
	inputChan := make(chan int)
	outputChan := make(chan int)

	// Запускаем горутину для записи чисел в inputChan
	go func() {
		numbers := []int{1, 2, 3, 4, 5}
		for _, num := range numbers {
			inputChan <- num
		}
		close(inputChan)
	}()

	// Запускаем горутину для обработки чисел и записи в outputChan
	go func() {
		for num := range inputChan {
			outputChan <- num * 2
		}
		close(outputChan)
	}()

	// Выводим результат из outputChan в stdout
	for num := range outputChan {
		fmt.Println(num)
	}
}

//-----------------------------------------

func Read(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func PowFromChan(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func mai() {
	ch := Read(1, 2, 3, 4, 5)
	pow2nums := PowFromChan(ch)
	for v := range pow2nums {
		fmt.Print(v, " ")
	}
}
