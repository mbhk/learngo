package fibo

import "testing"

func TestNext(t *testing.T) {
	var f Fibonacci

	var known = []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377}
	t.Log("Doing a test of known numbers.")
	for i, fib := range known {
		t.Logf("checking %dth equal to %d", i, fib)
		if got := f.Next(); got.Uint64() != fib {
			t.Errorf("Next() = %q, want %q", got, fib)
		}
	}
}
