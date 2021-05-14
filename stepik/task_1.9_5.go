package main

import "fmt"

func main() {
    var year int
    fmt.Scan(&year)
    isLeap := year % 400 == 0 || year % 4 == 0 && year % 100 != 0
    if isLeap {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    } 
}






