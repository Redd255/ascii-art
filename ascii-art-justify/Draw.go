package asciiart

import (
	"fmt"
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

	count := GetTerminalSize() - (GetArtSize(banner, inpultsplit))
	fmt.Println(count)

	if count % 10 != 0 {
		count = count +1
	}
	fmt.Println(count)

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
					if v[j] == ' '{
						for z := 0; z < lastCount-1; z++ {
							fmt.Print(" ")
						}
						continue
					}
					fmt.Print(banner[rune(v[j])][i])

				}
				fmt.Println()
			}
		}
	}
}
