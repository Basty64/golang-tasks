package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func main() {

	//1st variant

	var v interface{} = 10
	detectType(v)

	v = "Hello, world!"
	detectType(v)

	v = true
	detectType(v)

	ch := make(chan int)
	v = ch
	detectType(v)

	//2nd variant

	arrOfTypes := []interface{}{'a', "hello", 5, 0.23, true, struct{}{}}
	for _, typ := range arrOfTypes {
		typ := RecognizeType(typ)
		fmt.Printf("Value: %v, type: %s\n", typ, typ.Kind().String())
	}
}

func detectType(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Printf("Тип переменной: %s\n", t.String())
}

func RecognizeType(unknown interface{}) reflect.Value {
	return reflect.ValueOf(unknown)
}
