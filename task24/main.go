/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

/*
Определяю структуру Point с инкапсулированными параметрами x и y.
Определяю конструктор NewPoint, который создает новый объект Point с заданными координатами.
Определяю метод DistanceTo для вычисления расстояние между двумя точками.
Метод принимает в качестве аргумента другой объект Point и использует формулу расстояния между двумя точками на плоскости,
чтобы вычислить расстояние между ними.

В функции main создаю две точки p и q с помощью NewPoint и вызываю метод DistanceTo на точке p,
передавая в качестве аргумента точку q.
*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
} // для определения новой точки

func (p *Point) DistanceTo(q *Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
} // рассчет дистанции до q

func main() {
	p := NewPoint(1, 2)
	q := NewPoint(4, 6)
	distance := p.DistanceTo(q)
	fmt.Println(distance)
}
