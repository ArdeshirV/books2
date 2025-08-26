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

func TestTestSub(t *testing.T) {
	a, b := 20, 10
	exp := a - b
	res := testSub(a, b)

	if res != exp {
		t.Errorf("Expected testSub(%d, %d) = %d, but got %d", a, b, exp, res)
	}
}

func BenchmarkTestAdd(b *testing.B) {
	x, y := 10, 20
	for range b.N {
		testAdd(x, y)
	}
}
