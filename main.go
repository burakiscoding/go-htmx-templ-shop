package main

import (
	"net/http"

	"github.com/burakiscoding/go-htmx-templ/handlers"
	"github.com/burakiscoding/go-htmx-templ/stores"
	"github.com/burakiscoding/go-htmx-templ/utils"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	store := stores.NewStore()
	handler := handlers.NewHandler(store)
	router.HandleFunc("/", utils.MakeHandler(handler.HandleHome)).Methods(http.MethodGet)
	router.HandleFunc("/", utils.MakeHandler(handler.HandleAdd)).Methods(http.MethodPost)
	router.HandleFunc("/{id}", utils.MakeHandler(handler.HandleProduct)).Methods(http.MethodGet)
	router.HandleFunc("/{id}", utils.MakeHandler(handler.HandleUpdate)).Methods(http.MethodPut)
	router.HandleFunc("/{id}", utils.MakeHandler(handler.HandleDelete)).Methods(http.MethodDelete)
	router.HandleFunc("/{id}/edit", utils.MakeHandler(handler.HandleEdit)).Methods(http.MethodGet)
	router.HandleFunc("/add/", utils.MakeHandler(handler.HandleShowAddForm)).Methods(http.MethodGet)
	router.HandleFunc("/nocontent/", utils.MakeHandler(handler.HandleNoContent)).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)
}
