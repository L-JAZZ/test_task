package router

import "github.com/gorilla/mux"

func StartServer() {
	router := mux.NewRouter()

	router.HandleFunc("/", homeLink)
}
