/* Alta3 Research | RZFeeser@alta3.com
   Understanding Go Receiver Functions (i.e. methods)

   func(receiver_name Type) method_name(parameter_list)(return_type){
   // Code
   }

*/

package main

import "fmt"

type Player struct {
    Lives int
    Stage int
    Inventory []string
}

// method to increase life by 1
func (p *Player) Greenmushroom() {
    p.Lives++
}

func main() {
    mario := Player{3, 1, []string{"mushroom"}}

    fmt.Println(mario.Lives)

    mario.Greenmushroom()

    fmt.Println(mario.Lives)

}
