package main

import "fmt"

func main() {
    var i int
    fmt.Scan(&i)
    switch {
    case i > 0: fmt.Println("Число положительное")
    case i < 0: fmt.Println("Число отрицательное")
    default: fmt.Println("Ноль")
    }
}