package main

import (
	. "database/sql/driver"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func startCalc(slice []int, nextMove bool, posInfo [6]bool) string{
	var aiPlayer AiPlayer
	aiPlayer.init(slice, nextMove, posInfo)
	return aiPlayer.stringMove()
}

func stringToBool(bool string) bool{
	if bool == "false" {
		return false
	} else { return true }
}

func receiveAjax(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nextMove := stringToBool(r.FormValue("next_move"))
		whiteKingMoved := stringToBool(r.FormValue("whiteKingMoved"))
		blackKingMoved := stringToBool(r.FormValue("blackKingMoved"))
		rookA1Moved := stringToBool(r.FormValue("rookA1Moved"))
		rookH1Moved:= stringToBool(r.FormValue("rookH1Moved"))
		rookA8Moved := stringToBool(r.FormValue("rookA8Moved"))
		rookH8Moved := stringToBool(r.FormValue("rookH8Moved"))

		boardPosition := r.FormValue("board_position")
		cleaned := strings.Replace(boardPosition, ",", " ", -1)
		strSlice := strings.Fields(cleaned)
		// create new slice with boolean's
		var intSlice []int
		for i := 0; i < len(strSlice); i++ {
			if strSlice[i] == "1" {
				intSlice = append(intSlice, i)
			}
		}
		aiMove := startCalc(intSlice, nextMove, [6]bool{whiteKingMoved, blackKingMoved, rookA1Moved, rookH1Moved, rookA8Moved, rookH8Moved})
		w.Write([]byte(aiMove))
	}
}

func sendDeepEvaluation(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(getDeepEval()))
}

func sendCalcProgression (w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(getCalcProgression()))
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	_ = t.Execute(w, Null{})
	fmt.Println("hier")
}

func main() {
	r := mux.NewRouter()
	r.
		PathPrefix("/js/").
		Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("."+"/js/"))))
	r.
		PathPrefix("/images/").
		Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("."+"/images/"))))
	r.HandleFunc("/", handler)
	r.HandleFunc("/receive", receiveAjax)
	r.HandleFunc("/calcProgress", sendCalcProgression)
	r.HandleFunc("/deepEval", sendDeepEvaluation)
	log.Fatal(http.ListenAndServe(":8080", r))
}