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
	tests := []struct{
		desc string
		name string
		recorder recorder
		want string
		expectErr bool
	}{
		{
			desc : "Error: recorder had some server error",
			name: "John",
			recorder: fakeRecorder{err: true},
			expectErr: true,
		},
		{
			desc: "Error: server returned wrong name",
			name: "John",
			recorder: fakeRecorder{data: Record{Name: "Bob", Age: 20}},
			expectErr: true,
		},
		{
			desc: "Success",
			name: "John",
			recorder: fakeRecorder{data: Record{Name: "John", Age: 20}},
			want: "Greetings John",
		},
	}

	for _, test := range tests {
		got, err := Greeter(test.name, test.recorder)
		switch {
		case err == nil && test.expectErr:
			t.Errorf("TestGreet (%s); got err == nil, want err != nil", test.desc)
			continue
		case err != nil && !test.expectErr:
			t.Errorf("TestGreet (%s); got err != %s, want err == nil", test.desc, err)
			continue
		case err != nil:
			continue
		}

		if got != test.want {
			t.Errorf("TestGreet (%s); got result %q, want %q", test.desc, got, test.want)
		}
	}
}