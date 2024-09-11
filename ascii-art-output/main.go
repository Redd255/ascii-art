package main

import (
	"fmt"
	"os"
	"strings"

	asciiart "asciiart/src"
)

func main() {
	forlder := "banners/"
	if len(os.Args) == 1 {
		fmt.Println("Error: need to write one argument or more.")
		return
	} else if len(os.Args) == 2 && !strings.HasPrefix(os.Args[1], "--output") { // ex: go run . [STRING] 
		input := os.Args[1]
		if !asciiart.Checkchars(input) {
			fmt.Println("Error: Input should contain only printable ASCII characters.")
			return
		}
		banner, err := asciiart.MapBanner(forlder+"standard.txt")
		if err != nil {
			fmt.Println(" Error: loading banner:", err)
			return
		}
		inputsplit := strings.Split(input, "\\n")
		asciiart.Draw(banner, inputsplit)
	} else if len(os.Args) == 3 && !strings.HasPrefix(os.Args[1], "--output") { //ex: go run . [STRING] [BANNER]
		if !asciiart.Checkchars(os.Args[1]) {
			fmt.Println("Error: Input should contain only printable ASCII characters.")
			return
		}
		input := os.Args[1]
		bannerName := os.Args[2]
		banner, err := asciiart.MapBanner(forlder+bannerName + ".txt")
		if err != nil {
			fmt.Println(" Error loading banner:\n", err)
			return
		}
		inputsplit := strings.Split(input, "\\n")
		asciiart.Draw(banner, inputsplit)
	} else if strings.HasPrefix(os.Args[1], "--output") { // ex: go run . [OPTION] [STRING] [BANNER]
		fileName := os.Args[1]
		errorTxt := "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
		if len(os.Args) < 3 || len(os.Args) > 4 {
			fmt.Println(errorTxt)
			return
		}
		input := os.Args[2]
		bannerName := "standard"
		if len(os.Args) == 4 {
			bannerName = os.Args[3]
		}
		if len(os.Args) > 2 && !asciiart.Checkchars(os.Args[2]) {
			fmt.Println("Input should contain only printable ASCII characters.")
			return
		}
		if os.Args[1] != "--output=.txt" && (strings.HasPrefix(os.Args[1], "--output=") && strings.HasSuffix(os.Args[1], ".txt")) {
			fileName = strings.TrimPrefix(fileName, "--output=")
		} else if os.Args[1] == "--output=.txt" {
			fmt.Println(errorTxt)
			return
		} else {
			fmt.Println(errorTxt)
			return
		}
		banner, err := asciiart.MapBanner(forlder+bannerName + ".txt")
		if err != nil {
			fmt.Println(" Error loading banner:\n", err)
			return
		}
		inputsplit := strings.Split(input, "\\n")
		asciiart.DrawInFile(banner, inputsplit, fileName)
	} else {
		fmt.Println("Error: Invalid number of arguments.")
		return
	}
}
