package main

import (
	"fmt"
	"os"
)

func practice1() []string {
	return os.Args[0:]
}

func main() {
	fmt.Println(practice1(), " ")
}
