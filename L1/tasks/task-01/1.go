package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action
//от родительской структуры Human (аналог наследования).

//Мы встраиваем структуру Human в структуру Action.
//Это позволяет Action использовать все поля и методы Human.
//Метод DoAction в Action использует метод Greet из Human,
//показывая, как встраивание позволяет реализовать функциональность, похожую на наследование.

// Создаём структуру human
type Human struct {
	Name    string
	Surname string
	Age     int
}

// Создаём структуру action
type Action struct {
	Human
	Case string
}

func (h Human) DescribeHuman() {
	fmt.Printf("Hello, this is a %s %s, %d years old\n", h.Name, h.Surname, h.Age)
}

// Вызываем метод DescribeHuman
func (a Action) DoAction() {

	//Встраивание
	a.DescribeHuman()
	fmt.Printf("Now i can do: %s\n", a.Case)
}

func main() {

	action := Action{
		Human: Human{
			Name:    "Joe",
			Surname: "McDir",
			Age:     45,
		},
		Case: "walking",
	}
	action.DoAction()
}
