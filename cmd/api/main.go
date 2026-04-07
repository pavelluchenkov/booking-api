package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})

}
func main(){
	http.HandleFunc("/ping", pingHandler)

	log.Println("Server started on localhost 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal("Server failed:", err)
	}

}