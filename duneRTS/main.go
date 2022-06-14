/* Tuesday Morning Review / Intro to Recv Methods


  func (receiver_name Type) method_name(parameter_list)(return_type){
  // Code
  }

*/

package main

import (
    // std library
    "fmt"

    // other stuff 3rd party imports
    "github.com/rzfeeser/duneRTS/models"
)

func main() {

    fmt.Println("Welcome to the Dune RTS game.")


    harvester1 := models.SpiceHarvester{5, false, false, 2}
    fmt.Println(harvester1)

    harvester1.AddCrew()
    fmt.Println(harvester1)

    fmt.Println(harvester1.IsWorried())

}
