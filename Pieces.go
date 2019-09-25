package main

import (
	"math"
)

type Piece interface {
	allMoves(boardPos BoardPosition) []MoveAndEval
	validMove(boardPos BoardPosition, move Move) bool
	getPosition() Point
	getValue() byte
	setPosition(point Point)
	clone() Piece
}

type Pawn struct {
	position Point
	value byte
	color bool
}

type Knight struct {
	position Point
	value byte
	color bool
}

type Bishop struct {
	position Point
	value byte
	color bool
}

type Rook struct {
	position Point
	value byte
	color bool
}

type Queen struct {
	position Point
	value byte
	color bool
}

type King struct {
	position Point
	value byte
	color bool
}

// Pawn functions
func (piece Pawn) validMove(boardPos BoardPosition, move Move) bool {
	if move.end.x > 8 || move.end.x < 1 || move.end.y > 8 || move.end.y < 1 {return false}
	if pieceAtColor(boardPos, move.end, piece.color) {return false}
	if !freeWay(boardPos, move) {return false}
	// white Pawns
	if piece.color == true {
		// diagonal taking
		if move.start.y+1 == move.end.y && (move.start.x-1 == move.end.x || move.start.x+1 == move.end.x) && pieceAtColor(boardPos, move.end, !piece.color) {
			return true
		}
		if pieceAtColor(boardPos, move.end, !piece.color) { return false }
		// hasn't been moved yet
		if move.start.y == 2 {
			if (move.start.y + 1 == move.end.y || move.start.y + 2 == move.end.y) && move.start.x == move.end.x {
				return true
			}
		} else {
			if move.start.y + 1 == move.end.y && move.start.x == move.end.x {
				return true
			}
		}
	}

	// black Pawns
	if piece.color == false {
		// diagonal taking
		if move.start.y-1 == move.end.y && (move.start.x-1 == move.end.x || move.start.x+1 == move.end.x) && pieceAtColor(boardPos, move.end, !piece.color) {
			return true
		}
		if pieceAtColor(boardPos, move.end, !piece.color) { return false }
		// hasn't been moved yet
		if move.start.y == 7 {
			if (move.start.y - 1 == move.end.y || move.start.y - 2 == move.end.y) && move.start.x == move.end.x {
				return true
			}
		} else {
			if move.start.y - 1 == move.end.y && move.start.x == move.end.x {
				return true
			}
		}
	}
	return false
}

func (piece Pawn) allMoves(boardPos BoardPosition) []MoveAndEval{
	var retMoveList []MoveAndEval
	var move Move
	if piece.color {
		move = Move{piece.position, Point{piece.position.x+1, piece.position.y+1}}
		if piece.validMove(boardPos, move) {
		 	retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		}
		move = Move{piece.position, Point{piece.position.x-1, piece.position.y+1}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		}
		move := Move{piece.position, Point{piece.position.x, piece.position.y+1}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval:0})
		}
		move = Move{piece.position, Point{piece.position.x, piece.position.y+2}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval:0})
		}
	}

	if !piece.color {
		move = Move{piece.position, Point{piece.position.x+1, piece.position.y-1}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		}
		move = Move{piece.position, Point{piece.position.x-1, piece.position.y-1}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		}
		move := Move{piece.position, Point{piece.position.x, piece.position.y-1}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval:0})
		}
		move = Move{piece.position, Point{piece.position.x, piece.position.y-2}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval:0})
		}
	}
	return retMoveList
}

func (piece *Pawn) setPosition(point Point) {
	piece.position = point
}

func (piece Pawn) getPosition() Point {
	return piece.position
}

func (piece Pawn) getValue() byte {
	return piece.value
}

func (piece Pawn) clone() Piece {
	var newPawn Pawn
	newPawn.position.x = piece.position.x
	newPawn.position.y = piece.position.y
	newPawn.color = piece.color
	newPawn.value = piece.value
	return &newPawn
}


