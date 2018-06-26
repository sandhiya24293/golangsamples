package main

import (
	"fmt"
	//"math"
)


type Shape interface {
	GetArea()
	GetPerimeter()
}


type Rectangle struct {
	Len int
	Brh int
}
func (rect *Rectangle) GetArea()  {
	fmt.Println("RectArea:==",rect.Len * rect.Brh)
}
func (rect *Rectangle) GetPerimeter()  {
	fmt.Println("Rect Pri:==", 2*(rect.Len + rect.Brh))
}

type Circle struct {
	radius int
}
func (rect Circle ) GetArea()  {
	fmt.Println("CrcArea:==", 2* 3.14 * (float64(rect.radius * rect.radius)))
}
func (rect Circle ) GetPerimeter() {
	fmt.Println("CrcPeri:==", 2* 3.14 * (float64(rect.radius)))
}


func main() {
	rect := Rectangle{4,5}
	circl := Circle{3}
	var MyShape Shape
	MyShape = &rect //As pointer reciver
	MyShape.GetArea()

	MyShape = circl// As value reciver
	MyShape.GetArea()

}


/*type geometry interface {
	area() float64
	perim() float64
}
//For our example we’ll implement this interface on rect and circle types.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

//To implement an interface in Go, we just need to implement all the methods in the interface. Here we implement geometry on rects.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

//The implementation for circles.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//If a variable has an interface type, then we can call methods that are in the named interface. Here’s a generic measure function taking advantage of this to work on any geometry.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	//The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
	var G geometry
	G = r
	G.area()
	measure(r)
	measure(c)
}*/
