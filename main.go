package main

import (
	"test_task/configs"
	"test_task/router"
)

func main() {
	configs.LoadConfiguration("config.json")

	router.StartServer()
}
