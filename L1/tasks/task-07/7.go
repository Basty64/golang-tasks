package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Реализовать конкурентную запись данных в map.

func main() {
	// Создаем map с защитой от одновременного доступа
	var dataMap sync.Map
	var wg sync.WaitGroup
	wg.Add(10)

	// Запускаем 10 горутин, которые записывают данные в map
	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				// Записываем данные в map с ключом "key" + id + " " + j
				key := fmt.Sprintf("key%d %d", id, j)
				dataMap.Store(key, fmt.Sprintf("Значение %d %d", id, j))
				time.Sleep(time.Millisecond * 500)
			}
		}(i)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим содержимое map
	dataMap.Range(func(key, value interface{}) bool {
		fmt.Printf("Ключ: %v, Значение: %v\n", key, value)
		return true
	})
}

//-----------------------------------------------

// Структура безопасной мапы
type SafeMap struct {
	mu   sync.RWMutex // RWMutex для безопасного чтения и записи в мапу
	data map[string]int
}

func NewMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock() // Блокируем запись в мапу для других горутин, горутина ждет завершения метода Set
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *SafeMap) Get(key string) int {
	s.mu.RLock() // Блокирует только чтение из мапы, при этом запись не заблокирована
	defer s.mu.RUnlock()
	return s.data[key]
}

func mai() {
	var wg sync.WaitGroup // Создадим объект WaitGroup для ожидания выполнения горутин

	safemap := NewMap()
	stringelems := []string{"test", "hello", "world", "mutex", "gopher"}
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		for i, v := range stringelems {
			safemap.Set(v, i)
		}
		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		for _, v := range stringelems {
			fmt.Println(safemap.Get(v))
		}
		wg.Done()
	}(&wg)

	wg.Wait() // Программа ждет выполнение всех горутин
}

//---------------------------------

type Data struct { //объявляем вспомогательную структуру
	id   int
	name int
}

func ma() {
	//Channel()
	Mutex()
}

func Channel() {
	n := 5                         // количество писателей
	data := make(map[int]int)      //создаем мапу для записи
	chanData := make(chan Data, n) //создаем канал для передачи данных

	for i := 0; i < n; i++ { //запускаем n горутин

		go func(ch chan Data) {
			for {
				ch <- Data{id: i, name: i} //отправляем данные в канал
			}
		}(chanData)

	}
	for { //бесконечный цикл чтения из канала
		tmp := <-chanData //получаем данных
		data[tmp.id] = tmp.name
	}
}

func Mutex() {
	n := 5 //задаем количество писателей

	data := make(map[int]int) //задаем мапу

	var mtx sync.Mutex //создаем Mutex

	for i := 0; i < n; i++ {
		go func(data map[int]int, mtx *sync.Mutex) {
			for {
				id := rand.Int()
				name := rand.Int()
				mtx.Lock()      //блокируем Mutex
				data[id] = name //записываем
				mtx.Unlock()    //открываем Mutex
			}
		}(data, &mtx)

	}

}

func SyncMap() {
	n := 5                   //задаем количество писателей
	var data sync.Map        //создаем мапу с
	for i := 0; i < n; i++ { //начинаем цикл
		go func(data *sync.Map) {
			for {
				data.Store(rand.Int(), rand.Int()) //записываем значения в мапу
			}
		}(&data)
	}
}
