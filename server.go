package main

import (
	"log"
	"net/http"
	"strconv"

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
	mux.HandleFunc("/index", func(w http.ResponseWriter, req *http.Request) {
		users := models.AllUsers()
		ren.JSON(w, http.StatusOK, users)
	})
	mux.HandleFunc("/show", func(w http.ResponseWriter, req *http.Request) {
		id, err := strconv.Atoi(req.FormValue("id"))
		if err != nil {
			log.Println(err)
		}
		user := models.GetUser(id)
		ren.JSON(w, http.StatusOK, user)
	})
	n := neg.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
