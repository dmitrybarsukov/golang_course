package main

import "fmt"

func getNumberLength(num int) int {
    for i := 1;; i++ {
        num /= 10
        if num == 0 {
            return i
        }
    }
}

func getDigit(num int, digit int) int {
    for ; digit > 0; digit-- {
        num /= 10
    }
    return num % 10
}

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    for i := getNumberLength(a) - 1; i >= 0; i-- {
        da := getDigit(a, i)
        for j := getNumberLength(b) - 1; j >= 0; j-- {
            db := getDigit(b, j)
            if da == db {
                fmt.Print(da, " ")
            }
        }
    }
}
