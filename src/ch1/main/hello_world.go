package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args))
	if len(os.Args) >= 1 {
		fmt.Println("Hello World" + os.Args[0])
	}
}
