package router

import (
	"fmt"
	"net/http"
	"test_task/api"
	"test_task/configs"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()

	router.HandleFunc("/currency/save/{date}", api.SaveAPI)
	router.HandleFunc("/currency/{date}/{code}", api.GetCurrency)

	fmt.Println("HTTP server started...")
	http.ListenAndServe(fmt.Sprintf("%s:%s", configs.Config.ListenIP, configs.Config.ListenPort), router)
}
