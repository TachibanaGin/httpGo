package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":7000", nil))

}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "hello , go ")
	fmt.Println("hello , go ")

}

