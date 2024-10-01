package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "evev" {
		t.Errorf("exepected: even, actual: %s", result)
	}
}