package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"time"
)

var MAX_SORTING_DEPTH byte

type MoveAndEval struct {
	move Move
	eval float32
}

type MoveAndDepth struct {
	move Move
	maxDepth byte
}

type AiPlayer struct {
	boardPos BoardPosition
	bestMove MoveAndEval
	moveList []MoveAndEval
	count int
	firstMove bool
}

func (aiPlayer *AiPlayer) init(slice []int, nextMove bool) {
	aiPlayer.boardPos.init(slice, nextMove)
	aiPlayer.count = 0
	aiPlayer.firstMove = true
	aiPlayer.TreeSearch(aiPlayer.boardPos, 1, -10000, 10000, aiPlayer.boardPos.nextMove, MoveAndDepth{maxDepth:100})
	fmt.Println("Durchsuchte Positionen : ", aiPlayer.count)
}

func (aiPlayer *AiPlayer) TreeSearch(position BoardPosition, depth byte, alpha float32, beta float32, color bool, prevMove MoveAndDepth) float32{
	aiPlayer.count++
	posEval := eval(position)
	maxDepth := prevMove.maxDepth
	if depth == prevMove.maxDepth {
		return posEval
	}
	var newPos BoardPosition
	allMoves := allValidMoves(position, 1)
	start := time.Now()
	if depth == 1 {aiPlayer.SortMoveList(position, &allMoves, 6, color, false)}
	elapsed := time.Since(start)
	if depth == 1 {
		log.Printf("Binomial took %s", elapsed.Seconds())
		fmt.Println(allMoves)
	}
	if depth == 2 {aiPlayer.SortMoveList(position, &allMoves, 3, color, false)}
	if depth == 3 {aiPlayer.SortMoveList(position, &allMoves, 3, color, false)}
	if depth == 4 {aiPlayer.SortMoveList(position, &allMoves, 2, color, false)}
	if prevMove.maxDepth == 8 {
		if depth == 5 {aiPlayer.SortMoveList(position, &allMoves, 2, color, false)}
	}

	if len(allMoves) == 0 {return posEval}
	// if depth == 4 {aiPlayer.SortMoveList(position, &allMoves, 2, color)}
	// if color is white
	if position.nextMove {
		var maxEval float32 = -10000
		for index, move := range allMoves {
			if depth == 1 {
				if index < 2 {maxDepth = 8}
				if index >= 2 && index < 6 {maxDepth = 7}
				if index >= 6 && index < 10 {maxDepth = 6}
				if index >= 10 && index < 15 {maxDepth = 5}
				if index >= 15 {maxDepth = 5}
			}
			newPos = clone(position)
			newPos.movePiece(move.move)
			newPos.nextMove = !position.nextMove
			depthEval := aiPlayer.TreeSearch(newPos, depth+1, alpha, beta, !color, MoveAndDepth{move:move.move, maxDepth:maxDepth})

			if depthEval > maxEval {maxEval = depthEval}
			alpha = float32(math.Max(float64(alpha), float64(depthEval)))
			if beta <= alpha {break}
		}
		if depth == 2 {
			if maxEval < aiPlayer.bestMove.eval || aiPlayer.firstMove {
				aiPlayer.firstMove = false
				aiPlayer.bestMove = MoveAndEval{eval: maxEval, move:prevMove.move}
				fmt.Println(aiPlayer.bestMove)
			}
		}
		return maxEval
	}

	// if color is black
	if !position.nextMove {
		var minEval float32 = 10000
		for index, move := range allMoves {
			if depth == 1 {
				if index < 3 {maxDepth = 8}
				if index >= 3 && index < 6 {maxDepth = 7}
				if index >= 6 && index < 10 {maxDepth = 6}
				if index >= 10 && index < 15 {maxDepth = 5}
				if index >= 15 {maxDepth = 5}
			}
			newPos = clone(position)
			newPos.movePiece(move.move)
			newPos.nextMove = !position.nextMove
			depthEval := aiPlayer.TreeSearch(newPos, depth+1, alpha, beta, !color, MoveAndDepth{move:move.move, maxDepth:maxDepth})

			if depthEval < minEval {minEval = depthEval}
			beta = float32(math.Min(float64(beta), float64(depthEval)))
			if beta <= alpha {break}
		}
		if depth == 2 {
			if minEval > aiPlayer.bestMove.eval || aiPlayer.firstMove {
				aiPlayer.firstMove = false
				aiPlayer.bestMove = MoveAndEval{eval: minEval, move:prevMove.move}
			}
		}
		return minEval
	}
	return 0
}

