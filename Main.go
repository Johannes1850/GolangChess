package main

import (
	. "database/sql/driver"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func startCalc(slice []int, nextMove bool) string{
	var aiPlayer AiPlayer
	aiPlayer.init(slice, nextMove)
	return aiPlayer.stringMove()
}

func receiveAjax(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		next_move_string := r.FormValue("next_move")
		var next_move bool
		if next_move_string == "false" {
			next_move = false
		} else { next_move = true }
		board_position := r.FormValue("board_position")
		cleaned := strings.Replace(board_position, ",", " ", -1)
		strSlice := strings.Fields(cleaned)
		// create new slice with boolean's
		intSlice := []int {}
		for i := 0; i < len(strSlice); i++ {
			if strSlice[i] == "1" {
				intSlice = append(intSlice, i)
			}
		}
		aiMove := startCalc(intSlice, next_move)
		w.Write([]byte(aiMove))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, Null{})
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
	log.Fatal(http.ListenAndServe(":8080", r))
}