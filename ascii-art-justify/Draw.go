package asciiart

import (
	"fmt"
	"strings"
)

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

func LeftText(banner map[rune][]string, inpultsplit []string) {
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
func JustifyText(banner map[rune][]string, inpultsplit []string, input string) {
	if CountSpaces(input) >= 0 {
		
		LeftText(banner, inpultsplit)
		return
	}
	if CountSpaces(input) == 1 {
		sp := strings.Split(inpultsplit[0], " ")
		inpultsplit[0] = sp[1]
		RightText(banner, inpultsplit)
		return
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

				for j := 0; j < len(v); j++ {
					fmt.Print(banner[rune(v[j])][i])
				}
				for c := 0; c < lastCount-1; c++ {
					fmt.Print(string(32))
				}
				fmt.Println()
			}
		}
	}
}
