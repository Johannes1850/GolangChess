package main

import "math"

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
	for _, element := range boardPos.WhitePieces {
		if element.getPosition() == point { return true }
	}
	for _, element := range boardPos.BlackPieces {
		if element.getPosition() == point { return true }
	}
	return false
}

// returns true, if piece of given color exists at point
func pieceAtColor(boardPos BoardPosition, point Point, color bool) bool{
	if color == true {
		for _, element := range boardPos.WhitePieces {
			if element.getPosition() == point { return true }
		}
	}
	if color == false {
		for _, element := range boardPos.BlackPieces {
			if element.getPosition() == point { return true }
		}
	}
	return false
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

// returns false if way is blocked
func freeWay(boardPos BoardPosition, move Move) bool {
	// horizontal
	if move.start.y == move.end.y && move.start.x != move.end.x {
		horizontalDiff := int(move.start.x-move.end.x)
		if horizontalDiff < 0 {horizontalDiff *= -1}
		if move.start.x < move.end.x {
			for i := 1; i < horizontalDiff; i++ {
				if pieceAt(boardPos, Point{move.start.x+byte(i), move.start.y}){ return false }
			}
		}
		if move.start.x > move.end.x {
			for i := 1; i < horizontalDiff; i++ {
				if pieceAt(boardPos, Point{move.start.x-byte(i), move.start.y}){ return false }
			}
		}
	}

	// vertical
	if move.start.y != move.end.y && move.start.x == move.end.x {
		verticalDiff := int(move.start.y-move.end.y)
		if verticalDiff < 0 {verticalDiff *= -1}
		if move.start.y < move.end.y {
			for i := 1; i < verticalDiff; i++ {
				if pieceAt(boardPos, Point{move.start.x, move.start.y+byte(i)}){ return false }
			}
		}
		if move.start.y > move.end.y {
			for i := 1; i < verticalDiff; i++ {
				if pieceAt(boardPos, Point{move.start.x, move.start.y-byte(i)}){ return false }
			}
		}
	}

	//diagonal
	var pointX int = int(move.end.x - move.start.x)
	var pointY int = int(move.end.y - move.start.y)
	if math.Abs(float64(pointX)) == math.Abs(float64(pointY)) {
		var i byte
		// topRight
		if pointX > 0 && pointY > 0 {
			for i = 1; i < byte(pointY); i++ {
				if pieceAt(boardPos, Point{move.start.x+i, move.start.y+i}) { return false }
			}
		}
		// topLeft
		if pointX < 0 && pointY > 0 {
			for i = 1; i < byte(pointY); i++ {
				if pieceAt(boardPos, Point{move.start.x-i, move.start.y+i}) { return false }
			}
		}
		// bottomRight
		if pointX > 0 && pointY < 0 {
			for i = 1; i < byte(pointX); i++ {
				if pieceAt(boardPos, Point{move.start.x+i, move.start.y-i}) { return false }
			}
		}
		//bottomLeft
		if pointX < 0 && pointY < 0 {
			for i = 1; i < byte(math.Abs(float64(byte(pointX)))); i++ {
				if pieceAt(boardPos, Point{move.start.x-i, move.start.y-i}) { return false }
			}
		}
	}
	return true
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