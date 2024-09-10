package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

// Интерфейс для источника данных
type DataSource interface {
	GetData() string
}

// Конкретный источник данных (в данном случае, файл)
type FileDataSource struct {
	filename string
}

func (f *FileDataSource) GetData() string {
	// Здесь должна быть логика чтения данных из файла
	return fmt.Sprintf("Данные из файла: %s", f.filename)
}

// Адаптер, который адаптирует FileDataSource к интерфейсу DataSource
type FileDataSourceAdapter struct {
	fileDataSource *FileDataSource
}

func (a *FileDataSourceAdapter) GetData() string {
	return a.fileDataSource.GetData()
}

func main() {
	// Создаем источник данных (файл)
	fileDataSource := &FileDataSource{filename: "data.txt"}

	// Создаем адаптер
	adapter := &FileDataSourceAdapter{fileDataSource: fileDataSource}

	// Используем адаптер как DataSource
	data := adapter.GetData()
	fmt.Println(data)
}

//Паттерн Адаптер - это структурный паттерн проектирования, который позволяет использовать классы
//с несовместимыми интерфейсами, делая их совместимыми.
//Он решает проблему, когда у вас есть класс, который не соответствует тому интерфейсу,
//который требуется другому классу или части системы.
