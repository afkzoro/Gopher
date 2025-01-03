package main

import "fmt"


func main() {
	fmt.Println("Using send-only channels")

	c := make(chan int, 4) 
	fmt.Printf("lencgth of c: %v\n", len(c))
	producer(c)




	for v := range c {
		fmt.Printf("value of v: %v\n", v)		
	}

	// Reading from a channel
	// var v = <- c
	// fmt.Printf("value of v: %v\n", v)


}


// Send only channel
func producer (c chan <- int) {
	// Writing to a channel
	c <- 10
	c <- 5
	close(c)
}
