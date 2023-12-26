package struct


import "fmt"

// Custom struct type
type Record struct {
	Name string
	Age int
}


// Type method of Record struct
func (r Record) String() string {
	return fmt.Sprintf("%s,%d", r.Name , r.Age)
}


// Changing a field's value
func changeName(r *Record) {
	r.Name = "Peter"
	fmt.Println("inside changeName: ", r.Name)
}


// Changing a field's value in a method
func (r *Record) IncrAge() {
	r.Age++
}


// func main () {
// 	david := Record{Name: "David", Age: 28}
// 	sarah := Record{Name: "Sarah", Age: 28}

// 	john := Record{Name: "John", Age: 100}
// 	fmt.Println(john.String()) // The String method can only be used by the Record object

// 	fmt.Printf("%+v\n", david)
// 	fmt.Printf("%+v\n", sarah)

// 	//Changing the name value of a record object
// 	rec := Record{Name: "Emmanuel"}
// 	changeName(&rec) // Passing the address of the rec variable to the changeName function
// 	fmt.Println("main: " ,rec.Name)

// 	// Testing the IncrAge method
// 	joseph := &Record{Age: 50}
// 	joseph.IncrAge()
// 	fmt.Println(joseph.Age)

// }