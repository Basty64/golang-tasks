package main

import (
	"fmt"
	"os"
)

func main() {

	createFolders()

}

func createFolders() {

	// Получаем путь к родительской папке и количество папок
	var path string
	var n int
	fmt.Print("Введите путь к родительской папке: ")
	fmt.Scanln(&path)
	fmt.Print("Введите количество папок: ")
	fmt.Scanln(&n)

	// Создаем папки и файлы
	for i := 7; i <= n; i++ {
		var folderName string
		if i < 10 {
			folderName = fmt.Sprintf("task-0%d", i)
		} else {
			folderName = fmt.Sprintf("task-%d", i)
		}
		folderPath := path + "/" + folderName
		// Создаем папку
		err := os.Mkdir(folderPath, 0755)

		if err != nil {
			fmt.Println("Ошибка при создании папки:", err)
		}

		// Создаем файл .go
		fileName := fmt.Sprintf("%d.go", i)
		filePath := folderPath + "/" + fileName

		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Ошибка при создании файла:", err)
			return
		}
		defer file.Close()

		// Записываем код в файл
		_, err = file.WriteString(fmt.Sprintf(`package main

func main() {

}
`))
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}

		fmt.Println("Создана папка:", folderPath)
		fmt.Println("Создан файл:", filePath)
	}
}
