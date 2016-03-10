package server

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github/elmoswelt/Geekon-2016/Source/server/routes"
)

type Server struct  {

}

func (s *Server)Start() {

	setupRouter()
}

func setupRouter() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handleIndex)
	router.HandleFunc("/v1/diff", routes.Diff)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("I was here."))
}
