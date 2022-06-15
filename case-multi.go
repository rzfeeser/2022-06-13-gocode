package main

import "fmt"

// Zach prints and returns n.
func Zach(n int) int {
    fmt.Println(n)
    return n
}

func main() {
    switch Zach(9) {
    case Zach(1), Zach(2), Zach(9), Zach(10):
        fmt.Println("First case")
        fallthrough
    case Zach(4):
        fmt.Println("Second case")
    }
}
