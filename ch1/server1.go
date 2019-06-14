package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:10086", nil))
}


func handler(w http.ResponseWriter, r *http.Request){
	_, _ = fmt.Fprintf(w, "URL.Path= %q\n", r.URL.Path)
}