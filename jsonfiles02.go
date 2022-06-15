/* RZFeeser | Alta3 Research
Writing out JSON              */


package main

import (
    "fmt"
    "encoding/json"
    "os"
)

func main() {
    type Person struct {
        Fn string `json:"first name"`
        Ln string `json:"last name"`
    }
    type ColorGroup struct {
        ID     int
        Name   string
        Colors []string
        P      Person `json:"person"`
    }

    per := Person{Fn: "RZ",
        Ln: "Feeser",
    }

    group := ColorGroup{
        ID:     24601,
        Name:   "Greens",
        Colors: []string{"Moss", "Shamrock", "Lime", "Hunter"},
        P:      per,
    }
    
    // serialize values from struct to json format
    b, err := json.Marshal(group)
    
    if err != nil {
        panic(err)
    }


    fmt.Println(len("\n"))
    fmt.Println(len(`\n`))


    // print to standard out
    os.Stdout.Write(b)
}

