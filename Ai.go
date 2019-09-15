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

func (aiPlayer AiPlayer) TreeSearch(position BoardPosition, depth byte, alpha int, beta int, color bool, prevMove Move) {
	// time.Sleep(3 * time.Second)
}

func (aiPlayer AiPlayer) stringMove() string{
	allMoves := allValidMoves(aiPlayer.boardPos, 1)
	fmt.Println(allMoves)
	fmt.Println(len(allMoves))
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	randn := formatMove(allMoves[rand.Intn(len(allMoves))])
	retMove := fmt.Sprint(randn.start.x)+","+fmt.Sprint(randn.start.y)+","+fmt.Sprint(randn.end.x)+","+fmt.Sprint(randn.end.y)
	return retMove
}

func formatMove(move Move) Move {
	return Move{Point{move.start.x-1, 8-move.start.y}, Point{move.end.x-1, 8-move.end.y}}
}