package asciiart

import (
	"fmt"
	"strings"
)

var colors = map[string]string{
	"Reset": "\033[0m",
	"nlack":   "\033[30m",
	"red":   "\033[31m",
	"green": "\033[32m",
	"yellow": "\033[33m",
	"blue": "\033[34m",
	"magenta": "\033[35m",
	"cyan": "\033[36m",
	"white":   "\033[37m",
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
    for idx, line := range inputsplit {
        if Checknewline(inputsplit) && idx != len(inputsplit)-1 {
            fmt.Println()
            continue
        } else if len(line) == 0 && !Checknewline(inputsplit) {
            fmt.Println()
        } else if len(line) != 0 && !Checknewline(inputsplit) {
            asciiRows := [8]string{}
            // Process each word in the line
            words := strings.Split(line, " ")
            for _, word := range words {
                for i := 0; i < 8; i++ {
                    for j := 0; j < len(word); j++ {
                        char := rune(word[j])
                        subIndex := strings.Index(word, sub)

                        if subIndex != -1 && j >= subIndex && j < subIndex+len(sub) {
                            // Add colored substring to the appropriate row
                            asciiRows[i] += colorCode + maps[char][i] + colors["Reset"]
                        } else {
                            // Add the rest of the word normally
                            asciiRows[i] += maps[char][i]
                        }
                    }

                    // Add a space between words (except for the last word)
                    if word != words[len(words)-1] {
                        asciiRows[i] += "      " 
                    }
                }
            }

            // Print each row of the resulting ASCII art
            for i := 0; i < 8; i++ {
                fmt.Println(asciiRows[i])
            }
        }
    }
}
