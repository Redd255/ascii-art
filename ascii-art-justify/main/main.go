package main

import (
	"asciiart"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: need to write a one argument")
		return
	} else if len(os.Args) > 4 {
		fmt.Println("Error: Too many arguments")
		return
	} else {
		baname := "standard"
		if len(os.Args) > 3 {
			input2 := os.Args[3]
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
		}
		banner, err := asciiart.MapBanner(baname + ".txt")
		if err != nil {
			fmt.Println("Error loading banner:", err)
			return
		}
		if strings.HasPrefix(os.Args[1], "--align=") {
			// For Usage: go run . [OPTION] [STRING] [BANNER]
			input := os.Args[2]
			var inputsplit []string
			if asciiart.Checkchars(input) {
				inputsplit = strings.Split(input, "\\n")
			} else if !asciiart.Checkchars(input) {
				fmt.Println("Input should contain only printable ASCII characters.")
				return
			}
			alignName := "left"
			alignName = strings.TrimPrefix(os.Args[1], "--align=")
			if alignName == "left" {
				asciiart.LeftText(banner, inputsplit)
			} else if alignName == "right" {
				asciiart.RightText(banner, inputsplit)
			} else if alignName == "center" {
				asciiart.CenterText(banner, inputsplit)
			} else if alignName == "justify" {
				asciiart.JustifyText(banner, inputsplit,input)
			} else {
				fmt.Println("Invalid alignment option")
				return
			}
		}
	}
}
