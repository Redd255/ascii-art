package asciiart

import "fmt"

func Draw(banner map[rune][]string, inpultsplit []string) { // default drawing
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

func LeftText(banner map[rune][]string, inpultsplit []string) { // left drawing
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

func RightText(banner map[rune][]string, inpultsplit []string) { //right drawing
	count := GetTerminalSize() - (GetArtSize(banner, inpultsplit))-1 // terminal size - ascii art size -1= how many spaces we'll add 
	for idx, v := range inpultsplit {
		if Checknewline(inpultsplit) && idx != len(inpultsplit)-1 {
			fmt.Println()
			continue
		} else if len(v) == 0 && !Checknewline(inpultsplit) {
			fmt.Println()
		} else if len(v) != 0 && !Checknewline(inpultsplit) {
			for i := 0; i < 8; i++ {
				for c := 0; c < count; c++ {
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

func CenterText(banner map[rune][]string, inpultsplit []string) { // center drawing
	count := GetTerminalSize()/2 - (GetArtSize(banner, inpultsplit))/2 -1  //terminal size /2 - ascii art size/2 -1= how many spaces we'll add 
	for idx, v := range inpultsplit {
		if Checknewline(inpultsplit) && idx != len(inpultsplit)-1 {
			fmt.Println()
			continue
		} else if len(v) == 0 && !Checknewline(inpultsplit) {
			fmt.Println()
		} else if len(v) != 0 && !Checknewline(inpultsplit) {
			for i := 0; i < 8; i++ {
				for c := 0; c < count; c++ {
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

func JustifyText(banner map[rune][]string, inpultsplit []string, input string) { // justfy
	count := GetTerminalSize() - (GetArtSize(banner, inpultsplit))
	var lastCount int
	if CountSpaces(input) == 0 {
		lastCount = 0 
	}else if CountSpaces(input) >0 {
		lastCount = (count / CountSpaces(input))-1// the number of spaces we'll add between words
	}
	for idx, v := range inpultsplit {
		if Checknewline(inpultsplit) && idx != len(inpultsplit)-1 {
			fmt.Println()
			continue
		} else if len(v) == 0 && !Checknewline(inpultsplit) {
			fmt.Println()
		} else if len(v) != 0 && !Checknewline(inpultsplit) {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(v); j++ {
					if v[j] == ' ' {
						for z := 0; z < lastCount; z++ {
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
