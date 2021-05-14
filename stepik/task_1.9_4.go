package main

import "fmt"

func main() {
    var number int
    fmt.Scan(&number)
    
    if number < 100000 || number >= 1000000 {
        fmt.Println("Not 6-digit number")
        return
    }
    
    s1 := number / 100000 % 10 + number / 10000 % 10 + number / 1000 % 10
    s2 := number / 100 % 10 + number / 10 % 10 + number % 10
    if s1 == s2 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}




