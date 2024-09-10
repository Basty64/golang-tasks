package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point
//с инкапсулированными параметрами x,y и конструктором.

// Структура для представления точки

type Point struct {
	x float64
	y float64
}

// Конструктор для создания точки

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Метод для вычисления расстояния между двумя точками

func (p *Point) Distance(q *Point) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

func main() {
	// Создание двух точек
	p := NewPoint(1, 2)
	q := NewPoint(4, 6)

	// Вычисление расстояния между точками
	distance := p.Distance(q)

	// Вывод результата
	fmt.Printf("Расстояние между точками: %f\n", distance)
}
