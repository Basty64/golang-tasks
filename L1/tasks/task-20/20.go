package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».

func reverseWords(s string) string {
	words := strings.Fields(s) // Разбиваем строку на слова
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i] // Меняем слова местами
	}
	return strings.Join(words, " ") // Соединяем слова обратно в строку
}

func main() {
	input := "snow dog sun"
	reversed := reverseWords(input)
	fmt.Println("Перевернутая строка:", reversed)
}
