package main

import "testing"

func TestTestAdd(t *testing.T) {
	a, b := 10, 20
	exp := a + b
	res := testAdd(a, b)

	if res != exp {
		t.Errorf("Expected testAdd(%d, %d) = %d, but got %d", a, b, exp, res)
	}
}
