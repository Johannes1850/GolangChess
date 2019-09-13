package main

type MoveSortingMode byte
const (
	NoSorting        MoveSortingMode = 0
	BestMovesFirst   MoveSortingMode = 1
	OnlyHittingMoves MoveSortingMode = 2
)

// BoardPosition functions

func eval(boardPos BoardPosition) float32{
	return -0.3
}

// returns true, if piece exists at point
func pieceAt(boardPos BoardPosition, point Point) bool{
	return true
}

// returns piece at point, else emptyPiece
func getPiece(boardPos BoardPosition, point Point) Piece{
	a := Pawn{}
	return a
}

// returns all valid moves for given position
func allValidMoves(boardPos BoardPosition, sortingMode MoveSortingMode) []Move{
	return []Move{Move{start:Point{1,1}, end:Point{1,1}}}
}