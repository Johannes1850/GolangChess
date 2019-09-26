package main

import (
	"fmt"
	"math"
	"sort"
)

var MAX_SORTING_DEPTH byte

type MoveAndEval struct {
	move Move
	eval float32
}

type MoveListAndEval struct {
	moveList []Move
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
	moveSequence []MoveListAndEval
	moveSequence2 []MoveListAndEval
	bestDeepSearch MoveAndEval
	count int
	firstMove bool
	searchProgress DeepSearchProgression
}

// TODO moveSequence2 gets always changed in move sorting

func (aiPlayer *AiPlayer) init(slice []int, nextMove bool, posInfo [6]bool) {
	aiPlayer.boardPos.init(slice, nextMove, posInfo)
	aiPlayer.bestDeepSearch = MoveAndEval{}
	aiPlayer.count = 0
	aiPlayer.firstMove = true
	aiPlayer.moveSequence = []MoveListAndEval{}
	aiPlayer.TreeSearch(&aiPlayer.boardPos, 1, -10000, 10000, aiPlayer.boardPos.nextMove, MoveAndDepth{maxDepth:5}, []Move{})
	aiPlayer.StartDeepSearch()
	deepEval = aiPlayer.bestMove.eval
	fmt.Println("Durchsuchte Positionen : ", aiPlayer.count)
}

func (aiPlayer *AiPlayer) StartDeepSearch() {
	for _, moveSequence := range aiPlayer.moveSequence {
		aiPlayer.DeepSearch(&aiPlayer.boardPos, 2, 5, -10000, 10000, aiPlayer.boardPos.nextMove, moveSequence)
	}
	fmt.Println("hier")
	for _, moveSequence2 := range aiPlayer.moveSequence2 {
		if moveSequence2.eval < aiPlayer.bestMove.eval + 0.13 && len(moveSequence2.moveList) == 6{
			aiPlayer.DeepSearch(&aiPlayer.boardPos, 1, 4, -10000, 10000, aiPlayer.boardPos.nextMove, moveSequence2)
		}
	}
	/**
	fmt.Println("hier")
	for _, moveSequence2 := range aiPlayer.moveSequence2 {
		if moveSequence2.eval < aiPlayer.bestMove.eval + 0.07 && len(moveSequence2.moveList) == 8{
			aiPlayer.DeepSearch(&aiPlayer.boardPos, 1, 4, -10000, 10000, aiPlayer.boardPos.nextMove, moveSequence2)
		}
	}
	fmt.Println("hier")
	for _, moveSequence2 := range aiPlayer.moveSequence2 {
		if moveSequence2.eval < aiPlayer.bestMove.eval + 0.06 && len(moveSequence2.moveList) == 10{
			aiPlayer.DeepSearch(&aiPlayer.boardPos, 1, 4, -10000, 10000, aiPlayer.boardPos.nextMove, moveSequence2)
		}
	}
	**/
}

func (aiPlayer *AiPlayer) DeepSearch(position *BoardPosition, offset byte, depth byte, alpha float32, beta float32, color bool, moveList MoveListAndEval) {
	var newPos BoardPosition
	var currentMoveList []Move
	newPos = clone(*position)
	for i := 0; i < len(moveList.moveList)-int(offset); i++ {
		currentMoveList = append(currentMoveList, moveList.moveList[i])
		nextMove := moveList.moveList[i]
		newPos.movePiece(nextMove)
		newPos.nextMove = !newPos.nextMove
	}
	allMoves := allValidMoves(newPos, 1)
	aiPlayer.SortMoveList(newPos, &allMoves, depth, color, false, currentMoveList)
	if allMoves[0].eval <= aiPlayer.bestMove.eval {
		aiPlayer.bestMove.move = moveList.moveList[0]
		aiPlayer.bestMove.eval = allMoves[0].eval
		fmt.Println("Bester Zug : ", aiPlayer.bestMove)
	}
}

