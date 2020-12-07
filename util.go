package main

// Opposite gets the opposite of a piece/flips it over
func Opposite(piece Piece) Piece {
	if piece == White {
		return Black
	} else if piece == Black {
		return White
	}
	return Empty
}
