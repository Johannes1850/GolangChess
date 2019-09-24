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

	whiteKingMoved bool
	blackKingMoved bool
	RookA1Moved bool
	RookH1Moved bool
	RookA8Moved bool
	RookH8Moved bool

	// move that lead to position
	prevMove Move

	WhitePieces []Piece
	BlackPieces []Piece
	wholeBoard [8][8]Piece
}

func (boardPos *BoardPosition) removePiece(point Point, color bool)  {
	if color {
		for index, piece := range boardPos.WhitePieces {
			if piece.getPosition() == point {
				boardPos.WhitePieces[index] = boardPos.WhitePieces[len(boardPos.WhitePieces)-1]
				boardPos.WhitePieces =  boardPos.WhitePieces[:len(boardPos.WhitePieces)-1]
				return
			}
		}
	}
	if !color {
		for index, piece := range boardPos.BlackPieces {
			if piece.getPosition() == point {
				boardPos.BlackPieces[index] = boardPos.BlackPieces[len(boardPos.BlackPieces)-1]
				boardPos.BlackPieces =  boardPos.BlackPieces[:len(boardPos.BlackPieces)-1]
				return
			}
		}
	}
}

func (boardPos *BoardPosition) movePiece(move Move) {
	if boardPos.nextMove {
		for _, piece := range boardPos.WhitePieces {
			if piece.getPosition() == move.start {
				boardPos.removePiece(move.end, !boardPos.nextMove)
				// castling
				if piece.getValue() == 10 {
					if move.start.x-2 == move.end.x && !boardPos.RookA1Moved {
						boardPos.movePiece(Move{start:Point{1,1}, end:Point{4,1}})
						boardPos.RookA1Moved = true
						boardPos.whiteKingMoved = true
					}
					if move.start.x+2 == move.end.x && !boardPos.RookH1Moved {
						boardPos.movePiece(Move{start:Point{8,1}, end:Point{6,1}})
						boardPos.RookH1Moved = true
						boardPos.whiteKingMoved = true
					}
				}
				// pawn promotion
				if piece.getValue() == 1 && move.end.y == 8{
					boardPos.removePiece(move.start, boardPos.nextMove)
					boardPos.WhitePieces = append(boardPos.WhitePieces, &Queen{move.end, 9, true})
					return
				}
				piece.setPosition(move.end)
				return
			}
		}
	}
	if !boardPos.nextMove {
		for _, piece := range boardPos.BlackPieces {
			if piece.getPosition() == move.start {
				boardPos.removePiece(move.end, !boardPos.nextMove)
				// castling
				if piece.getValue() == 10 {
					if move.start.x-2 == move.end.x && !boardPos.RookA8Moved {
						boardPos.movePiece(Move{start:Point{1,8}, end:Point{4,8}})
						boardPos.RookA8Moved = true
						boardPos.blackKingMoved = true
					}
					if move.start.x+2 == move.end.x && !boardPos.RookH8Moved {
						boardPos.movePiece(Move{start:Point{8,8}, end:Point{6,8}})
						boardPos.RookH8Moved = true
						boardPos.blackKingMoved = true
					}
				}
				// pawn promotion
				if piece.getValue() == 1 && move.end.y == 1 {
					boardPos.removePiece(move.start, boardPos.nextMove)
					boardPos.BlackPieces = append(boardPos.BlackPieces, &Queen{move.end, 9, false})
					return
				}
				piece.setPosition(move.end)
				return
			}
		}
	}
}

func (boardPos *BoardPosition) init(slice []int, nextMove bool, posInfo [6]bool) {
/**
	// Load model
	module, _ := torch.LoadJITModule("EvalNN.pt")

	// Create an input tensor
	inputTensor, _ := torch.NewTensor(slice)

	// Forward propagation
	res, _ := module.Forward(inputTensor)

	fmt.Println("Dadada : ", res)
**/
	boardPos.nextMove = nextMove
	boardPos.whiteKingMoved = posInfo[0]
	boardPos.blackKingMoved = posInfo[1]
	boardPos.RookA1Moved = posInfo[2]
	boardPos.RookH1Moved = posInfo[3]
	boardPos.RookA8Moved = posInfo[4]
	boardPos.RookH8Moved = posInfo[5]
	for _, element := range slice {
		pieceInt := (element-1) / 64
		positionInt := element % 64
		piecePosition := posIntToPoint(positionInt)
		
		// <6 for white Pieces, >=6 for black Pieces
		switch pieceInt {
		case 0:
			boardPos.WhitePieces = append(boardPos.WhitePieces, &Pawn{piecePosition, 1, true})
		case 1:
			boardPos.WhitePieces= append(boardPos.WhitePieces, &King{piecePosition, 10, true})
		case 2:
			boardPos.WhitePieces= append(boardPos.WhitePieces, &Queen{piecePosition, 9, true})
		case 3:
			boardPos.WhitePieces= append(boardPos.WhitePieces, &Bishop{piecePosition, 4, true})
		case 4:
			boardPos.WhitePieces= append(boardPos.WhitePieces, &Rook{piecePosition, 5, true})
		case 5:
			boardPos.WhitePieces= append(boardPos.WhitePieces, &Knight{piecePosition, 3, true})

		case 6:
			boardPos.BlackPieces= append(boardPos.BlackPieces, &Pawn{piecePosition, 1, false})
		case 7:
			boardPos.BlackPieces = append(boardPos.BlackPieces, &King{piecePosition, 10, false})
		case 8:
			boardPos.BlackPieces = append(boardPos.BlackPieces, &Queen{piecePosition, 9, false})
		case 9:
			boardPos.BlackPieces = append(boardPos.BlackPieces, &Bishop{piecePosition, 4, false})
		case 10:
			boardPos.BlackPieces = append(boardPos.BlackPieces, &Rook{piecePosition, 5, false})
		case 11:
			boardPos.BlackPieces = append(boardPos.BlackPieces, &Knight{piecePosition, 3, false})
		}
	}
}