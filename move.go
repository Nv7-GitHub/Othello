package main

type change struct {
	row int
	col int
}

var changes = []change{
	change{row: 0, col: -1},
	change{row: 1, col: 0},
	change{row: 0, col: 1},
	change{row: -1, col: 0},
}

// Move flips pieces as needed, provided with the piece that just got changed
func (b *Board) Move(row, col int) {
	for _, change := range changes {
		if (b.Pieces[row-change.row][col-change.col] == Opposite(b.Pieces[row][col])) && (b.Pieces[row-change.row*2][col-change.col*2] == b.Pieces[row][col]) {
			b.Pieces[row-change.row][col-change.col] = b.Pieces[row][col]
		}
	}
}
