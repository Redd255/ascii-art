package asciiart

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Checkchars(s string) bool {
	for _, c := range s {
		if c < 32 || c > 126 {
			return false
		}
	}
	return true
}

func GetTerminalSize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	s := strings.TrimSpace(string(out))
	ssp := strings.Split(s, " ")
	width, err := strconv.Atoi(ssp[1])
	if err != nil {
		log.Fatal(err)
	}
	return width
}
func GetArtSize(banner map[rune][]string, inpultsplit []string) int {
	count := 0
	for _, v := range inpultsplit {
		if len(v) != 0 && !Checknewline(inpultsplit) {
			for i := 0; i < 1; i++ {
				for j := 0; j < len(v); j++ {
					count += len(banner[rune(v[j])][i])
				}
			}
		}
	}
	return count
}
func CenterText(banner map[rune][]string, inpultsplit []string) {
	count := GetTerminalSize()/2 - (GetArtSize(banner, inpultsplit))/2
	fmt.Println(count, (GetArtSize(banner, inpultsplit)))
	for idx, v := range inpultsplit {
		if Checknewline(inpultsplit) && idx != len(inpultsplit)-1 {
			fmt.Println()
			continue
		} else if len(v) == 0 && !Checknewline(inpultsplit) {
			fmt.Println()
		} else if len(v) != 0 && !Checknewline(inpultsplit) {
			for i := 0; i < 8; i++ {
				for c := 0; c < count-1; c++ {
					fmt.Print(string(32))
				}
				for j := 0; j < len(v); j++ {
					fmt.Print(banner[rune(v[j])][i])
				}
				fmt.Println()
			}
		}
	}
}
func RightText(banner map[rune][]string, inpultsplit []string) {
	count := GetTerminalSize() - (GetArtSize(banner, inpultsplit))
	fmt.Println(count, (GetArtSize(banner, inpultsplit)))
	for idx, v := range inpultsplit {
		if Checknewline(inpultsplit) && idx != len(inpultsplit)-1 {
			fmt.Println()
			continue
		} else if len(v) == 0 && !Checknewline(inpultsplit) {
			fmt.Println()
		} else if len(v) != 0 && !Checknewline(inpultsplit) {
			for i := 0; i < 8; i++ {
				for c := 0; c < count-1; c++ {
					fmt.Print(string(32))
				}
				for j := 0; j < len(v); j++ {
					fmt.Print(banner[rune(v[j])][i])
				}
				fmt.Println()
			}
		}
	}
}

func LeftText(banner map[rune][]string, inpultsplit []string,input string) {
	banner, err := MapBanner(baname + ".txt")
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}
	inputsplit := strings.Split(input, "\\n")
	Draw(banner, inputsplit)
}
func CountSpaces(input string) int {
	count := 0
	for _, c := range input {
		if c == 32 {
			count++
		}
	}
	return count
}
func JustifyText(banner map[rune][]string, inpultsplit []string, input string) {
	if CountSpaces(input) == 0 {
		fmt.Println("No spaces found in input.")
		return
	}
	if CountSpaces(input) == 1 {
		RightText()
	}
	count := GetTerminalSize() - (GetArtSize(banner, inpultsplit))
	lastCount := count / CountSpaces(input)
	for idx, v := range inpultsplit {
		if Checknewline(inpultsplit) && idx != len(inpultsplit)-1 {
			fmt.Println()
			continue
		} else if len(v) == 0 && !Checknewline(inpultsplit) {
			fmt.Println()
		} else if len(v) != 0 && !Checknewline(inpultsplit) {
			for i := 0; i < 8; i++ {
				for c := 0; c < lastCount-1; c++ {
					fmt.Print(string(32))
				}
				for j := 0; j < len(v); j++ {
					fmt.Print(banner[rune(v[j])][i])
				}
				fmt.Println()
			}
		}
	}
}
func MapBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	banner := make(map[rune][]string)
	charIndex := 32
	for i := 0; i < len(lines); i += 9 {
		if i+8 < len(lines) {
			banner[rune(charIndex)] = lines[i+1 : i+9]
			charIndex++
		}
	}
	return banner, nil
}

func Checknewline(inpultsplit []string) bool {
	c := 0
	for _, line := range inpultsplit {
		if len(line) != 0 {
			c++
		}
	}
	if c == 0 {
		return true
	} else {
		return false
	}
}

func Draw(banner map[rune][]string, inpultsplit []string) {
	for idx, v := range inpultsplit {
		if Checknewline(inpultsplit) && idx != len(inpultsplit)-1 {
			fmt.Println()
			continue
		} else if len(v) == 0 && !Checknewline(inpultsplit) {
			fmt.Println()
		} else if len(v) != 0 && !Checknewline(inpultsplit) {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(v); j++ {
					fmt.Print(banner[rune(v[j])][i])
				}
				fmt.Println()
			}
		}
	}
}
