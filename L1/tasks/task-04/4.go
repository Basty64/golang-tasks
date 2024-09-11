package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
//которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

func mai() {
	mainChan := make(chan int)

	var nWorkers int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&nWorkers)

	//выход из программы через ctrl+C
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
