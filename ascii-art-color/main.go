package main

import (
	"flag"
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
	} else if (len(os.Args) == 2 || len(os.Args) == 3) && !strings.HasPrefix(os.Args[1], "--output=") &&
		!strings.HasPrefix(os.Args[1], "--align=") && !strings.HasPrefix(os.Args[1], "--color=") {
		// EX : go run . [string] [banner]
		if len(os.Args) == 3 {
			banner = os.Args[2]
		}
		asciiart.Draw(banner, os.Args[1])
	} else if (len(os.Args) == 3 || len(os.Args) == 4) && strings.HasPrefix(os.Args[1], "--output=") && strings.HasSuffix(os.Args[1], ".txt") {
		// EX : go run . [output] [string] [banner]
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
		// EX : go run . [align] [string] [banner]
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
	} else if (len(os.Args) == 3 || len(os.Args) == 4 || len(os.Args) == 5) && strings.HasPrefix(os.Args[1], "--color=") {
		// EX : go run . [color] [string] [banner]
		if !strings.ContainsAny(os.Args[1], "=") {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\"")
			return
		}
		colorFlag := flag.String("color", "", "Specify color for the substring or the entire string")
		flag.Parse()
		args := flag.Args()
		if len(args) == 1 || (len(args) == 2 && os.Args[3] == "shadow" || os.Args[3] == "standard" || os.Args[3] == "thinkertoy") {
			if len(args) == 2 {
				banner = args[1]
			}
			asciiart.DrawColor(*colorFlag, args[0], banner)
		} else {
			// EX : go run . [color] [sub] [string] [banner]
			if len(args) == 3 {
				banner = args[2]
			}
			asciiart.DrawSubColor(*colorFlag, args[1], args[0], banner)
		}
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\"")
	}
}
