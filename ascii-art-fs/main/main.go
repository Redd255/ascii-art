package main

import (
	"asciiart"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("need to write a one argument")
		return
	}
	if !asciiart.Checkchars(os.Args[1]) {
		fmt.Println("Input should contain only printable ASCII characters.")
		return
	}
	baname := "standard"
	if len(os.Args) > 2 {
		input2 := os.Args[2]
		switch input2 {
		case "standard":
			baname = "standard"
		case "shadow":
			baname = "shadow"
		case "thinkertoy":
			baname = "thinkertoy"
		default:
			fmt.Println("Invalid banner name")
			return
		}
		banner, err := asciiart.MapBanner(baname + ".txt")
		if err != nil {
			fmt.Println("Error loading banner:", err)
			return
		}
		input := os.Args[1]
		inputsplit := strings.Split(input, "\\n")
		asciiart.Draw(banner, inputsplit)
	} else {
		banner, err := asciiart.MapBanner(baname + ".txt")
		if err != nil {
			fmt.Println("Error loading banner:", err)
			return
		}
		input := os.Args[1]
		inputsplit := strings.Split(input, "\\n")
		asciiart.Draw(banner, inputsplit)
	}
}
