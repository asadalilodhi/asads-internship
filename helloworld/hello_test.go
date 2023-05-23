package main

import (
	"testing"
)

func OK() string {
	return "They  match!"
}

func NOTOK() string {
	return "They dont match!"
}

func TestHello(t *testing.T) {
	got := Hello("Asad")
	want := Hi("Asad")

	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// 	fmt.Println(NOTOK())
	// }

	// if got == want {
	// 	// t.Errorf("got %q want %q", got, want)
	// 	fmt.Println(OK())

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
