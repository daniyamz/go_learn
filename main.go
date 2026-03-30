package main

import (
	"fmt"
	"go_learn/processor"
)

func main() {
	input := processor.GetInput()

	fmt.Println("*****Make a selection*****")
	fmt.Println(" 1. Last word")
	fmt.Println(" 2. Capitalize")
	fmt.Println(" 3. Delete Index")
	choice := processor.GetChoice()
	processor.Process(input, choice)
}
