package test1

import "testing"

func TestHello(t *testing.T) {
	want := "Hello"
	if got := Hello(want); got != want {
		t.Errorf("Hello()=%s, want=%s", got, want)
	}
}
