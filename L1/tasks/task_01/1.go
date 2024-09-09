package task_01

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     int
}

type Action struct {
	Human
	Case string
}

func (h Human) DescribeHuman() {
	fmt.Printf("Hello, this is a %s %s, %d years old\n", h.Name, h.Surname, h.Age)
}

func (a Action) DoAction() {
	a.DescribeHuman()
	fmt.Printf("Now i can do: %s\n", a.Case)
}

//В этом примере мы встраиваем структуру Human в структуру Action.
//Это позволяет Action использовать все поля и методы Human.
//Метод DoAction в Action использует метод Greet из Human,
//показывая, как встраивание позволяет реализовать функциональность, похожую на наследование.

func Example1() {

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
