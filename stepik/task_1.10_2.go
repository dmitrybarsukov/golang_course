package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Println((b * b - a * a + a + b) / 2)
}
