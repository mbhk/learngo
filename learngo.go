package main

import (
	"fmt"
	"os"

	"github.com/mbhk/learngo/hello"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(hello.Hello())
	} else {
		fmt.Println(hello.Proverb())
	}
}
