package main

import "fmt"

func main() {
    defer fmt.Println("This will be executed last.")
    fmt.Println("This will be executed first.")
}