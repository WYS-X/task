package main

import "fmt"

type Shape interface {
	Area() float32
	Perimeter() float32
}
type Rectangle struct {
	Length float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}
func (r Rectangle) Perimeter() float32 {
	return (r.Length + r.Width) * 2
}

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return float32(3.14) * c.Radius * c.Radius
}
func (c Circle) Perimeter() float32 {
	return 2 * 3.14 * c.Radius
}

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeId int
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工姓名：%s, 年龄：%d", e.Name, e.Age)
}
func main() {
	r := Rectangle{Length: 2, Width: 3}
	fmt.Println("rectangle length=2, width=3, area=", r.Area(), ", perimeter=", r.Perimeter())

	c := Circle{Radius: 10}
	fmt.Println("circle radius=10, area=", c.Area(), ", perimeter=", c.Perimeter())

	e := Employee{Person: Person{Name: "xiao wang", Age: 40}}
	e.PrintInfo()
}
