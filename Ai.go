package main

type AiPlayer struct {
	boardPos BoardPosition
}

func (aiPlayer *AiPlayer) init(slice []int) {
	aiPlayer.boardPos.init(slice)
}

func (aiPlayer AiPlayer) TreeSearch() {

}