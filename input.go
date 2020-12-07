package main

import "fmt"

// Input gets board input, checks if it is valid, and changes the board
func (b *Board) Input() {
	var input string
	if b.Turn {
		fmt.Print("")
	}
	fmt.Scanln(&input)

}

func isValidInput(inp string) bool {
	return true
}
