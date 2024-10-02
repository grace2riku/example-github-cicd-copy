// コード4.2 テストコード
package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "even" {
		t.Errorf("exepected: even, actual: %s", result)
	}
}

func TestOdd(t *testing.T) {
	result := EvenOrOdd(5)
	if result != "odd" {
		t.Errorf("exepected: odd, actual: %s", result)
	}
}