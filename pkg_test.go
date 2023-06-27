package main

import "testing"

func TestAddition(t *testing.T) {
	got := Addition(11, 28)
	want := 39

	if got != want {
		t.Errorf("Got %v and wanted %v", got, want)
	}
}

func TestSubtraction(t *testing.T) {
	got := Subtraction(22, 30)
	want := -8

	if got != want {
		t.Errorf("Got %v and wanted %v", got, want)
	}
}
