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

//-------------------------------------------------------

//Second variant

var TestArray = []int{1, 3, 4, 6, 8, 10, 55, 56, 59, 70, 79, 81, 91, 10001}

func NewBinarySearch(mas []int, search int) (bool, int) {

	var index int

	for mas[index] != search {
		a := len(mas) / 2
		if mas[a] > search {
			mas = mas[:a]
		} else if mas[a] < search {
			mas = mas[a:]
		} else {
			fmt.Print("true", index)
			return true, index
		}
	}
	return false, 0
}