func (aiPlayer *AiPlayer) TreeSearch(position *BoardPosition, depth byte, alpha float32, beta float32, color bool, prevMove MoveAndDepth, moveList []Move) MoveListAndEval{
	aiPlayer.count++
	currentPos := *position
	posEval := eval(currentPos)
	if depth == prevMove.maxDepth {
		return MoveListAndEval{moveList:moveList, eval:posEval}
	}
	maxDepth := prevMove.maxDepth
	var newPos BoardPosition
	allMoves := allValidMoves(currentPos, 1)
	if allMoves[0].eval == 10 {
		if color {return MoveListAndEval{eval:1, moveList:moveList}} else {return MoveListAndEval{eval:-1, moveList:moveList}}
	}
	if depth == 1 {
		level1MoveCount = len(allMoves)
		searchedMoves = 0
		aiPlayer.SortMoveList(currentPos, &allMoves, 5, color, false, []Move{})
	}
	if depth == 2 {aiPlayer.SortMoveList(currentPos, &allMoves, 4, color, false, []Move{})}
	if depth == 3 {aiPlayer.SortMoveList(currentPos, &allMoves, 3, color, false, []Move{})}

	if len(allMoves) == 0 {return MoveListAndEval{moveList:moveList, eval:posEval}}
	if position.nextMove {
		var maxEval = MoveListAndEval{eval:-10000}
		for _, move := range allMoves {
			if maxDepth < 5 {maxDepth = 5}
			newPos = clone(currentPos)
			newPos.movePiece(move.move)
			newPos.nextMove = !position.nextMove
			var tmp []Move
			for _, move := range moveList {tmp = append(tmp, move)}
			tmp = append(tmp, move.move)
			depthEval := aiPlayer.TreeSearch(&newPos, depth+1, alpha, beta, !color, MoveAndDepth{move:move.move, maxDepth:maxDepth}, tmp)

			if depthEval.eval > maxEval.eval {maxEval = depthEval}
			alpha = float32(math.Max(float64(alpha), float64(depthEval.eval)))
			if beta <= alpha {break}
		}
		if depth == 2 {
			searchedMoves++
			if maxEval.eval <= aiPlayer.bestMove.eval+0.06 || aiPlayer.firstMove {
				aiPlayer.moveSequence = append(aiPlayer.moveSequence, maxEval)
				aiPlayer.firstMove = false
				aiPlayer.bestMove = MoveAndEval{eval: maxEval.eval, move:prevMove.move}
			}
		}
		return maxEval
	}

	// if color is black
	if !position.nextMove {
		var minEval = MoveListAndEval{eval: 10000}
		for _, move := range allMoves {
			if maxDepth < 5 {maxDepth = 5}
			newPos = clone(currentPos)
			newPos.movePiece(move.move)
			newPos.nextMove = !position.nextMove
			var tmp []Move
			for _, move := range moveList {tmp = append(tmp, move)}
			tmp = append(tmp, move.move)
			depthEval := aiPlayer.TreeSearch(&newPos, depth+1, alpha, beta, !color, MoveAndDepth{move:move.move, maxDepth:maxDepth}, tmp)

			if depthEval.eval < minEval.eval {minEval = depthEval}
			beta = float32(math.Min(float64(beta), float64(depthEval.eval)))
			if beta <= alpha {break}
		}
		if depth == 2 {
			searchedMoves++
			if minEval.eval >= aiPlayer.bestMove.eval-0.06 || aiPlayer.firstMove {
				aiPlayer.moveSequence = append(aiPlayer.moveSequence, minEval)
				aiPlayer.firstMove = false
				aiPlayer.bestMove = MoveAndEval{eval: minEval.eval, move:prevMove.move}
			}
		}
		return minEval
	}
	return MoveListAndEval{}
}

