package greetings

import (
	"errors"
	"testing"
	"fmt"
)

// type  Fetch struct{

// }

// func (f *Fetch) Record(name string) (Record, error) {
// 	//Some code that talks to a server
// }
type Record struct {
	Name string
	Age int
}

type recorder interface {
	Record(name string) (Record, error)
}


// Now this function can paass a *client.Fetch instance or anything that implements the recorder interface
func Greeter(name string, fetch recorder) (string, error) {
	rec, err := fetch.Record(name)
	if err != nil {
		return "", err
	}

	if rec.Name != name {
		return "", fmt.Errorf("server returned record for %s not %s", rec.Name, name)
	}
	if rec.Age < 18 {
		return "Greetings young one", nil
	}
	return fmt.Sprintf("Greetings %s", name), nil
}

type fakeRecorder struct{
	data Record
	err bool
}

func (f fakeRecorder) Record(name string) (Record, error) {
	if f.err {
		return Record{}, errors.New("error")
	}
	return f.data, nil
}

func TestGreeter (t *testing.T) {
	
}