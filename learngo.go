package main

import (
	"fmt"
	"os"

	"github.com/golang/example/stringutil"
	"github.com/mbhk/learngo/hello"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(hello.Hello())
	} else {
		fmt.Println(hello.Proverb())
		fmt.Printf("%v\n", os.Args)
	}
	fmt.Println(stringutil.Reverse(hello.Proverb()))
}
