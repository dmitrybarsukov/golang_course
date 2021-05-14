package main

import "fmt"

func main() {
    var f float64
    fmt.Scan(&f)
    if f <= 0 {
        fmt.Printf("число %2.2f не подходит", f)
    } else if f > 10000 {
        fmt.Printf("%e", f)
    } else {
        f *= f
        fmt.Printf("%.4f", f)
    }
}