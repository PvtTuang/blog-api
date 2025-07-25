package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	initDB()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/post/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/post/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/post/{id}", deletePost).Methods("DELETE")

	log.Println("เซิร์ฟเวอร์กำลังทำงานอยู่บน :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
