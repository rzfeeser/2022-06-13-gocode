/* intro to concurrent API looksup */

package main

import (
    "fmt"
    "sync"
)

// make a var @ package level to track threads (go routines)
var wg sync.WaitGroup

// a function that mimics an HTTP lookup
func apiLookup(api string) {

    // ensure main doesn't close before our worker finishes!
    defer wg.Done()

    // this should actually be an HTTP request
    fmt.Println("Looking up API", api)
}

func main() {
    restfulAPIs := []string{"https://example.com", "https://example.net", "https://alta3.com", "http://go.dev"}

    // loop across the lookups we need to perform
    for _, api := range restfulAPIs {
        wg.Add(1)          // add 1 to the wg per worker (thread) created
        go apiLookup(api)  // create API workers
    }

    // this causes us to wait until all threads have finished (running wg.Done() indicates finished)
    wg.Wait()

}
