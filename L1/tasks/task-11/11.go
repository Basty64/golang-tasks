package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func main() {
	set1 := []int{1, 3, 5, 7, 9}
	set2 := []int{2, 3, 4, 5, 6}

	intersection := findIntersection(set1, set2)
	fmt.Println("Пересечение:", intersection) // Вывод: Пересечение: [3 5]
}

func findIntersection(set1 []int, set2 []int) []int {
	intersection := make([]int, 0)

	// Создаем карту для хранения элементов первого множества
	set1Map := make(map[int]bool)
	for _, element := range set1 {
		set1Map[element] = true
	}

	// Перебираем элементы второго множества
	for _, element := range set2 {
		// Если элемент присутствует в первом множестве, добавляем его в пересечение
		if _, ok := set1Map[element]; ok {
			intersection = append(intersection, element)
		}
	}

	return intersection
}
