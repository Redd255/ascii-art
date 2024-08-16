package asciiart

import (
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

func CountSpaces(input string) int {
	count := 0
	for _, c := range input {
		if c == 32 {
			count++
		}
	}
	return count
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
