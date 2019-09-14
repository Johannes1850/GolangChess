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

// returns a clone of boardPos
func clone(boardPos BoardPosition) BoardPosition {
	newPos := boardPos
	newPos.WhitePieces = make([]Piece, len(boardPos.WhitePieces))
	copy(newPos.WhitePieces, boardPos.WhitePieces)
	return newPos
}

// returns board coordinate from 1-768 coordinates
func posIntToPoint(posInt int) Point{
	if posInt == 0 {
		return Point{8,8}
	}
	var helpY float32 = (float32)(posInt) / 8.0
	var posY byte
	if helpY <= 1 {
		posY = 1
	}
	if helpY > 1 && helpY <= 2 {
		posY = 2
	}
	if helpY > 2 && helpY <= 3 {
		posY = 3
	}
	if helpY > 3 && helpY <= 4 {
		posY = 4
	}
	if helpY > 4 && helpY <= 5 {
		posY = 5
	}
	if helpY > 5 && helpY <= 6 {
		posY = 6
	}
	if helpY > 6 && helpY <= 7 {
		posY = 7
	}
	if helpY > 7 {
		posY = 8
	}
	var posX byte = byte(posInt) - (posY-1) * 8
	return Point{posX, posY}
}