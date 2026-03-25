package main

import (
	"fmt"
	"go_learn/processor"
)

func main() {
	input := processor.GetInput()

	fmt.Println("Make a selection: ")
	choice := processor.GetChoice()
	processor.Process(input, choice)
}
