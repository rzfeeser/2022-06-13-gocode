/* Alta3 Research | RZFeeser

One of the most simple interfaces to impliment is the error interface. Provided by Go,
this only requires that you impliment a receiver-function that accepts no parameters and returns a string.

    type error interface {
        Error() string
    }

*/

package main

import(
    "fmt"
)

type hardwareProblem struct {
    system string
    location string
}

func (hp hardwareProblem) Error() string {
    return fmt.Sprintf("hardware error! %s, is happening at : %s", hp.system, hp.location)
}

type networkProblem struct {
	message string
	code    int
}

func (np networkProblem) Error() string {
	return fmt.Sprintf("network error! message: %s, code: %v", np.message, np.code)
}

func handleErr(err error) {
	fmt.Println(err.Error())
}

func main(){

np := networkProblem{
	message: "we received a problem",
	code:    404,
}

hp := hardwareProblem{
    system: "battery backup not initalized",
    location: "datacenter West",
}

handleErr(np)
handleErr(hp)

}
