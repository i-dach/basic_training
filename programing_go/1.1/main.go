package main

import (
	"fmt"
	"os"
	"strings"
)

func practice1() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func main() {
	practice1()
}
