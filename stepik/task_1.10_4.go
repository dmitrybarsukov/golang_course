package main

import "fmt"

func main() {
    var (
        i int
        max int = 0
        maxCount int = 0
    )
    for fmt.Scan(&i); i > 0; fmt.Scan(&i) {
        if i > max {
            max = i
            maxCount = 1
        } else if i == max {
            maxCount++
        }
    }
    fmt.Println(maxCount)
}
