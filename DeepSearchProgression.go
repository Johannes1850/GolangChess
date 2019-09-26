package main

import (
	"fmt"
	"strconv"
)

var level1MoveCount int
var searchedMoves int
var deepEval float32

type DeepSearchProgression struct {
	level2MoveCount int
	searchedMoves int
}

func getCalcProgression() string {
	if searchedMoves == 0 {return "0"}
	return strconv.Itoa(int(float32(searchedMoves)/float32(level1MoveCount)*100))
}

func getDeepEval() string {
	return fmt.Sprintf("%f", deepEval)
}
