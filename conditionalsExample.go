package main

import (
    "fmt"
    "math/rand"
    "time"
)

func init(){
    // rand.Seed(1)  // randomness has this "seed" value by default
    rand.Seed(time.Now().Unix()) 
}

func main(){

    // create a slice of strings
    instructorSlice := []string{"Jane", "Zach", "Talia", "Chad"}

    // take a random selection from our instructorSlice
    // x := instructorSlice[rand.Intn(len(instructorSlice))]

    //x := "Jane"
    if x := instructorSlice[rand.Intn(len(instructorSlice))]; x == "Zach" {
        fmt.Println("Zach is the assigned instructor")
    } else if x == "Jane" {
        fmt.Println("Jane will teach this class")
    } else {
        fmt.Println("A new hire will teach:", x)
    }

    // this should error, x is local to the if-else-if-else block
    //fmt.Println(x)

}
