package main

import "fmt"

func main() {
    var (
        n int
        sum = 0
    )
    fmt.Scan(&n)
    for ; n > 0; n-- {
        var num int
        fmt.Scan(&num)
        if num >= 10 && num < 100 && num % 8 == 0 {
            sum += num
        }
    }
    fmt.Println(sum)
}