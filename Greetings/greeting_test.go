package greetings

import (
	"fmt"
	"testing"
)

func Greet(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name was empty")
	}
	return "Hello " + name, nil
}

func TestGreet1(t *testing.T) {
	name := "Bob"
	want := "Hello Sarah"

	got, err := Greet(name)
	if got != want || err != nil {
		t.Fatalf("TestGreet(%s): got %q/%v, want %q/nil", name, got, err, want)
	}
}

func TestGreet (t *testing.T) {
	tests := []struct{
		desc string
		name string
		want string
		expectErr bool
	}{
		{
			desc: "Error: name is an empty string",
			expectErr: true,
			// name and want are "", the zero value of the string
		},
		{
			desc: "Success",
			name: "John",
			want: "Hello John",
			// expectErr  is set to zero value, false
		},
	}

	for _, test := range tests {
		got, err := Greet(test.name)
		switch {
			case err == nil && test.expectErr
		}
	}
}
