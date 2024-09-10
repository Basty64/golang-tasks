package main

import "fmt"

//Удалить i-ый элемент из слайса.

func main() {
	slice := []int{1, 2, 3, 4, 5}
	i := 2 // Индекс элемента, который нужно удалить

	// Удаление i-го элемента из слайса
	slice = append(slice[:i], slice[i+1:]...)

	fmt.Println("Слайс после удаления:", slice)
}
