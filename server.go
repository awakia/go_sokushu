package main

import (
	"net/http"

	"github.com/awakia/go_sokushu/models"
	neg "github.com/codegangsta/negroni"
	"github.com/unrolled/render"
)

func main() {
	ren := render.New()
	mux := http.NewServeMux()
	mux.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		ren.JSON(w, http.StatusOK, models.NewUser("Naoyoshi Aikawa", 29))
	})
	n := neg.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
