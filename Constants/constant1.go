package constants

import (
	"fmt"
	// "strings"
)


type specialStr string

func printSpecial(str specialStr) {
	fmt.Println(string(str))
}

// func main() {
// 	const unTyedstr = "Test 1 check"

// 	var testStr = "Test 2 check"
	
// 	printSpecial(unTyedstr) // Constants are untyped and don't need casting
// 	printSpecial(specialStr(testStr)) // testStr is type
// }