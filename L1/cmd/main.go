package main

import (
	"fmt"
	task_04 "golang-tasks/L1/tasks/task-04"
	"os"
)

func main() {
	//task_01.Check()
	//task_02.Check()
	//task_03.Check()
	task_04.Che()
	//task_06.Check()
	//task_07.Check()
	//task_08.Check()
	//task_09.Check()
	//task_10.Check()
	//task_11.Check()
	//task_12.Check()
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
		folderName := fmt.Sprintf("task-%d", i)
		folderPath := path + "/" + folderName

		// Создаем папку
		err := os.Mkdir(folderPath, 0755)
		if err != nil {
			fmt.Println("Ошибка при создании папки:", err)
			return
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
    // Это файл %d
}
`, i))
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}

		fmt.Println("Создана папка:", folderPath)
		fmt.Println("Создан файл:", filePath)
	}
}
