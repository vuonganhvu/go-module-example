package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"go-module-example/controller"
)

func MakeRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", wrapHandler(controller.Home)).Methods("GET")
	router.HandleFunc("/api/login", controller.Login).Methods("POST")
	return router
}

func wrapHandler(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	h := func(w http.ResponseWriter, r *http.Request) {
		if !isAuthenticate(r) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
	return h
}

func isAuthenticate(r *http.Request) bool {
	return true
}