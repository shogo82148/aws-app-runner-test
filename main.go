package main

import (
	"io"
	"log"
	"net/http"
	"runtime"
)

func main() {
	log.Println("starting the server...")
	err := http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello App Runner on "+runtime.GOOS+"/"+runtime.GOARCH)
	}))
	if err != nil {
		log.Fatal(err)
	}
}
