package main

import "fmt"

func main() {
    var i int
    fmt.Scan(&i)
    fmt.Println(i / 10 % 10)
}
