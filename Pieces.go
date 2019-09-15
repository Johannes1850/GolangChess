package main

type Piece interface {
	allMoves(boardPos BoardPosition) []Move
	validMove(boardPos BoardPosition, move Move) bool
	getPosition() Point
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
	// white Pawns
	if piece.color == true {
		// diagonal taking
		if move.start.y+1 == move.end.y && (move.start.x-1 == move.end.x || move.start.x+1 == move.end.x) && pieceAtColor(boardPos, move.end, !piece.color) {
			return true
		}
	}

	// black Pawns
	if piece.color == false {

	}
}

func (piece Pawn) allMoves(boardPos BoardPosition) []Move{
	 var retMoveList []Move

	return []Move{Move{start:Point{1,1}, end:Point{1,1}}}
}

func (piece Pawn) getPosition() Point{
	return piece.position
}


// King functions
func (piece King) validMove(boardPos BoardPosition, move Move) bool {
	if move.end.x > 8 || move.end.x < 1 || move.end.y > 8 || move.end.y < 1 {return false}
	return true
}

func (piece King) allMoves(boardPos BoardPosition) []Move{
	return []Move{Move{start:Point{1,1}, end:Point{1,1}}}
}

func (piece King) getPosition() Point{
	return piece.position
}


// Queen functions
func (piece Queen) validMove(boardPos BoardPosition, move Move) bool {
	if move.end.x > 8 || move.end.x < 1 || move.end.y > 8 || move.end.y < 1 {return false}
	return true
}

func (piece Queen) allMoves(boardPos BoardPosition) []Move{
	return []Move{Move{start:Point{1,1}, end:Point{1,1}}}
}

func (piece Queen) getPosition() Point{
	return piece.position
}


// Bishop functions
func (piece Bishop) validMove(boardPos BoardPosition, move Move) bool {
	if move.end.x > 8 || move.end.x < 1 || move.end.y > 8 || move.end.y < 1 {return false}
	return true
}

func (piece Bishop) allMoves(boardPos BoardPosition) []Move{
	return []Move{Move{start:Point{1,1}, end:Point{1,1}}}
}

func (piece Bishop) getPosition() Point{
	return piece.position
}


// Rook functions
func (piece Rook) validMove(boardPos BoardPosition, move Move) bool {
	if move.end.x > 8 || move.end.x < 1 || move.end.y > 8 || move.end.y < 1 {return false}
	return true
}

func (piece Rook) allMoves(boardPos BoardPosition) []Move{
	return []Move{Move{start:Point{1,1}, end:Point{1,1}}}
}

func (piece Rook) getPosition() Point{
	return piece.position
}


// Knight functions
func (piece Knight) validMove(boardPos BoardPosition, move Move) bool {
	if move.end.x > 8 || move.end.x < 1 || move.end.y > 8 || move.end.y < 1 {return false}
	return true
}

func (piece Knight) allMoves(boardPos BoardPosition) []Move{
	return []Move{Move{start:Point{1,1}, end:Point{1,1}}}
}

func (piece Knight) getPosition() Point{
	return piece.position
}