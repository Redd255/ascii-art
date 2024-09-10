package main

import (
	"fmt"
	"os"
	"strings"

	asciiart "asciiart/src"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}
	if !asciiart.Checkchars(os.Args[1]) {
		fmt.Println("Input should contain only printable ASCII characters.")
		return
	}
	baname := "standard" // this is the default banner
	folder := "banners/" // thsi is if the user want to change the name banner's folder
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
		banner, err := asciiart.MapBanner(folder + baname + ".txt")
		if err != nil {
			fmt.Println("Error loading banner:", err)
			return
		}
		input := os.Args[1]
		inputsplit := strings.Split(input, "\\n")
		asciiart.Draw(banner, inputsplit)
	} else if len(os.Args) == 2{ // if the len of arg are not
		banner, err := asciiart.MapBanner(folder + baname + ".txt")
		if err != nil {
			fmt.Println("Error loading banner:", err)
			return
		}
		input := os.Args[1]
		inputsplit := strings.Split(input, "\\n")
		asciiart.Draw(banner, inputsplit)
	} else { // if the len of arg are not 3
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
	}
}
