package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const MAX_DEPTH  = 5

type MoveAndEval struct {
	move Move
	eval float32
}

type AiPlayer struct {
	boardPos BoardPosition
	bestMove MoveAndEval
	count int
	firstMove bool
}

func (aiPlayer *AiPlayer) init(slice []int, nextMove bool) {
	aiPlayer.boardPos.init(slice, nextMove)
	aiPlayer.count = 0
	aiPlayer.firstMove = true
	aiPlayer.TreeSearch(aiPlayer.boardPos, 1, -10000, 10000, aiPlayer.boardPos.nextMove, Move{})
}

func (aiPlayer *AiPlayer) TreeSearch(position BoardPosition, depth byte, alpha float32, beta float32, color bool, prevMove Move) float32{
	aiPlayer.count++
	if depth == MAX_DEPTH {
		return eval(position)
	}
	var newPos BoardPosition
	allMoves := allValidMoves(position, 1)
	// if color is white
	if position.nextMove {
		var maxEval float32 = -10000
		for _, move := range allMoves {
			newPos = clone(position)
			newPos.movePiece(move)
			newPos.nextMove = !position.nextMove
			depthEval := aiPlayer.TreeSearch(newPos, depth+1, alpha, beta, color, move)

			if depthEval > maxEval {maxEval = depthEval}
			alpha = float32(math.Max(float64(alpha), float64(depthEval)))
			if beta <= alpha {break}
		}
		if depth == 2 {
			fmt.Println(maxEval, aiPlayer.bestMove.eval)
			if maxEval < aiPlayer.bestMove.eval || aiPlayer.firstMove {
				aiPlayer.firstMove = false
				aiPlayer.bestMove = MoveAndEval{eval: maxEval, move:prevMove}
				fmt.Println(aiPlayer.bestMove)
			}
		}
		return maxEval
	}

	// if color is black
	if !position.nextMove {
		var minEval float32 = 10000
		for _, move := range allMoves {
			newPos = clone(position)
			newPos.movePiece(move)
			newPos.nextMove = !position.nextMove
			depthEval := aiPlayer.TreeSearch(newPos, depth+1, alpha, beta, color, move)

			if depthEval < minEval {minEval = depthEval}
			beta = float32(math.Min(float64(beta), float64(depthEval)))
			if beta <= alpha {break}
		}
		if depth == 2 {
			if minEval > aiPlayer.bestMove.eval || aiPlayer.firstMove {
				aiPlayer.firstMove = false
				aiPlayer.bestMove = MoveAndEval{eval: minEval, move:prevMove}
				fmt.Println(aiPlayer.bestMove)
			}
		}
		return minEval
	}
	return 0
}

func (aiPlayer AiPlayer) stringMove() string{
	start := time.Now()
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	randn := formatMove(aiPlayer.bestMove.move)
	retMove := fmt.Sprint(randn.start.x)+","+fmt.Sprint(randn.start.y)+","+fmt.Sprint(randn.end.x)+","+fmt.Sprint(randn.end.y)
	return retMove
}

func formatMove(move Move) Move {
	return Move{Point{move.start.x-1, 8-move.start.y}, Point{move.end.x-1, 8-move.end.y}}
}