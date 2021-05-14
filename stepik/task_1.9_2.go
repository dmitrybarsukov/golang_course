package main

import "fmt"

func main() {
    var number int
    fmt.Scan(&number)
    if number < 100 || number > 999 {
        fmt.Println("Not 3-digit number")
        return
    }
    
    d1, d2, d3 := number / 100 % 10, number / 10 % 10, number % 10
    allDifferent := d1 != d2 && d1 != d3 && d2 != d3
    if allDifferent {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

