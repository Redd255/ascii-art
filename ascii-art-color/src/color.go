package asciiart

import (
	"fmt"
	"strings"
)

var colors = map[string]string{
	"Reset": "\033[0m",
	"red":   "\033[31m",
}

func DrawColor(color string, input string, banner string) {
	if banner != "standard" && banner != "thinkertoy" && banner != "shadow" {
		fmt.Println("Invalid banner name")
		return
	}
	maps, err := MapBanner("banners/" + banner + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	if Checkchars(input) {
		inputsplit := strings.Split(input, "\\n")
		for idx, v := range inputsplit {
			if Checknewline(inputsplit) && idx != len(inputsplit)-1 {
				fmt.Println()
				continue
			} else if len(v) == 0 && !Checknewline(inputsplit) {
				fmt.Println()
			} else if len(v) != 0 && !Checknewline(inputsplit) {
				for i := 0; i < 8; i++ {
					for j := 0; j < len(v); j++ {
						fmt.Print(colors[color] + maps[rune(v[j])][i] + colors["Reset"])
					}
					fmt.Println()
				}
			}
		}
	} else {
		fmt.Println("Input contains special characters")
	}
}

func DrawSubColor(color string, input string, sub string, banner string) {
	if banner != "standard" && banner != "thinkertoy" && banner != "shadow" {
		fmt.Println("Invalid banner name")
		return
	}

	maps, err := MapBanner("banners/" + banner + ".txt")
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}

	// check if the color exists
	colorCode, ok := colors[color]
	if !ok {
		fmt.Println("Invalid color:", color)
		return
	}

	inputsplit := strings.Split(input, "\\n")
	for idx, v := range inputsplit {
		if Checknewline(inputsplit) && idx != len(inputsplit)-1 {
			fmt.Println()
			continue
		} else if len(v) == 0 && !Checknewline(inputsplit) {
			fmt.Println()
		} else if len(v) != 0 && !Checknewline(inputsplit) {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(v); j++ {
					if strings.Contains(sub, string(v[j])) {
						fmt.Print("\033[" + colorCode + maps[rune(v[j])][i] + colors["Reset"]) // Example for different highlight color
					} else {
						fmt.Print(maps[rune(v[j])][i])
					}
				}
				fmt.Println()
			}
		}
	}
}
