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
			boardPos.WhitePieces = append(boardPos.WhitePieces, Pawn{piecePosition, 1, true})
		case 1:
			boardPos.WhitePieces= append(boardPos.WhitePieces, King{piecePosition, 10, true})
		case 2:
			boardPos.WhitePieces= append(boardPos.WhitePieces, Queen{piecePosition, 9, true})
		case 3:
			boardPos.WhitePieces= append(boardPos.WhitePieces, Bishop{piecePosition, 3, true})
		case 4:
			boardPos.WhitePieces= append(boardPos.WhitePieces, Rook{piecePosition, 5, true})
		case 5:
			boardPos.WhitePieces= append(boardPos.WhitePieces, Knight{piecePosition, 3, true})

		case 6:
			boardPos.BlackPieces= append(boardPos.BlackPieces, Pawn{piecePosition, 1, false})
		case 7:
			boardPos.BlackPieces = append(boardPos.BlackPieces, King{piecePosition, 10, false})
		case 8:
			boardPos.BlackPieces = append(boardPos.BlackPieces, Queen{piecePosition, 9, false})
		case 9:
			boardPos.BlackPieces = append(boardPos.BlackPieces, Bishop{piecePosition, 3, false})
		case 10:
			boardPos.BlackPieces = append(boardPos.BlackPieces, Rook{piecePosition, 5, false})
		case 11:
			boardPos.BlackPieces = append(boardPos.BlackPieces, Knight{piecePosition, 3, false})
		}
	}
}