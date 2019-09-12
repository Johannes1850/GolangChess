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

func startCalc(slice []int) {
	var aiPlayer AiPlayer
	aiPlayer.init(slice)
}

func receiveAjax(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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
		startCalc(intSlice)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
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