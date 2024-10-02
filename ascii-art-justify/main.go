package main

import (
	"fmt"
	"os"
	"strings"

	asciiart "asciiart/src"
)

func main() {
	banner := "standard"
	var fileName string
	if len(os.Args) < 2 {
		fmt.Println("Error: need to write a one argument")
		return
	} else if (len(os.Args) == 2 || len(os.Args) == 3 )&& !strings.HasPrefix(os.Args[1], "--output=") && !strings.HasPrefix(os.Args[1], "--align=") {
		//EX : go run . [string] [banner]
		if len(os.Args) == 3 {
			banner = os.Args[2]
		}
		asciiart.Draw(banner, os.Args[1])
	} else if (len(os.Args) == 3 || len(os.Args) == 4) && strings.HasPrefix(os.Args[1], "--output=") && strings.HasSuffix(os.Args[1], ".txt") {
		//EX : go run . [output] [string] [banner]
		if os.Args[1] == "--output=.txt" {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
		fileName = strings.TrimPrefix(os.Args[1], "--output=")
		if len(os.Args) == 4 {
			banner = os.Args[3]
		}
		asciiart.DrawInFile(banner, os.Args[2], fileName)
	} else if (len(os.Args) == 3 || len(os.Args) == 4) && strings.HasPrefix(os.Args[1], "--align=") {
		//EX : go run . [align] [string] [banner]
		if len(os.Args) == 4 {
			banner = os.Args[3]
		}
		option := strings.TrimPrefix(os.Args[1], "--align=")
		switch option {
		case "left":
			asciiart.LeftText(banner, os.Args[2])
		case "right":
			asciiart.RightText(banner, os.Args[2])
		case "center":
			asciiart.CenterText(banner, os.Args[2])
		case "justify":
			asciiart.JustifyText(banner, os.Args[2])
		default:
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		}
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
	}
}
