package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func Draw(banner string, input string) { // default drawing
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
						fmt.Print(maps[rune(v[j])][i])
					}
					fmt.Println()
				}
			}
		}
	} else {
		fmt.Println("Input contains special characters")
	}
}

func DrawInFile(banner string, input string, outputPath string) error {
	if banner != "standard" && banner != "thinkertoy" && banner != "shadow" {
		return fmt.Errorf("invalid banner name")
	}
	maps, err := MapBanner("banners/" + banner + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Create("result/" + outputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	if Checkchars(input) {
		inputsplit := strings.Split(input, "\\n")
		for idx, v := range inputsplit {
			if Checknewline(inputsplit) && idx != len(inputsplit)-1 {
				_, err := file.WriteString("\n")
				if err != nil {
					return err
				}
				continue
			} else if len(v) == 0 && !Checknewline(inputsplit) {
				_, err := file.WriteString("\n")
				if err != nil {
					return err
				}
			} else if len(v) != 0 && !Checknewline(inputsplit) {
				for i := 0; i < 8; i++ {
					for j := 0; j < len(v); j++ {
						_, err := file.WriteString(maps[rune(v[j])][i])
						if err != nil {
							return err
						}
					}
					_, err := file.WriteString("\n")
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	} else {
		return fmt.Errorf("input contains special characters")
	}
}

func LeftText(banner string, input string) { // left drawing
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
						fmt.Print(maps[rune(v[j])][i])
					}
					fmt.Println()
				}
			}
		}
	} else {
		fmt.Println("Input contains special characters")
	}
}

func RightText(banner string, input string) { //right drawing
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
		count := GetTerminalSize() - (GetArtSize(banner, inputsplit)) - 1 // terminal size - ascii art size -1= how many spaces we'll add
		for idx, v := range inputsplit {
			if Checknewline(inputsplit) && idx != len(inputsplit)-1 {
				fmt.Println()
				continue
			} else if len(v) == 0 && !Checknewline(inputsplit) {
				fmt.Println()
			} else if len(v) != 0 && !Checknewline(inputsplit) {
				for i := 0; i < 8; i++ {
					for c := 0; c < count; c++ {
						fmt.Print(string(32))
					}
					for j := 0; j < len(v); j++ {
						fmt.Print(maps[rune(v[j])][i])
					}
					fmt.Println()
				}
			}
		}
	}
}

func CenterText(banner string, input string) { // center drawing
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
		count := GetTerminalSize()/2 - (GetArtSize(banner, inputsplit))/2 - 1 //terminal size /2 - ascii art size/2 -1= how many spaces we'll add
		for idx, v := range inputsplit {
			if Checknewline(inputsplit) && idx != len(inputsplit)-1 {
				fmt.Println()
				continue
			} else if len(v) == 0 && !Checknewline(inputsplit) {
				fmt.Println()
			} else if len(v) != 0 && !Checknewline(inputsplit) {
				for i := 0; i < 8; i++ {
					for c := 0; c < count; c++ {
						fmt.Print(string(32))
					}
					for j := 0; j < len(v); j++ {
						fmt.Print(maps[rune(v[j])][i])
					}
					fmt.Println()
				}
			}
		}
	}
}

func JustifyText(banner string, input string) { // justfy
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
		count := GetTerminalSize() - (GetArtSize(banner, inputsplit))
		var lastCount int
		if CountSpaces(input) == 0 {
			lastCount = 0
		} else if CountSpaces(input) > 0 {
			lastCount = (count / CountSpaces(input)) - 1 // the number of spaces we'll add between words
		}
		for idx, v := range inputsplit {
			if Checknewline(inputsplit) && idx != len(inputsplit)-1 {
				fmt.Println()
				continue
			} else if len(v) == 0 && !Checknewline(inputsplit) {
				fmt.Println()
			} else if len(v) != 0 && !Checknewline(inputsplit) {
				for i := 0; i < 8; i++ {
					for j := 0; j < len(v); j++ {
						if v[j] == ' ' {
							for z := 0; z < lastCount; z++ {
								fmt.Print(" ")
							}
							continue
						}
						fmt.Print(maps[rune(v[j])][i])
					}
					fmt.Println()
				}
			}
		}
	}
}
