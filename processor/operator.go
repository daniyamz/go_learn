package processor

import (
	"fmt"
	"strings"
)

func GetInput() string {
	var input string
	fmt.Print("Enter Word: ")
	fmt.Scanln(&input)
	return input
}

func GetChoice() int {
	for {
		var choice int
		fmt.Print("Enter Digit: ")
		fmt.Scanln(&choice)

		if choice != 1 && choice != 2 && choice != 3 {
			fmt.Println("Wrong Selection. Try again.")
			continue
		}
		return choice
	}
}

func Process(GetInput string, Getchoice int) {
	for {
		switch Getchoice {
		case 1:
			last := string(GetInput[len(GetInput)-1])
			fmt.Println("Last word", last)
		case 2:
			fmt.Println("Capitalized", strings.ToUpper(GetInput))
		case 3:
			var index int
			fmt.Print("Enter index: ")
			fmt.Scanln(&index)
			char := []rune(GetInput)
			if index < 0 || index >= len(char) {
				fmt.Println("Error: Index out of range. Try again")
				continue
			}
			result := string(append(char[:index], char[index+1:]...))
			fmt.Println(result)
		}

	}
}