// King functions
func (piece King) validMove(boardPos BoardPosition, move Move) bool {
	if move.end.x > 8 || move.end.x < 1 || move.end.y > 8 || move.end.y < 1 {return false}
	if pieceAtColor(boardPos, move.end, piece.color) {return false}
	// castling
	if move.start.x - move.end.x == 2 || move.end.x - move.start.x == 2 {
		if piece.color == true {
			if !boardPos.whiteKingMoved && piece.position.x == 5 && piece.position.y == 1 {
				if move.start.x - 2 == move.end.x && freeWay(boardPos, Move{start:Point{x:5,y:1}, end:Point{3,1}}) {if !boardPos.RookA1Moved && freeWay(boardPos, Move{start:Point{x:1,y:1}, end:Point{4,1}}) {return true}}
				if move.start.x + 2 == move.end.x && freeWay(boardPos, Move{start:Point{x:5,y:1}, end:Point{7,1}}) {if !boardPos.RookH1Moved && freeWay(boardPos, Move{start:Point{x:8,y:1}, end:Point{6,1}}) {return true}}
			}
		}
		if piece.color == false {
			if !boardPos.blackKingMoved && piece.position.x == 5 && piece.position.y == 8 {
				if move.start.x - 2 == move.end.x && freeWay(boardPos, Move{start:Point{x:5,y:8}, end:Point{3,8}}) {if !boardPos.RookA8Moved && freeWay(boardPos, Move{start:Point{x:1,y:8}, end:Point{4,8}}) {return true}}
				if move.start.x + 2 == move.end.x && freeWay(boardPos, Move{start:Point{x:5,y:8}, end:Point{7,8}}) {if !boardPos.RookH8Moved && freeWay(boardPos, Move{start:Point{x:8,y:8}, end:Point{6,8}}) {return true}}
			}
		}
	}
	var diffX = int(move.start.x) - int(move.end.x)
	var diffY = int(move.start.y) - int(move.end.y)
	if math.Abs(float64(diffX)) <= 1 && math.Abs(float64(diffY)) <= 1 {
		// if !KingBlockingKing(boardPos, move.end) {return true}
		return true
	}
	return false
}

func (piece King) allMoves(boardPos BoardPosition) []MoveAndEval{
	var retMoveList []MoveAndEval
	var move Move
	move = Move{piece.position, Point{piece.position.x+2, piece.position.y}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x-2, piece.position.y}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x+1, piece.position.y+1}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x+1, piece.position.y}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x+1, piece.position.y-1}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x-1, piece.position.y+1}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x-1, piece.position.y}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x-1, piece.position.y-1}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x, piece.position.y+1}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x, piece.position.y-1}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	return retMoveList
}

func (piece *King) setPosition(point Point) {
	piece.position = point
}

func (piece King) getPosition() Point{
	return piece.position
}

func (piece King) getValue() byte {
	return piece.value
}

func (piece King) clone() Piece {
	var newPiece King
	newPiece.position.x = piece.position.x
	newPiece.position.y = piece.position.y
	newPiece.color = piece.color
	newPiece.value = piece.value
	return &newPiece
}


// Queen functions
func (piece Queen) validMove(boardPos BoardPosition, move Move) bool {
	if move.end.x > 8 || move.end.x < 1 || move.end.y > 8 || move.end.y < 1 {return false}
	if pieceAtColor(boardPos, move.end, piece.color) {return false}
	var diffX int = int(move.start.x) - int(move.end.x)
	var diffY int = int(move.start.y) - int(move.end.y)
	if math.Abs(float64(diffX)) == math.Abs(float64(diffY)) || (move.start.x == move.end.x &&
		move.start.y != move.end.y || move.start.x != move.end.x && move.start.y == move.end.y) {
		if freeWay(boardPos, move) {return true}
	}
	return false
}

