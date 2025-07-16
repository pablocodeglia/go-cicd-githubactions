package greet

import "testing"

func TestGreet(t *testing.T) {
	want := "Hey"
	got := SayHi()

	if got != want {
		t.Errorf("\n-Got: '%s' \n-Expected: '%s'", got, want)
	}
}
