package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает,
//вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	var aStr, bStr string

	// Ввод чисел a и b
	fmt.Println("Введите число a (больше 2^20):")
	fmt.Scan(&aStr)
	fmt.Println("Введите число b (больше 2^20):")
	fmt.Scan(&bStr)

	// Преобразовываем строки в большие числа
	a := new(big.Int)
	b := new(big.Int)

	a.SetString(aStr, 10)
	b.SetString(bStr, 10)

	// Проверка условия задачи
	if a.Cmp(big.NewInt(1<<20)) <= 0 || b.Cmp(big.NewInt(1<<20)) <= 0 {
		fmt.Println("Оба числа должны быть больше 2^20.")
		return
	}

	// Выполнение арифметических операций над большими числами
	sum := new(big.Int).Add(a, b)
	difference := new(big.Int).Sub(a, b)
	product := new(big.Int).Mul(a, b)
	quotient := new(big.Rat).SetFrac(a, b) // Используем Rat для деления

	// Вывод результатов
	fmt.Printf("Сумма: %s\n", sum.String())
	fmt.Printf("Разность: %s\n", difference.String())
	fmt.Printf("Произведение: %s\n", product.String())
	fmt.Printf("Частное: %s\n", quotient.RatString())
}
