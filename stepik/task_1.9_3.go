package main

import "fmt"

func main() {
    var number int
    for fmt.Scan(&number); number >= 10; number /= 10 {}
    fmt.Println(number)
}
