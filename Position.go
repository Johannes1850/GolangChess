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
				piece.setPosition(move.end)
				return
			}
		}
	}
	if !boardPos.nextMove {
		for _, piece := range boardPos.BlackPieces {
			if piece.getPosition() == move.start {
				boardPos.removePiece(move.end, !boardPos.nextMove)
				piece.setPosition(move.end)
				return
			}
		}
	}
}

func (boardPos *BoardPosition) init(slice []int, nextMove bool) {
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
			boardPos.WhitePieces= append(boardPos.WhitePieces, &Bishop{piecePosition, 3, true})
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
			boardPos.BlackPieces = append(boardPos.BlackPieces, &Bishop{piecePosition, 3, false})
		case 10:
			boardPos.BlackPieces = append(boardPos.BlackPieces, &Rook{piecePosition, 5, false})
		case 11:
			boardPos.BlackPieces = append(boardPos.BlackPieces, &Knight{piecePosition, 3, false})
		}
	}
}