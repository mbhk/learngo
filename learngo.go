package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mbhk/learngo/fibo"
)

func main() {
	var iterations uint
	flag.UintVar(&iterations, "iterations", 10, "number of iterations")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-iterations n]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	f := fibo.Fibonacci{}
	i := uint(0)
	for i < iterations {
		i++
		fmt.Println(f.Next())
	}
}
