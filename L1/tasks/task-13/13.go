package main

import (
	"fmt"
)

//Поменять местами два числа без создания временной переменной.

func swap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func main() {
	a, b := 10, 20
	fmt.Printf("До: a = %d, b = %d\n", a, b)

	a, b = swap(a, b)
	fmt.Printf("После: a = %d, b = %d\n", a, b)

	a, b = swap2(a, b)
	fmt.Printf("После второй функции: a = %d, b = %d\n", a, b)

	//Для go 1.23:
	//fmt.Println(math.swap(&a, &b))
}

func swap2(a, b int) (int, int) {

	a = a + b // a теперь содержит сумму a и b
	b = a - b // b теперь содержит исходное значение a (a - b = a)
	a = a - b // a теперь содержит исходное значение b (a - b = b)

	return a, b
}
