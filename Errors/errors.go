package errors

import (
	"errors"
	"fmt"
)


// errors.New constructs a basic error value with the given error message
func f1(arg int) (int,  error) {
	if arg  == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil // A nil value in the place of error indicates there's no error
}


// Creating custom error types as errors by implementing the Error method on them
type argError struct {
	arg int
	prob string
}

// Error method to provide a string representation of the error when printed
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"} //using argError to create a new error type
	}
	return arg + 3, nil
}

// func main() {

// 	// error returning function f1
// 	for _, i := range []int{7, 42} {
// 		if result, error := f1(i); error != nil {
// 			fmt.Println("f1 failed:", error)
// 		} else {
// 			fmt.Println("f2 worked:", result)
// 		}
// 	}


// 	for _, i := range []int{7, 42} {
// 		if result, error := f2(i); error != nil {
// 			fmt.Println("f2 failed:", error)
// 		} else {
// 			fmt.Println("f2 worked:", result)
// 		}
// 	}

// 	// To programmatically use the data in the custom type
// 	// The error is taken as an instance of the custom error type via assertion
	
// 	_ , e := f2(42)
// 	if ae, ok := e.(*argError); ok {
// 		fmt.Println(ae.arg)
// 		fmt.Println(ae.prob)
// 	}
// }

