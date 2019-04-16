package main

import (
	"fmt"

	"github.com/mbhk/learngo/fibo"
)

func main() {
	f := fibo.Fibonacci{}
	i := 0
	for i < 20 {
		i++
		fmt.Println(f.Next())
	}
}
