package main

import "fmt"

func main() {
    var x, p, y int
    var s = 0
    for fmt.Scan(&x, &p, &y); x <= y; x += x * p / 100 {
        s++
    }
    fmt.Println(s)
}