func (piece Queen) allMoves(boardPos BoardPosition) []MoveAndEval{
	var retMoveList []MoveAndEval
	var move Move
	for i := 1; i <= 8; i++ {
		move = Move{piece.position, Point{piece.position.x+byte(i), piece.position.y+byte(i)}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	for i := 1; i <= 8; i++ {
		move = Move{piece.position, Point{piece.position.x+byte(i), piece.position.y-byte(i)}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	for i := 1; i <= 8; i++ {
		move = Move{piece.position, Point{piece.position.x-byte(i), piece.position.y+byte(i)}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	for i := 1; i <= 8; i++ {
		move = Move{piece.position, Point{piece.position.x-byte(i), piece.position.y-byte(i)}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	for i := 1; i <= 7; i++ {
		move = Move{piece.position, Point{piece.position.x, piece.position.y+byte(i)}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	for i := 1; i <= 7; i++ {
		move = Move{piece.position, Point{piece.position.x, piece.position.y-byte(i)}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	for i := 1; i <= 7; i++ {
		move = Move{piece.position, Point{piece.position.x+byte(i), piece.position.y}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	for i := 1; i <= 7; i++ {
		move = Move{piece.position, Point{piece.position.x-byte(i), piece.position.y}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	return retMoveList
}

func (piece *Queen) setPosition(point Point) {
	piece.position = point
}

func (piece Queen) getPosition() Point{
	return piece.position
}

func (piece Queen) getValue() byte {
	return piece.value
}

func (piece Queen) clone() Piece {
	var newPiece Queen
	newPiece.position.x = piece.position.x
	newPiece.position.y = piece.position.y
	newPiece.color = piece.color
	newPiece.value = piece.value
	return &newPiece
}


// Bishop functions
func (piece Bishop) validMove(boardPos BoardPosition, move Move) bool {
	if move.end.x > 8 || move.end.x < 1 || move.end.y > 8 || move.end.y < 1 {return false}
	if pieceAtColor(boardPos, move.end, piece.color) {return false}
	var diffX int = int(move.start.x) - int(move.end.x)
	var diffY int = int(move.start.y) - int(move.end.y)
	if math.Abs(float64(diffX)) == math.Abs(float64(diffY)) {
		if freeWay(boardPos, move) {return true}
	}
	return false
}

func (piece Bishop) allMoves(boardPos BoardPosition) []MoveAndEval{
	var retMoveList []MoveAndEval
	var move Move
	for i := 1; i <= 8; i++ {
		move = Move{piece.position, Point{piece.position.x+byte(i), piece.position.y+byte(i)}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	for i := 1; i <= 8; i++ {
		move = Move{piece.position, Point{piece.position.x+byte(i), piece.position.y-byte(i)}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	for i := 1; i <= 8; i++ {
		move = Move{piece.position, Point{piece.position.x-byte(i), piece.position.y+byte(i)}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	for i := 1; i <= 8; i++ {
		move = Move{piece.position, Point{piece.position.x-byte(i), piece.position.y-byte(i)}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	return retMoveList
}

func (piece *Bishop) setPosition(point Point) {
	piece.position = point
}

func (piece Bishop) getPosition() Point{
	return piece.position
}

func (piece Bishop) getValue() byte {
	return piece.value
}

func (piece Bishop) clone() Piece {
	var newPiece Bishop
	newPiece.position.x = piece.position.x
	newPiece.position.y = piece.position.y
	newPiece.color = piece.color
	newPiece.value = piece.value
	return &newPiece
}


// Rook functions
func (piece Rook) validMove(boardPos BoardPosition, move Move) bool {
	if move.end.x > 8 || move.end.x < 1 || move.end.y > 8 || move.end.y < 1 {return false}
	if pieceAtColor(boardPos, move.end, piece.color) {return false}
	if move.start.x == move.end.x &&
		move.start.y != move.end.y || move.start.x != move.end.x && move.start.y == move.end.y {
		if freeWay(boardPos, move) {return true}
	}
	return false
}

func (piece Rook) allMoves(boardPos BoardPosition) []MoveAndEval{
	var retMoveList []MoveAndEval
	var move Move
	for i := 1; i <= 7; i++ {
		move = Move{piece.position, Point{piece.position.x, piece.position.y+byte(i)}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	for i := 1; i <= 7; i++ {
		move = Move{piece.position, Point{piece.position.x, piece.position.y-byte(i)}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	for i := 1; i <= 7; i++ {
		move = Move{piece.position, Point{piece.position.x+byte(i), piece.position.y}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	for i := 1; i <= 7; i++ {
		move = Move{piece.position, Point{piece.position.x-byte(i), piece.position.y}}
		if piece.validMove(boardPos, move) {
			retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
		} else {break}
	}
	return retMoveList
}

func (piece *Rook) setPosition(point Point) {
	piece.position = point
}

func (piece Rook) getPosition() Point{
	return piece.position
}

func (piece Rook) getValue() byte {
	return piece.value
}

func (piece Rook) clone() Piece {
	var newPiece Rook
	newPiece.position.x = piece.position.x
	newPiece.position.y = piece.position.y
	newPiece.color = piece.color
	newPiece.value = piece.value
	return &newPiece
}


// Knight functions
func (piece Knight) validMove(boardPos BoardPosition, move Move) bool {
	if move.end.x > 8 || move.end.x < 1 || move.end.y > 8 || move.end.y < 1 {return false}
	if pieceAtColor(boardPos, move.end, piece.color) {return false}
	diffX := int(move.start.x)-int(move.end.x)
	diffY := int(move.start.y)-int(move.end.y)
	if math.Abs(float64(diffX)) == 2 && math.Abs(float64(diffY)) == 1 || (math.Abs(float64(diffX)) == 1 && math.Abs(float64(diffY)) == 2) { return true }
	return false
}

func (piece Knight) allMoves(boardPos BoardPosition) []MoveAndEval{
	var retMoveList []MoveAndEval
	var move Move
	move = Move{piece.position, Point{piece.position.x+2, piece.position.y+1}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x+2, piece.position.y-1}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x+1, piece.position.y+2}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x+1, piece.position.y-2}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x-2, piece.position.y+1}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x-2, piece.position.y-1}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x-1, piece.position.y+2}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	move = Move{piece.position, Point{piece.position.x-1, piece.position.y-2}}
	if piece.validMove(boardPos, move) {
		retMoveList = append(retMoveList, MoveAndEval{move:move, eval: float32(pieceAtValue(boardPos, move.end, !piece.color))})
	}
	return retMoveList
}

func (piece *Knight) setPosition(point Point) {
	piece.position = point
}

func (piece Knight) getPosition() Point{
	return piece.position
}

func (piece Knight) getValue() byte {
	return piece.value
}

func (piece Knight) clone() Piece {
	var newPiece Knight
	newPiece.position.x = piece.position.x
	newPiece.position.y = piece.position.y
	newPiece.color = piece.color
	newPiece.value = piece.value
	return &newPiece
}