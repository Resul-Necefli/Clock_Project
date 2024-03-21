package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

func main() {
	type clockArrey [5]string

	zero := clockArrey{
		"000",
		"0 0",
		"0 0",
		"0 0",
		"000",
	}

	one := clockArrey{
		" 0 ",
		"00 ",
		" 0 ",
		" 0 ",
		"000",
	}

	two := clockArrey{
		"000",
		"  0",
		"000",
		"0  ",
		"000",
	}

	three := clockArrey{
		"000",
		"  0",
		"000",
		"  0",
		"000",
	}
	four := clockArrey{
		"0 0",
		"0 0",
		"000",
		"  0",
		"  0",
	}
	five := clockArrey{
		"000",
		"0  ",
		"000",
		"  0",
		"000",
	}
	six := clockArrey{
		"000",
		"0  ",
		"000",
		"0 0",
		"000",
	}
	seven := clockArrey{
		"000",
		"  0",
		" 0 ",
		" 0 ",
		" 0 ",
	}
	eight := clockArrey{
		"000",
		"0 0",
		"000",
		"0 0",
		"000",
	}
	nine := clockArrey{
		"000",
		"0 0",
		"000",
		"  0",
		"000",
	}

	digits := [...]clockArrey{zero, one, two, three, four, five, six, seven, eight, nine}

	screen.Clear()
	for {
		now := time.Now()
		screen.MoveTopLeft()

		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		clean := clockArrey{
			"  ",
			"::",
			"  ",
			"::",
			"  ",
		}
		clock := [...]clockArrey{
			digits[hour/10], digits[hour%10],
			clean,
			digits[min/10], digits[min%10],
			clean,
			digits[sec/10], digits[sec%10],
		}
		c := color.New(color.FgRed)

		for i := 0; i < 5; i++ {
			for _, digit := range clock {
				for _, char := range digit[i] {
					if char == '0' {
						c.Printf("%s", string(char))
					} else {
						fmt.Printf("%s", string(char))
					}
				}
				fmt.Print("  ")
			}
			fmt.Println("")
		}
		time.Sleep(time.Second)
	}
}
