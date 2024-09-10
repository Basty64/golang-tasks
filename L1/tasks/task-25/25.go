package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C // Ожидание завершения таймера
}

func main() {
	fmt.Println("Начало выполнения...")
	var seconds int
	fmt.Println("Введите кол-во секунд для сна программы:")
	fmt.Scanln(&seconds)
	sleep(time.Duration(seconds) * time.Second) // Приостановка на 2 секунды
	fmt.Println("Продолжение выполнения...")
}
