package main

type Point struct {
	x byte
	y byte
}

type Move struct {
	start Point
	end Point
}

type BoardPosition struct {
	// white is true, black is false
	nextMove bool

	WhitePieces []Piece
	BlackPieces []Piece
	wholeBoard [8][8]Piece
}

func (boardPos *BoardPosition) init(slice []int) {

}