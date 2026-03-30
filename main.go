package main

import (
	"bufio"
	"fmt"
	"go_learn/processor"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
	begin:
		processor.ClearScreen()

		input := processor.GetInput(reader)
		if input == "exit" {
			fmt.Println("Press enter to exit...")
			reader.ReadString('\n')
		}
	start:
		fmt.Println("*****MENU*****")
		fmt.Println(" 1. Last word")
		fmt.Println(" 2. Capitalize")
		fmt.Println(" 3. Delete Index")
		choice := processor.GetChoice(reader)
		if choice != 1 && choice != 2 && choice != 3 {
			fmt.Println("\033[31mInvalid choice. Try again.\033[0m")
			goto start
		}
		processor.Process(input, choice)
		fmt.Println("Press Enter to continue...")
		reader.ReadString('\n')
		goto begin
	}
}
