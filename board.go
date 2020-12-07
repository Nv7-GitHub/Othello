package main

// Board contains the data for an othello board
type Board struct {
	Turn   bool
	Pieces [][]bool // 8x8
}
