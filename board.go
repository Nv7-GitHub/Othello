package main

// Board contains the data for an othello board
type Board struct {
	Turn   bool
	Pieces [8][8]Piece // 8x8
}

// NewBoard initializes an Othello board
func NewBoard() Board {
	return Board{}
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
