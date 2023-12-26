package constructor

import "fmt"

// Constructors in Golang have a New[Type] () pattern when declaring a public constructor
// Use New() if there are types in the package or they won't be even in the future

type Record struct {
	Name string
	Age int
}

func NewRecord (name string, age int) (*Record, error) {
	if name ==  ""{
		return nil, fmt.Errorf("name cannot be an empty string")
	}
	if age <= 0 {
		return nil, fmt.Errorf("age cannot be <= 0")
	}

	return &Record{Name: name, Age: age}, nil
}

// func main() {
// 	rec, err := NewRecord("John", 100)
// 	if err != nil {
// 		return err
// 	}

// 	return rec
// } 