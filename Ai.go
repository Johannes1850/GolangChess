package main

import "fmt"
import "time"

const MAX_DEPTH  = 4

type AiPlayer struct {
	boardPos BoardPosition
}

func (aiPlayer *AiPlayer) init(slice []int, nextMove bool) {
	aiPlayer.boardPos.init(slice, nextMove)
	aiPlayer.TreeSearch(aiPlayer.boardPos, 1, -10000, 10000, aiPlayer.boardPos.nextMove, Move{})
}

func (aiPlayer AiPlayer) TreeSearch(position BoardPosition, depth byte, alpha int, beta int, color bool, prevMove Move) {
	fmt.Println(color)
	time.Sleep(3 * time.Second)
}

func (aiPlayer AiPlayer) stringMove() string{
	return "1,7,1,5"
}