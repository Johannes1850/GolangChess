package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_DEPTH  = 4

type AiPlayer struct {
	boardPos BoardPosition

}

func (aiPlayer *AiPlayer) init(slice []int, nextMove bool) {
	aiPlayer.boardPos.init(slice, nextMove)
	aiPlayer.TreeSearch(aiPlayer.boardPos, 1, -10000, 10000, aiPlayer.boardPos.nextMove, Move{})
}

func (aiPlayer AiPlayer) TreeSearch(position BoardPosition, depth byte, alpha int, beta int, color bool, prevMove Move) float32{
	if depth == MAX_DEPTH {
		return eval(position)
	}
	var newPos BoardPosition
	allMoves := allValidMoves(aiPlayer.boardPos, 1)
	// if color is white
	if position.nextMove {
		for _, move := range allMoves {
			newPos = clone(position)
			newPos.movePiece(move)
			fmt.Println(position.BlackPieces)
			fmt.Println("newPos", newPos.BlackPieces)
		}
	}

	// if color is black
	if !position.nextMove {
		for _, move := range allMoves {
			fmt.Println(move)
			newPos = clone(position)
			newPos.movePiece(move)
			fmt.Println(position.BlackPieces)
			fmt.Println("newPos", newPos.BlackPieces)
		}
	}
	return 0
}

func (aiPlayer AiPlayer) stringMove() string{
	start := time.Now()
	allMoves := allValidMoves(aiPlayer.boardPos, 1)
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	randn := formatMove(allMoves[rand.Intn(len(allMoves))])
	retMove := fmt.Sprint(randn.start.x)+","+fmt.Sprint(randn.start.y)+","+fmt.Sprint(randn.end.x)+","+fmt.Sprint(randn.end.y)
	return retMove
}

func formatMove(move Move) Move {
	return Move{Point{move.start.x-1, 8-move.start.y}, Point{move.end.x-1, 8-move.end.y}}
}