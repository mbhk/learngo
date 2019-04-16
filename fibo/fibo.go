/*
Package fibo implements a simple library for Fibonacci numbers.
*/
package fibo

import (
	"math/big"
)

var initial big.Int

// Fibonacci holds the current state of a Fibonacci sequence
// initialize fresh one using New
type Fibonacci struct {
	a *big.Int
	b *big.Int
}

// Next returns the next Fibonacci number
func (m *Fibonacci) Next() *big.Int {
	if m.a == nil || m.b == nil {
		m.a = &initial
		m.b = big.NewInt(1)
	}

	if m.a == &initial {
		m.a = copy(&initial)
	} else {
		m.a.Add(m.a, m.b)
		// Swap a and b so that b is the next number in the sequence.
		m.a, m.b = m.b, m.a
	}

	return copy(m.a)
}

// copy creates a deep copy of an *big.Int
func copy(a *big.Int) *big.Int {
	result := big.Int{}
	return result.Set(a)
}
