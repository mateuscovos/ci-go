package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(10, 10)

	if total != 20 {
		t.Errorf("The result of the sum was invalid. Result %d. Expected: %d.", total, 20)
	}
}
