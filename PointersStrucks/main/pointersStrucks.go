package main

import "fmt"
import "math"

func main()  {
	/*fmt.Println("Hello")

	x:=0
	//pass pointer &
	changeValNow(&x)
	fmt.Println("x= ", x)

	rect1 := Rectangle{10,10}
	// if you dont know order of struck rect1 := Rectangle{leftX: 0, TopY: 50, height: 10, width: 10}
	fmt.Println("hright", rect1.height , " width", rect1.width)

	fmt.Println("Area of rect1 is ", rect1.area())*/
	rect := Rectangle{20,50}
	circ := Circle{4}
	fmt.Println("Rectangle Area= ", getArea(rect))
	fmt.Println("Circle Area = ", getArea(circ))


}
/// strucks interfaces
type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width float64
}


type Circle struct {
	radius float64
}

func (c Circle) area() float64  {
	return math.Pi * math.Pow(c.radius,2)
}


//STRUCKS

func (rect Rectangle) area() float64{
	return rect.width * rect.height
}
func getArea(shape Shape) float64 {
	return shape.area()
}

// pointers * references a pointer
func changeValNow(x *int){
	*x=3

}