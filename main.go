package main

import (
	"fmt"
	"go_learn/processor"
)

func main() {
	processor.Scanner.Scan()
	for {
		processor.ClearScreen()

		input := processor.GetInput()
		if input == "exit" {
			fmt.Println("Exiting Program...")
			//reader.ReadString('\n')
			break
		}
		fmt.Println("*****MENU*****")
		fmt.Println(" 1. Last word")
		fmt.Println(" 2. Capitalize")
		fmt.Println(" 3. Delete Index")
		choice := processor.GetChoice()
		if choice != 1 && choice != 2 && choice != 3 {
			fmt.Println("\033[31mInvalid choice. Try again.\033[0m")

		}
		processor.Process(input, choice)
		fmt.Println("Press Enter to continue...")
	}
}
