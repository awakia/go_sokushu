package main

import (
	"fmt"
	"net/http"

	neg "github.com/codegangsta/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})
	n := neg.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
