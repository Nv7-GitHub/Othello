package main

import "fmt"

// Board contains the data for an othello board
type Board struct {
	Turn   bool
	Pieces [8][8]Piece // 8x8
}

// NewBoard initializes an Othello board
func NewBoard() Board {
	b := Board{}
	b.Pieces[3][3] = White
	b.Pieces[4][3] = Black
	b.Pieces[4][4] = White
	b.Pieces[3][4] = Black
	return b
}

// Print prints out the Othello board
func (b *Board) Print() {
	for _, row := range b.Pieces {
		for _, piece := range row {
			if piece == White {
				fmt.Print("W ")
			} else if piece == Black {
				fmt.Print("B ")
			} else if piece == Empty {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}

// Piece is a position on an Othello Board
type Piece int

const (
	// Empty has an empty piece (*)
	Empty Piece = 0
	// Black is a black piece (B)
	Black Piece = 1
	// White is a white piece (W)
	White Piece = 2
)