// sorts moveList by TreeSearch of depth
func (aiPlayer *AiPlayer) SortMoveList(boardPos BoardPosition, unsortedMoveList *[]MoveAndEval, depth byte, color bool, onlyImproving bool) {
	aiPlayer.moveList = nil
	MAX_SORTING_DEPTH = depth
	posEval := eval(boardPos)
	aiPlayer.SortTreeSearch(boardPos, 1, -10000, 10000, color, Move{}, *unsortedMoveList)
	if color{
		// descending
		sort.SliceStable(aiPlayer.moveList, func(i, j int) bool {
			return aiPlayer.moveList[i].eval > aiPlayer.moveList[j].eval
		})
		if onlyImproving {
			*unsortedMoveList = nil
			for _, element := range aiPlayer.moveList {
				if element.eval >= posEval {*unsortedMoveList = append(*unsortedMoveList, element)} else {return}
			}
		}
	}
	if !color {
		// ascending
		sort.SliceStable(aiPlayer.moveList, func(i, j int) bool {
			return aiPlayer.moveList[i].eval < aiPlayer.moveList[j].eval
		})
		if onlyImproving {
			*unsortedMoveList = nil
			for _, element := range aiPlayer.moveList {
				if element.eval <= posEval {*unsortedMoveList = append(*unsortedMoveList, element)} else {return}
			}
		}
	}
	*unsortedMoveList = nil
	*unsortedMoveList = aiPlayer.moveList
}

func (aiPlayer *AiPlayer) SortTreeSearch(position BoardPosition, depth byte, alpha float32, beta float32, color bool, prevMove Move, moveList []MoveAndEval) float32{
	aiPlayer.count++
	posEval := eval(position)
	if depth == MAX_SORTING_DEPTH {
		if depth == 2 {
			aiPlayer.moveList = append(aiPlayer.moveList, MoveAndEval{eval:posEval, move:prevMove})
		}
		return posEval
	}
	var newPos BoardPosition
	var allMoves []MoveAndEval
	if depth == 1 {allMoves = moveList} else {allMoves = allValidMoves(position, 1)}
	// if color is white
	if position.nextMove {
		var maxEval float32 = -10000
		for _, move := range allMoves {
			newPos = clone(position)
			newPos.movePiece(move.move)
			newPos.nextMove = !position.nextMove
			depthEval := aiPlayer.SortTreeSearch(newPos, depth+1, alpha, beta, !color, move.move, nil)

			if depthEval > maxEval {maxEval = depthEval}
			alpha = float32(math.Max(float64(alpha), float64(depthEval)))
			if beta <= alpha {break}
		}
		if depth == 2 {
			aiPlayer.moveList = append(aiPlayer.moveList, MoveAndEval{eval:maxEval, move:prevMove})
		}
		return maxEval
	}

	// if color is black
	if !position.nextMove {
		var minEval float32 = 10000
		for _, move := range allMoves {
			newPos = clone(position)
			newPos.movePiece(move.move)
			newPos.nextMove = !position.nextMove
			depthEval := aiPlayer.SortTreeSearch(newPos, depth+1, alpha, beta, !color, move.move, nil)

			if depthEval < minEval {minEval = depthEval}
			beta = float32(math.Min(float64(beta), float64(depthEval)))
			if beta <= alpha {break}
		}
		if depth == 2 {
			aiPlayer.moveList = append(aiPlayer.moveList, MoveAndEval{eval:minEval, move:prevMove})
		}
		return minEval
	}
	return 0
}

func (aiPlayer AiPlayer) stringMove() string{
	start := time.Now()
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())
	randn := formatMove(aiPlayer.bestMove.move)
	retMove := fmt.Sprint(randn.start.x)+","+fmt.Sprint(randn.start.y)+","+fmt.Sprint(randn.end.x)+","+fmt.Sprint(randn.end.y)
	return retMove
}

func formatMove(move Move) Move {
	return Move{Point{move.start.x-1, 8-move.start.y}, Point{move.end.x-1, 8-move.end.y}}
}