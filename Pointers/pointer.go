package pointers

import (
	"fmt"
)

func changeWord(word *string){

	// "*" on a variable or function argument is dereferencing
	*word += "change"
}

func someCode(){
	var x int = 30;
 fmt.Println(x)

	//Pointer stores address
var intPr *int = &x
fmt.Println(intPr)

var y string = "Worked";
changeWord(&y)
fmt.Println(y)
}
