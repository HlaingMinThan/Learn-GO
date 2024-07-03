package main

import "fmt"

type rectangle struct {
	height float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	getArea() float64
}

func (r rectangle) getArea() float64 {
	return r.height * r.width
}
func (c circle) getArea() float64 {
	return c.radius * c.radius * 3.14
}

func printShapeInfo(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	printShapeInfo(rectangle{height: 10, width: 10})
	printShapeInfo(circle{radius: 10})
}
