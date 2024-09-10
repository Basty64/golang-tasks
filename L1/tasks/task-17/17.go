package main

import (
	"fmt"
	"sort"
)

//Реализовать бинарный поиск встроенными методами языка.

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13}
	target := 9

	// Сортируем массив, если он еще не отсортирован
	sort.Ints(arr)

	// Используем метод sort.SearchInts для бинарного поиска
	index := sort.SearchInts(arr, target)

	if index < len(arr) && arr[index] == target {
		fmt.Printf("Целевое значение %d найдено на позиции %d\n", target, index)
	} else {
		fmt.Printf("Целевое значение %d не найдено в массиве\n", target)
	}
}
