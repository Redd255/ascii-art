package main

import (
	asciiart "asciiart/Func"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("need to write a one argument")
		return
	}
	banner, err := asciiart.MapBanner("standard.txt")
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}
	input := os.Args[1]
	if asciiart.Checkchars(input) {
		inputsplit := strings.Split(input, "\\n")
		asciiart.Draw(banner, inputsplit)
	} else {
		fmt.Println("Input should contain only printable ASCII characters.")
	}
}
