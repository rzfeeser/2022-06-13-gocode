/* Alta3 Research | RZFeeser
   goroutine - with waitgroups   */

package main

import (
        "fmt"
        "time"
        "sync"
)

func astroTalk() {
    defer wg.Done()

    fmt.Println("Wow this should be fun!")
    time.Sleep(time.Second)
    fmt.Println("Going to the moon!")
    time.Sleep(time.Second)
    fmt.Println("Almost time.....")
}

func countDown(s int) {
        defer wg.Done()  // this is run when the thread is ended

        for i := s; i > 0; i-- {
                fmt.Println(i)
                time.Sleep(time.Second)   // wait one second
        }
}


// WaitGroup is used to wait for the program to finish goroutines
var wg sync.WaitGroup

func main() {

    fmt.Println("NASA launch sequence starting:")

    // the waitGroup has "1" in the queue
     wg.Add(1)

    // Begin the countdown
    go countDown(10)


    wg.Add(1)

    go astroTalk()

    fmt.Println("Launch sequence starting:")

    time.Sleep(time.Second)
    fmt.Println("Hey wait a second...")
    
    time.Sleep(time.Second)
    fmt.Println("I forgot my wallet!")

    // wait for the waitGroup queue to reach 0
    wg.Wait()

    fmt.Println("TAKE OFF!")
}

