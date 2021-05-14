package main

import "fmt"

func main() {
    for {
        var i int
        fmt.Scan(&i)
        if i > 100 {
            break
        } else if i >= 10 {
            fmt.Println(i)
        }
    }
}
