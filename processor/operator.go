package processor

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Scanner = bufio.NewScanner(os.Stdin)

func GetInput() string {
	//reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[33mEnter Sentence:\033[0m ")
	Scanner.Scan()
	input := Scanner.Text()
	return strings.TrimSpace(input)
}

func GetChoice() int {
	fmt.Print("\033[34mChoose an option:\033[0m ")
	Scanner.Scan()
	input := Scanner.Text()
	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}

	return num
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func Process(GetInput string, Getchoice int) {
	switch Getchoice {
	case 1:
		last := string(GetInput[len(GetInput)-1])
		fmt.Println("Last character", last)
	case 2:
		fmt.Println("Capitalized", strings.ToUpper(GetInput))
	case 3:
		for {
			fmt.Print("\033[36mEnter index:\033[0m ")
			reader := bufio.NewReader(os.Stdin)
			indexst, _ := reader.ReadString('\n')
			indexst = strings.TrimSpace(indexst)
			index, err := strconv.Atoi(indexst)
			if err != nil {
				fmt.Println("Invalid input.")
				fmt.Println("Press Enter to continue...")
				reader.ReadString('\n')
				continue
			}
			char := []rune(GetInput)
			if index < 0 || index >= len(char) {
				fmt.Println("\033[31mError: Index out of range. Try again\033[0m ")
				fmt.Println("Press Enter to continue...")
				reader.ReadString('\n')
				continue
			}
			result := string(append(char[:index], char[index+1:]...))
			fmt.Println(result)
		}
	}
}
