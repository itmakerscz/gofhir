package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http server address")

func main() {
	flag.Parse()

	var p People

	fmt.Println(p)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	log.Fatal(http.ListenAndServe(*addr, nil))

}
