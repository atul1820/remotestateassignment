package main

import "fmt"

type Car struct {
    Brand  string
    Model  string
    Year   int
    Engine string
}

func main() {
    
    myCar := Car{
        Brand:  "Toyota",
        Model:  "Camry",
        Year:   2022,
        Engine: "V6",
    }

    // Access fields of the struct
    fmt.Printf("My car is a %d %s %s with a %s engine.\n", myCar.Year, myCar.Brand, myCar.Model, myCar.Engine)
}