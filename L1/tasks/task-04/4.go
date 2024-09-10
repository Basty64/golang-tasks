package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
//которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

func main() {
	mainChan := make(chan int)

	var nWorkers int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&nWorkers)

	info := make(chan os.Signal, 1)
	signal.Notify(info, syscall.SIGINT, os.Interrupt)

	for i := 1; i <= nWorkers; i++ {
		go woker(mainChan, i)
	}

	for {
		select {
		case <-info:
			close(mainChan)
			close(info)
			return
		default:
			mainChan <- rand.Intn(100)
		}
	}
}

func woker(ch <-chan int, wid int) {
	for v := range ch {
		fmt.Printf("Получено число %d из канала от воркера %d\n", v, wid)
		time.Sleep(2 * time.Second)
	}

}

//---------------------------------------------------------------------------

func Check() {

	// Получение количества воркеров из командной строки
	var numWorkers int
	fmt.Print("Введите количество воркеров: ")

	numWorkers, err := fmt.Scan(&numWorkers)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Создание канала для передачи данных
	dataChan := make(chan string)

	// Создание группы воркеров
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Запуск воркеров
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			for data := range dataChan {
				fmt.Printf("Воркер %d: %s\n", workerID, data)
			}
		}(i)
	}

	// Запуск записи данных в канал
	go func() {
		for {
			dataChan <- fmt.Sprintf("Данные %s", time.Now().Format("15:04:05"))
			time.Sleep(time.Second)
		}
	}()

	// Обработка сигнала Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Завершение работы воркеров при получении сигнала Ctrl+C
	go func() {
		<-c
		fmt.Println("\nЗавершение работы...")
		close(dataChan)
		wg.Wait()
		os.Exit(0)
	}()

	// Запуск бесконечного цикла, чтобы программа не завершалась до нажатия Ctrl+C
	select {}
}

//----------------------------------------------------------------------------------------

func Ch() {
	var n int //количество воркеров

	fmt.Printf("Введите необходимое количество воркеров ") //выводим в stdout обращение к пользователю
	_, err := fmt.Scan(&n)                                 //получаем данные введенные от пользователя в консоль
	if err != nil {                                        //в случае ошибки, выводим ее ошибку в консоль
		fmt.Println(err)
		return
	}

	mainChannel := make(chan int, 100) //объявляем канал (главный поток)
	defer close(mainChannel)           //после выполнения функции main() закрываем канал

	ctx, cancel := context.WithCancel(context.Background()) //объявляем контекст с функцией отмены

	finishChannel := make(chan bool, n) // объявляем канал для контроля завершения всех воркеров
	defer close(finishChannel)

	for i := 0; i < n; i++ { //запускаем n количество воркеров
		go Worker(ctx, finishChannel, mainChannel, i)
	}
	go ExitSignal(cancel) //запускаем функцию перехвата сигнала о том что надо закрыть программу
loop: //даем лейбл циклу чтобы потом завершить его
	for {
		select {

		case mainChannel <- rand.Intn(10): //отправляем в главный канал случайные числа от 0 до 9
			time.Sleep(time.Second) // немного замедляем выполнение
		case <-ctx.Done():
			break loop // выходим из цикла при получении сигнала от контекста
		}
	}
	for i := 0; i < n; i++ {
		<-finishChannel //ожидаем выполнения всех воркеров
	}
	os.Exit(0) //завершаем программу с кодом 0

}
func Worker(ctx context.Context, finishChan chan<- bool, mainChan <-chan int, id int) {
	//используем лейбл чтобы потом выйти из цикла
loop2:
	for {
		select {
		case info := <-mainChan:
			fmt.Println(id, "-ий Воркер получил:", info) //если получаем данные из канала
			// - выводим в консоль
		case <-ctx.Done():
			break loop2 // если получаем сигнал выхода, выходим из цикла
		}
		fmt.Println("Воркер номер ", id, " закончил работу") //отправляем в канал данные что работа воркера окончена
		finishChan <- true
	}

}
func ExitSignal(cancel context.CancelFunc) {
	a := make(chan os.Signal, 1)                      //объявим канал для отправки сигнала о завершении программы
	signal.Notify(a, syscall.SIGINT, syscall.SIGTERM) //перехватываем сигнал
	fmt.Println("Получаем сигнал о завершении", <-a)

	cancel() // вызываем функцию завершения в контексте
}
