package hello

import (
	"rsc.io/quote/v3"
)

// Hello resturns some greeting
func Hello() string {
	return quote.HelloV3()
}

// Proverb Some good advice
func Proverb() string {
	return quote.Concurrency()
}
