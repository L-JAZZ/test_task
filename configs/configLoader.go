package configs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ConfigStruct struct {
	Url        string        `json:"url"`
	ListenIP   string        `json:"listenIP"`
	ListenPort string        `json:"listenPort"`
	DB         DatabaseModel `json:"db"`
}

type DatabaseModel struct {
	Driver   string `json:"driver"`
	DbServer string `json:"db_server"`
	DbPort   int    `json:"db_port"`
	DbName   string `json:"db_name"`
	DbUser   string `json:"db_user"`
	DbPasswd string `json:"db_passwd"`
}

var Config ConfigStruct

func LoadConfiguration(file string) {

	configFile, err := os.Open(file)
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			log.Println("[LoadConfiguration]: .Close() fail")
		}
	}(configFile)

	if err != nil {
		fmt.Println("[LoadConfiguration]: fail")
		panic(fmt.Sprintf("Ошибка чтения конфигурации: %+v", err))
	}
	//fmt.Println("[LoadConfiguration]: success")

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&Config)
	if err != nil {
		return
	}
}
