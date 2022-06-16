package main

import(
    "fmt"
)


// coding contract that promises certain behavior(s)
type Player interface {
    Defend() string
    //Attack()
}

// type Wizard
type Wizard struct{}

// Wizard recv-function - Defend()
func (w Wizard) Defend() string {
    return "Expelliarmus"
}

// Wizard recv-function - CastForget()
func (w Wizard) CastForget() string {
    return "Obliviate"
}


// type Barbarian
type Barbarian struct{}

func (b Barbarian) Defend() string{
    return "Dodge"
}


// main function
func main() {
    
    gandalf := Wizard{}
    conan := Barbarian{}

    // make both instances Defend()
    //fmt.Println(gandalf.Defend())
    //fmt.Println(conan.Defend())

    // the interface promises us there is some commonality between everything in this slice
    // in this case, they are all type Player
    playerCharacters := []Player{gandalf, conan}

    for _, x := range playerCharacters {
        fmt.Println(x.Defend())
    }

}


