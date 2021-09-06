package main

import (
	"fmt"
	"os"
	"net/http"
    "sample/pkg/api"
	//"github.com/PacktPublishing/Cloud-Native-Go/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", api.EchoHandleFunc)
	//http.HandleFunc("/api/hello", api.HelloHandleFunc)

	http.HandleFunc("/api/song", api.SongsHandleFunc) // store and retrieve all songs
	http.HandleFunc("/api/song/", api.SongHandleFunc) // retrieve songs by IDs, update nad delete
	//http.HandleFunc("/api/books/", api.BookHandleFunc)

	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Cloud Native Go!")
}
