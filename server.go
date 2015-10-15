package main

import (
	"fmt"
	"net/http"

	"github.com/awakia/go_sokushu/models"
	neg "github.com/codegangsta/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%v", models.NewUser("Naoyoshi Aikawa", 29))
	})
	n := neg.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
