package deferpanicrecover


import "fmt"

func printStuff() (value string) {
	defer fmt.Println("exiting")
	defer func() {
		value = "we returned this"
	} ()
	fmt.Println("I am printing stuff")
	return ""
}

// func main() {
// 	v := printStuff()
// 	fmt.Println(v)
// }