// sorts moveList by TreeSearch of depth
func (aiPlayer *AiPlayer) SortMoveList(boardPos BoardPosition, unsortedMoveList *[]MoveAndEval, depth byte, color bool, onlyImproving bool, moveSequence []Move) {
	aiPlayer.moveList = nil
	MAX_SORTING_DEPTH = depth
	posEval := eval(boardPos)
	if len(moveSequence) == 0 {
		aiPlayer.SortTreeSearch(boardPos, 1, -10000, 10000, color, Move{}, *unsortedMoveList)
	} else {
		aiPlayer.SortAndMakeSequence(boardPos, 1, -10000, 10000, color, Move{}, *unsortedMoveList, moveSequence)
	}
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

func (aiPlayer *AiPlayer) SortAndMakeSequence(position BoardPosition, depth byte, alpha float32, beta float32, color bool, prevMove Move, moveList []MoveAndEval, moveSequence []Move) MoveListAndEval{
	aiPlayer.count++
	posEval := eval(position)
	if depth == MAX_SORTING_DEPTH {
		if depth == 2 {
			aiPlayer.moveList = append(aiPlayer.moveList, MoveAndEval{eval:posEval, move:prevMove})
		}
		return MoveListAndEval{eval:posEval, moveList:moveSequence}
	}
	var newPos BoardPosition
	var allMoves []MoveAndEval
	if depth == 1 {allMoves = moveList} else {allMoves = allValidMoves(position, 1)}
	// if color is white
	if position.nextMove {
		var maxEval = MoveListAndEval{eval:-10000}
		for _, move := range allMoves {
			newPos = clone(position)
			newPos.movePiece(move.move)
			newPos.nextMove = !position.nextMove
			var tmp []Move
			for _, curMove := range moveSequence {tmp = append(tmp, curMove)}
			tmp = append(tmp, move.move)
			depthEval := aiPlayer.SortAndMakeSequence(newPos, depth+1, alpha, beta, !color, move.move, nil, tmp)

			if depthEval.eval > maxEval.eval {maxEval = depthEval}
			alpha = float32(math.Max(float64(alpha), float64(depthEval.eval)))
			if beta <= alpha {break}
		}
		if depth == 2 {
			if maxEval.eval <= aiPlayer.bestMove.eval+0.06 {
				aiPlayer.moveSequence2 = append(aiPlayer.moveSequence2, maxEval)
			}
			aiPlayer.moveList = append(aiPlayer.moveList, MoveAndEval{eval:maxEval.eval, move:prevMove})
		}
		return maxEval
	}

	// if color is black
	if !position.nextMove {
		var minEval = MoveListAndEval{eval:10000}
		for _, move := range allMoves {
			newPos = clone(position)
			newPos.movePiece(move.move)
			newPos.nextMove = !position.nextMove
			var tmp []Move
			for _, curMove := range moveSequence {tmp = append(tmp, curMove)}
			tmp = append(tmp, move.move)
			depthEval := aiPlayer.SortAndMakeSequence(newPos, depth+1, alpha, beta, !color, move.move, nil, tmp)

			if depthEval.eval < minEval.eval {minEval = depthEval}
			beta = float32(math.Min(float64(beta), float64(depthEval.eval)))
			if beta <= alpha {break}
		}
		if depth == 2 {
			if minEval.eval >= aiPlayer.bestMove.eval-0.06 {
				aiPlayer.moveSequence2 = append(aiPlayer.moveSequence2, minEval)
			}
			aiPlayer.moveList = append(aiPlayer.moveList, MoveAndEval{eval:minEval.eval, move:prevMove})
		}
		return minEval
	}
	return MoveListAndEval{}
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
	randn := formatMove(aiPlayer.bestMove.move)
	retMove := fmt.Sprint(randn.start.x)+","+fmt.Sprint(randn.start.y)+","+fmt.Sprint(randn.end.x)+","+fmt.Sprint(randn.end.y)
	return retMove
}

func formatMove(move Move) Move {
	return Move{Point{move.start.x-1, 8-move.start.y}, Point{move.end.x-1, 8-move.end.y}}
}