package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello vent!")
	args := os.Args
	fmt.Println("count: ", len(args))
	for _, arg := range args[1:] {
		fmt.Println(arg)
	}
}
