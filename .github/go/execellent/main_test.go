// コード4.2 テストコード
package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "even" {
		t.Errorf("exepected: even, actual: %s", result)
	}
}