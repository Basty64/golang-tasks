package main

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func main() {

	// Создаем канал для сигналов
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Запускаем бесконечный цикл
	for {
		// Проверяем, не пришёл ли сигнал прерывания
		select {
		case <-c:
			fmt.Println("Получен сигнал прерывания. Завершаем работу...")
			return
		default:
			time.Sleep(3 * time.Second)
			arrOfTypes := []interface{}{'a', "hello", 5, 0.23, true, struct{}{}}
			for _, typ := range arrOfTypes {
				typ := detectType(typ)
				fmt.Printf("Value: %v, type: %s\n", typ, typ.Kind().String())
			}
		}
	}
}

func detectType(unknown interface{}) reflect.Value {
	return reflect.ValueOf(unknown)
}
