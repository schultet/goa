package main

import (
	"fmt"
	"os"

	"golang.org/x/exp/ebnf"
)

func main() {
	fmt.Println("hi")
	f, _ := os.Open("eb.txt")
	g, err := ebnf.Parse("./eb.txt", f)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%+v\n", g)
}
