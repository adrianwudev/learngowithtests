package main

import (
	"log"
	"net/http"

	"github.com/adrianwudev/learngowithtests/iowriter"
)

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	iowriter.Greet(w, "net/http world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
}
