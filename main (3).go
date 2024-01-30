package main

import "fmt"

// Define an interface
type Shape interface {
    Area() float64
}

// Implement the interface for a rectangle
type Rectangle struct {
    Width  float64
    Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rectangle := Rectangle{Width: 5, Height: 10}
    fmt.Println("Area of rectangle:", rectangle.Area())
}