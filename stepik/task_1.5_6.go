package main

import "fmt"

func main() {
    var i int
    fmt.Scan(&i)
    fmt.Println("It is", i / 30, "hours", i % 30 * 2, "minutes.")
}
