package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var memos []Memo
var idCounter int

type Memo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func createMemo(w http.ResponseWriter, r *http.Request) {
	var memo Memo
	json.NewDecoder(r.Body).Decode(&memo)
	idCounter++
	memo.ID = idCounter
	memos = append(memos, memo)
	json.NewEncoder(w).Encode(memo)
}

func getMemos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(memos)
}

func getMemo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, memo := range memos {
		if memo.ID == id {
			json.NewEncoder(w).Encode(memo)
			return
		}
	}
	http.Error(w, "Memo Not Found", http.StatusNotFound)
}

func updateMemo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, memo := range memos {
		if memo.ID == id {
			json.NewDecoder(r.Body).Decode(&memo)
			memo.ID = id
			memos[i] = memo
			json.NewEncoder(w).Encode(memo)
			return
		}
	}
	http.Error(w, "Memo Not Found", http.StatusNotFound)
}

func deleteMemo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, memo := range memos {
		if memo.ID == id {
			memos = append(memos[:i], memos[i+1:]...)
			json.NewEncoder(w).Encode("the memo was deleted successfully")
			return
		}
	}
	http.Error(w, "Memo Not Found", http.StatusNotFound)
}

func initializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/memos", createMemo).Methods("POST")
	router.HandleFunc("/memos", getMemos).Methods("GET")
	router.HandleFunc("/memos/{id}", getMemo).Methods("GET")
	router.HandleFunc("/memos/{id}", updateMemo).Methods("PUT")
	router.HandleFunc("/memos/{id}", deleteMemo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	memos = append(memos, Memo{
		ID:      1,
		Title:   "First memo",
		Content: "Hello World",
	})
	idCounter = 1
	initializeRouter()
}
