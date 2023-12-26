package interfaces

import (
	"fmt"
	"math"
	"strings"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// For rectangles : The interface is accessed by calling a method in it
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * r.width + 2 * r.height;
}

// For circles:

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}


// Example 2

type Stringer interface {
	String() string
}

type Person struct {
	First, Last string
}

func (p Person) String() string {
	return fmt.Sprintf("%s,%s", p.Last, p.Last)
}

type StrList []string

func (s StrList) String() string {
	return strings.Join(s, ",")
}

// PrintStringer prints the value of Stringer to Stdout
func PrintStringer(s Stringer) {
	fmt.Println(s)
	fmt.Println(s.String())
}
// func main() {
// 	r := rect{width: 3, height: 4}
// 	c := circle{radius: 5}

// 	measure(r)
// 	measure(c)
// }