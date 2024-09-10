package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.

//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func main() {
	testCases := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, testCase := range testCases {
		fmt.Printf("Строка: %s - %t\n", testCase, isUniqueChars(testCase))
	}
}

func isUniqueChars(s string) bool {
	// Преобразуем строку в нижний регистр
	s = strings.ToLower(s)

	// Создаем мапу для отслеживания символов
	charMap := make(map[rune]bool)

	// Проходимся по каждому символу в строке
	for _, char := range s {
		// Если символ уже есть в мапе, значит он не уникальный
		if _, ok := charMap[char]; ok {
			return false
		}

		// Добавляем символ в мапу
		charMap[char] = true
	}

	// Если все символы уникальны, возвращаем true
	return true
}
