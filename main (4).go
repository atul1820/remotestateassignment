package main

import "fmt"

// Function with parameters and return value
func add(x, y int) int {
    return x + y
}

func main() {
    result := add(3, 5)
    fmt.Println("Sum:", result)
}