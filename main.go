package main

import (
	"flag"
	"fmt"
)

func main() {
	flagTest := flag.String("wings", "w", "wings")

	flag.Parse()
	fmt.Println(flagTest)
	fmt.Println("Booting up!")
}
